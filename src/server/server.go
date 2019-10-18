// server.go
package server

import (
	//"encoding/json"
	"log"
	//"strconv"
	"strings"

	"github.com/fasthttp/router"
	//"github.com/google/uuid"
	//"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	log.Println("/")
	fasthttp.ServeFile(ctx, "../html/index.html")
}

func Link(ctx *fasthttp.RequestCtx) {
	log.Println("/link")
	ctx.SetContentType("text/html")
	ctx.Write([]byte("<a href='/'>Back</a>"))
}

func GetResource(ctx *fasthttp.RequestCtx) {
	log.Println(string(ctx.RequestURI()))
	var uriSegements = strings.Split(string(ctx.RequestURI()), "/")
	var imgfile = uriSegements[len(uriSegements)-1]
	// Path where the public resource files (i.e. images) are stored
	var str = "../html/resources/" + imgfile
	fasthttp.ServeFile(ctx, str)
}

func Submit(ctx *fasthttp.RequestCtx) {
	log.Println("/submit")
	log.Println(ctx.PostArgs())
	//log.Println(ctx.FormValue("firstname"))
	//log.Printf("Firstname = %s\n", ctx.FormValue("firstname"))
	ctx.Write(append(ctx.FormValue("firstname"), " welcome!"...))
}

// Cross Origin Resource Sharing for Frontend-Backend-Communication
func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next(ctx)
	}
}

func Run() {
	log.Println("Server running on port 8000")
	// setup routing
	r := router.New()
	r.GET("/", Index)
	r.GET("/link", Link)
	r.POST("/submit", Submit)
	r.GET("/resources/:resource", GetResource)
	// setting listening port
	err := fasthttp.ListenAndServe(":8000", CORS(r.Handler))
	// err := fasthttp.ListenAndServeTLS(":443", "cert.pem", "key.pem", CORS(r.Handler))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
