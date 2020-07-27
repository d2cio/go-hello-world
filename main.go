package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/valyala/fasthttp"
)

func main() {
	port := "3000"
	if env, ok := os.LookupEnv("PORT"); ok {
		if _, err := strconv.ParseUint(env, 10, 16); err == nil {
			port = env
		}
	}

	hostname, _ := os.Hostname()

	handler := func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("text/html; charset=utf8")
		fmt.Fprintf(ctx, "Hello from Go!<br>Server listening port %s on host: %s", port, hostname)
	}

	if err := fasthttp.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
