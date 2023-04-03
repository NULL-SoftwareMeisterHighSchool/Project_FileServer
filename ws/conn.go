package ws

import (
	"github.com/gofiber/websocket/v2"
)

func Connect(c *websocket.Conn) {

	go read()
	go write()
}

func read() {
	
}

func write() {

}
