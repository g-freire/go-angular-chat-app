package main

import (
	"fmt"
	"log"
	"time"
	"context"
	"github.com/gorilla/websocket"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleVersion(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "V 1.0.0")
}

// IBOP WEB SOCKET HANDLER
// Test ws status @ https://www.websocket.org/echo.html
func handlerRoot(writer http.ResponseWriter, request *http.Request){
	enableCors(&writer)
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
	}
	// WRITE ONLY STREAM INFINITE LOOP
	for {
		//wsMsg := []byte("TEST MSG FROM GO SERVER")
		// STRING 1 / BYTES 2
		//err = socket.WriteMessage(1, wsMsg)
		//JSON
		wsMsg := queryLastIBOP()
		err =socket.WriteJSON(wsMsg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("SENDING MESSAGE", wsMsg)
		time.Sleep(1 * time.Second)
	}}


// IBOP QUERY DATA
type IBOPInfo struct {
	ID      *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	IbopId 		 string   `json:"ibop_id" bson:"ibop_id"`
	CurrentState       string   `json:"current_state" bson:"current_state"`
	IsActive     bool `json:"is_active" bson:"is_active"`
}

func queryLastIBOP() IBOPInfo{
	host := ""
	collection := ""
	db := ""

	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(host)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	//QUERY RAW BSON
	var ibopInfo IBOPInfo
	collectionIBOP := client.Database(db).Collection(collection)
	//var result bson.M
	filter := bson.M{}
	err = collectionIBOP.FindOne(ctx, filter).Decode(&ibopInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", ibopInfo)
	return ibopInfo
}


func main(){
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/v", handleVersion)
	http.ListenAndServe(":5000", nil)
}