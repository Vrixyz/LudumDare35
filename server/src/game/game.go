package game

import (
	"fmt"
	"encoding/json"
	"udpServer"
	"time"
)

var player map[string]interface{}

func Start() {
	player = map[string]interface{}{
		"name": "Vrixyz",
		"position": map[string]interface{}{
			"x": 1,
			"y": 2},
    }
	
    i := 0
    for {
        i++
		//player.time = i
        buf, err := json.Marshal(player)
		if (err != nil) {
			fmt.Println("Couldn't marshall player", err)
		} else {
			fmt.Println(player)
			//fmt.Println(buf)
			udpServer.Broadcast(buf)
		}
        time.Sleep(time.Second * 1)
    }
}