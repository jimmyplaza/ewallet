package app

// ClearWatch worker

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	Version     = "1.0"
	PackageName = "api"
	LastUpdated = "2018/03/13"
	Authors     = "Jimmy Ko"

	// share in application
	App         *Context
	Timezone, _ = time.LoadLocation("Asia/Taipei")
	debug       bool
	port        = ""
	env         = ""
)

// Context
type Context struct {
	Timezone *time.Location
	Port     string
	Debug    bool
	Env      string
}

// ContextInit for initialize
func ContextInit(rurl string, dburi []string, port, env string, debug bool) *Context {
	App = new(Context)
	App.Port = port
	App.Timezone = Timezone
	App.Debug = debug
	App.Env = env
	return App
}

func init() {
	// Give the default value here
	flag.StringVar(&port, "port", ":3000", `address for listen default is :3000`)
	flag.BoolVar(&debug, "debug", false, `Flag for DEBUG, Default is: false`)
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("DEBUG") != "" {
		debug = true
	}

	if os.Getenv("ENV") != "" {
		env = os.Getenv("ENV")
	}

	fmt.Println(PackageName, Version)
	if debug {
		fmt.Println("Running in DEBUG mode")
	}
}

func printhelp() {
	fmt.Println("Name:", PackageName, Version)
	fmt.Println("Usage:")
	flag.PrintDefaults()
}

func NewContext() *Context {
	var dbs []string
	var redisurl string
	return ContextInit(redisurl, dbs, port, env, debug)
}
