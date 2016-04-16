package game

import (
	"fmt"
	"encoding/json"
	"udpServer"
	"time"
)

var players []map[string]interface{}

func NewPlayer(id int) {
	// TODO: thread safe (write)
	n := len(players)
	if n == cap(players) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
		
		newSlice := make([]map[string]interface{}, n, 2*n+1)
        copy(newSlice, players)
        players = newSlice
    }
    players = players[0 : n+1]
    players[n] = map[string]interface{}{
		"id": id,
		"name": "Vrixyz",
		"position": map[string]interface{}{
			"x": 1,
			"y": 2},
    }
}

func LostPlayer(id int) {
	// TODO: thread safe (write)
	for i := range players {
		if (players[i]["id"] == id) {
			players[i] = players[len(players) - 1]
			players[len(players) - 1] = nil
		}
	}
}

func Start() {
	
    i := 0
    for {
        i++
		// TODO: thread safe (read)
        buf, err := json.Marshal(players)
		if (err != nil) {
			fmt.Println("Couldn't marshall player", err)
		} else {
			fmt.Println(players)
			//fmt.Println(buf)
			udpServer.Broadcast(buf)
		}
        time.Sleep(time.Second * 1)
    }
}