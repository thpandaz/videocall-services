package handlers

import (
	"fmt"
	"os"
	w "videochat-system/webrtc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Stream(c *fiber.Ctx) error {
	suuid := c.Params("suuid")
	if suuid == "" {
		c.Status(400)
		return nil
	}

	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		ws = "wss"
	}

	w.RoomsLock.Lock()
	if _, ok := w.Stream[suuid]; !ok {
		w.RoomsLock.Unlock()
		return c.Render("stream", fiber.Map{
			"StreamWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/websocket", ws, c.Hostname(), suuid),
			"ChatWebsocketAddr":   fmt.Sprintf("%s://%s/stream/%s/chat/websocket", ws, c.Hostname(), suuid),
			"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/viewer/websocket", ws, c.Hostname(), suuid),
			"Type":                "stream",
		}, "layouts/main")
	}
	w.RoomsLock.Unlock()
	return c.Render("stream", fiber.Map{
		"NoStream": true,
		"Leave":    true,
	}, "layouts/main")
}

func StreamWebsocket(c *websocket.Conn) {

}

func StreamViewerWebSocket(c *websocket.Conn) {

}

func ViewerConn(c *websocket.Conn, p *w.Peers) {

}
