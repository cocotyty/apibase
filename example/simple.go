package main
import (
	"github.com/cocotyty/apibase"
	"gopkg.in/macaron.v1"
)
func main() {
	api:=apibase.Api()
	api.Get("/", func(ctx *macaron.Context) {
		ctx.JSON(200,[]string{"json","api"})
	})
	api.Run(8080)
}