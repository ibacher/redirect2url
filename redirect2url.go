package main

import (
	"log"
	"net/http"

	"github.com/valyala/fasthttp"
)

func redirect(ctx *fasthttp.RequestCtx) {
	args := ctx.QueryArgs()
	if args.Has("redirect-to") {
		ctx.Redirect(string(args.Peek("redirect-to")), http.StatusSeeOther)
	} else {
		ctx.NotFound()
	}
}

func main() {
	log.Fatal(fasthttp.ListenAndServe(":4000", redirect))
}
