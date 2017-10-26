package handler

import (
	"net/http"

	"github.com/jeffprestes/go-example-vps-locaweb/lib/context"
)

//Ola Manipula a chamada basica do sistema
func Ola(ctx *context.Context) {
	ctx.NativeHTML(http.StatusOK, "http2")
}
