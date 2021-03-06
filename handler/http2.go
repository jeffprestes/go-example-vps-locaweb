package handler

import (
	"net/http"

	"github.com/jeffprestes/go-example-vps-locaweb/lib/context"
)

func HTTP2Handler(ctx *context.Context) {
	pusher, err := ctx.Resp.Push()
	if err != nil {
		ctx.PlainText(500, []byte("the ResponseWriter doesn't support the Pusher interface"))
		return
	}
	pusher.Push("/css/bootstrap.min.css", nil)
	pusher.Push("/js/bootstrap.min.js", nil)
	ctx.NativeHTML(http.StatusOK, "http2")
}
