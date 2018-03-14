package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"ewallet/api/app"

	"github.com/didip/tollbooth"
	"github.com/go-zoo/bone"
	"github.com/justinas/alice"
	"github.com/onrik/ethrpc"
	"github.com/rs/cors"
	"github.com/russross/blackfriday"
	httpware "github.com/terryh/go-httpware"

	"github.com/unrolled/render"
)

var (
	AppContext     *app.Context
	Render         *render.Render
	templateForDoc *template.Template
	ethClient      *ethrpc.EthRPC
)

type Result struct {
	State   int
	Content map[string]interface{}
}

func New() *Result {
	r := Result{0, map[string]interface{}{"error": ""}}
	return &r
}

//DocHandler document
func DocHandler(w http.ResponseWriter, r *http.Request) {
	readme := FSMustByte(AppContext.Debug, "/static/README.md")
	content := blackfriday.MarkdownCommon(readme)

	var doc = struct {
		AUTHOR      string
		PKG_NAME    string
		VERSION     string
		LAST_UPDATE string
		CONTENT     template.HTML
	}{app.Authors, app.PackageName, app.Version, app.LastUpdated, template.HTML(content)}

	w.Header().Set("Content-type", "text/html;charset=utf-8")
	templateForDoc.Execute(w, doc)
}

func Main(context *app.Context) *bone.Mux {
	// init template
	Render = render.New(render.Options{
		Directory:  "/static",
		Extensions: []string{".html"},
		Asset: func(name string) ([]byte, error) {
			return FSByte(context.Debug, name)
		},
		AssetNames: func() []string {
			return []string{
				"/static/index.html",
			}
		},
	})

	templateForDoc, _ = template.New("doc").Parse(FSMustString(context.Debug, "/static/doc.html"))

	// middleware
	common := alice.New(
		httpware.SimpleLogger,
		httpware.Recovery,
		cors.New(cors.Options{AllowedHeaders: []string{"*"}, AllowCredentials: true}).Handler,
	)

	// in production mode
	if !context.Debug {
		common = common.Append(
			// one minute no more than 10 from same ip address
			httpware.Limiter(tollbooth.NewLimiter(10, time.Minute, nil)),
		)
	}

	mux := bone.New()

	// html
	mux.Get("/", http.HandlerFunc(IndexHandler))

	// handlereth.go
	mux.Get("/node", common.ThenFunc(nodeInfo))
	mux.Get("/block/:block_number", common.ThenFunc(blockInfo))
	mux.Get("/transation/:transation_hash", common.ThenFunc(transationInfo))
	mux.Put("/startminer", common.ThenFunc(startMiner))
	mux.Delete("/stopminer", common.ThenFunc(stopMiner))
	mux.Post("/sendtrans", common.ThenFunc(sendTrans))

	// static
	mux.Get("/apidoc", http.HandlerFunc(DocHandler))
	return mux
}

func main() {

	// Start app service
	AppContext = app.NewContext()

	mux := Main(AppContext)

	apiUrl := "http://127.0.0.1:8545"

	ethClient = ethrpc.New(apiUrl)

	log.Printf("Starting HTTP service on %s ...", AppContext.Port)
	http.ListenAndServe(AppContext.Port, mux)

}
