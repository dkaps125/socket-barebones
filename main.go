package main

import (
    //"net/url"
    "fmt"

    "github.com/kataras/iris"
    "github.com/kataras/iris/websocket"
)

func main() {
    app := iris.Default()

    app.Get("/", func(ctx iris.Context) {
        ctx.ServeFile("public/views/index.html", false);
    })

    app.StaticWeb("/public", "./public")

    initWS(app)

    app.Run(iris.Addr(":3030"))
}

func initWS(app *iris.Application) {
    ws := websocket.New(websocket.Config{
        ReadBufferSize: 1024,
        WriteBufferSize: 1024,
    })

    ws.OnConnection(handleConnection)

    app.Get("/io", ws.Handler())

    app.Any("/ws.js", func(ctx iris.Context) {
        ctx.Write(websocket.ClientSource)
    })
}

func handleConnection(c websocket.Connection) {
    c.On("test", func(msg string) {
        fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)

        c.To(websocket.Broadcast).Emit("test", msg)
    })
}
