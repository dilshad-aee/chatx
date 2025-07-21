package main

import(
"fmt"
"net/http"
"github.com/gorilla/websocket"
"time"
"sync"
)

var upgrader=websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct{
	ID : 

}