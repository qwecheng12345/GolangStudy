package cg

import (
	"encoding/json"
	"errors"
	"sync"
	"ipc"
)

var _, ipc.Server = &CenterServer{}

type Message struct {
	From string "from"
	To string "to"
	Content string "content"
}

type CenterServer struct {
	servers map[string] ipc.Server
	players [] *Player
	rooms [] *Room
	mutex sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string] ipc.Server)
	players := make([] *Player, 0)
	
	return &CenterServer{servers : servers, players : players}
}