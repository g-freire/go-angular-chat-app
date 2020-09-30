package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
	"./services"
)


//The ClientManager will keep track of all the connected clients, clients that are trying to become registered,
// clients that have become destroyed and are waiting to be removed, and messages that are to be broadcasted to and from all connected clients.
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

//Each Client has a unique id, a socket connection, and a message waiting to be sent.
type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}


var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

//send to all clients
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
			fmt.Println("send:", string(message))

		}
	}
}

//goroutine for reading websocket data sent from the clients
func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}


func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})

				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
			fmt.Println("write: ", string(message))

		}
	}
}

func (c *Client) writeLoop() {
	defer func() {
		c.socket.Close()
	}()
	for {
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: "TEST"})
		c.socket.WriteJSON(string(jsonMessage))

		fmt.Println("Socket write: ", string(jsonMessage))
		time.Sleep(2 * time.Second)
	}
}

func wsHandler(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}

	u1 := uuid.Must(uuid.NewV4()) 	// Creating UUID Version 4
	client := &Client{id: u1.String(), socket: conn, send: make(chan []byte)}
	fmt.Printf("\n New client connected: %s\n", client.id)

	manager.register <- client

	//go client.read()
	//go client.write()
	go client.writeLoop()

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	host := "mongodb://rioAdmin:Rio!2020@13.78.169.175:27017/?authSource=test"
	//a := services.GetIbopInfo(host)
	//a := services.GetRealTimeInfo(host)
	a := services.GetMatrixProfile(host)

	fmt.Print(a)
	//PORT := ":4444"
	//fmt.Println("Starting application...")
	//fmt.Printf("Serving at %s", PORT)
	//go manager.start()
	//http.HandleFunc("/ws", wsHandler)
	//http.ListenAndServe(PORT, nil)

}
