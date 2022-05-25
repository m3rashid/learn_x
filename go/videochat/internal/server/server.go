package server

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
	"github.com/m3rashid/learn_x/go/videochat/internal/handlers"
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

	engine := html.new("./views", ".html")
	app := fiber.New(fiber.Config{ Views: engine })
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.roomwebsocket, websocket.Config{
		HandshakeTimeout: 10*time.Second,
	}))

	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomChatWebsocket))

	app.Get("/stream/:uuid", handlers.Stream)
	app.Get("/stream/:uuid/websocket")
	app.Get("/stream/:uuid/chat/websocket")
	app.Get("/stream/:uuid/viewer/websocket")
}
