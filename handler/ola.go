package handler

import (
	"net/http"
	"time"

	"github.com/jeffprestes/go-example-vps-locaweb/lib/context"
	"github.com/jeffprestes/go-example-vps-locaweb/model"
)

//Ola Manipula a chamada basica do sistema
func Ola(ctx *context.Context) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	ctx.Data["pagina"] = pagina
	ctx.NativeHTML(http.StatusOK, "ola")
}
