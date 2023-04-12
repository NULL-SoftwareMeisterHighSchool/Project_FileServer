package core_client

import (
	"github.com/gofiber/websocket/v2"
)

// var requestMap = make(map[string]chan interface{})

func Connect(c *websocket.Conn) {
	// TODO: create map of request ids so that i can organize my requests

	go read(c)
	go write(c)
}

func read(c *websocket.Conn) {
	defer c.Close()
	for {
		// c.ReadJSON()
	}
}

func write(c *websocket.Conn) {
	defer c.Close()
	for {
		// c.WriteJSON()
	}
}
