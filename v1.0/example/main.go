package main

import (
	"log"
	"net/http"

	"github.com/cndpost/go-socket.io/v1.0"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			m := make(map[string]interface{})
			m["a"] = "你好"
			e := so.Emit("cn1111", m)

			b := make(map[string]string)
			b["u-a"] = "中文内容" //这个不能是中文
			m["b-c"] = b
			e = so.Emit("cn2222", m)
			log.Println(e)

			log.Println("emit:", so.Emit("chat message", msg))
			server.BroadcastTo("chat", "chat message", msg)
		})
		// Socket.io acknowledgement example
		// The return type may vary depending on whether you will return
		// For this example it is "string" type
		so.On("chat message with ack", func(msg string) string {
			return msg
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
