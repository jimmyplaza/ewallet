package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"ewallet/api/app"

	"github.com/go-zoo/bone"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"github.com/russross/blackfriday"
	httpware "github.com/terryh/go-httpware"

	"github.com/unrolled/render"
)

var (
	AppContext     *app.Context
	Render         *render.Render
	Curl           = &http.Client{Timeout: time.Duration(time.Second * 10)}
	templateForDoc *template.Template
	Bucket         string
)

type Result struct {
	State   int
	Content map[string]interface{}
}

func New() *Result {
	r := Result{0, map[string]interface{}{"success": "", "error": ""}}
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

	//templateForDoc.Execute(w, doc)
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

	// commonLimit := common.Append(
	// one minute no more than 100 from same ip address
	// httpware.Limiter(tollbooth.NewLimiter(30, time.Minute)),
	// )

	//in production mode
	if !context.Debug {
		common = common.Append(
		// one minute no more than 200 from same ip address
		// httpware.Limiter(tollbooth.NewLimiter(200, time.Minute)),
		)
	}

	mux := bone.New()

	// html
	mux.Get("/", http.HandlerFunc(IndexHandler))

	// handlereth.go
	mux.Get("/node", common.ThenFunc(nodeInfo))

	// static
	mux.Get("/apidoc", http.HandlerFunc(DocHandler))
	return mux
}

func main() {

	// start app service
	AppContext = app.NewContext()

	mux := Main(AppContext)

	log.Printf("Starting HTTP service on %s ...", AppContext.Port)
	http.ListenAndServe(AppContext.Port, mux)

}
