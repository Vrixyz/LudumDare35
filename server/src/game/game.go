package game

import (
	"fmt"
	"encoding/json"
	"udpServer"
	"time"
)

var players []map[string]interface{}

const leftAction string = "left"
const upAction string = "up"
const rightAction string = "right"
const downAction string = "down"

func PlayerMessage(id int, body []byte) {
	var player map[string]interface{}
	for i := range(players) {
		if (players[i]["id"] == id) {
			player = players[i]
		}
	}
	if player == nil {
		fmt.Println("should not happen unknown player: ", id)
		return
	}
	
	var message interface{}
	err := json.Unmarshal(body, &message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("player: ", id, " ; says: ", message)
}

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
			n := len(players) - 1
			players[i] = players[n]
			players = players[:n]
		}
	}
}

func Start() {
	
    i := 0
    for {
        i++
		// TODO: thread safe (read)
		data := map[string]interface{}{
				"players": players,
				"time": time.Now(),
			}
        buf, err := json.Marshal(data)
		if (err != nil) {
			fmt.Println("Couldn't marshall player", err)
		} else {
			fmt.Println(data)
			//fmt.Println(buf)
			udpServer.Broadcast(buf)
		}
        time.Sleep(time.Second * 1)
    }
}