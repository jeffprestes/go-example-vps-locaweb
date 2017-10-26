package handler

import (
	"net/http"

	"github.com/jeffprestes/go-example-vps-locaweb/lib/context"
)

//HTTP2TestHandler Manipula uma chamada a uma pagina HTML
func HTTP2TestHandler(ctx *context.Context) {
	ctx.NativeHTML(http.StatusOK, "http2")
}
