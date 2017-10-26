package handler

import (
	"net/http"

	"github.com/jeffprestes/go-example-vps-locaweb/lib/context"
)

//HTTP2TestHandler Manipula uma chamada a uma pagina HTML
func HTTP2TemplateHandler(ctx *context.Context) {
	ctx.NativeHTML(http.StatusOK, "http2")
}

func HTTP2Handler(ctx *context.Context) {
	pusher, ok := ctx.Resp.(http.Pusher)
	if !ok {
		ctx.PlainText(500, []byte("the ResponseWriter doesn't support the Pusher interface"))
		return
	}
	pusher.Push("/http2template", nil)
}
