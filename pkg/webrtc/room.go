package webrtc

import (
	"log"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v3"
)

type Room struct {
	Peers Peers
	Hub   *chat.hub
}

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuration

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Printf(err)
		return
	}

	newPeer := PeerConnectionState{
		PeerConnection: peerConnection,
		WebSocket:      &ThreadSafeWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}
