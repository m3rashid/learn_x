package server

import (
	"flag"
	"os"
)

var (
	addr = flag.String("addr":"", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.Sstring("key", "", "")
)

func Run() error {
	flag.Parse()
	if addr == ":" {
		*addr = ":5000"
	}

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket")
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomChatWebsocket))

	app.Get("/stream/:uuid", handlers.Stream)
	app.Get("/stream/:uuid/websocket")
	app.Get("/stream/:uuid/chat/websocket")
	app.Get("/stream/:uuid/viewer/websocket")
}
