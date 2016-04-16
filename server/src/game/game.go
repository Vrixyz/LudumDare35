package game

import (
	"fmt"
	"encoding/json"
	"time"
	"udpServer"
	"game/maze"
)

type Vector2 struct {
	X float64
	Y float64
}

type Player struct {
	Id int
    Name string
    Position Vector2
	Speed Vector2
}

//var players []map[string]interface{}
var players []Player

type MoveMessage struct {
    XSpeed    float64
    YSpeed    float64
}

func PlayerMessage(id int, body []byte) {
	// TODO: thread safe (read)
	gameId := -1
	for i := range(players) {
		if (players[i].Id == id) {
			gameId = i
		}
	}
	if gameId == -1 {
		fmt.Println("should not happen unknown player: ", id)
		return
	}
	
	var message MoveMessage
	err := json.Unmarshal(body, &message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("player: ", gameId, " ; says: ", message)
	//if (message["action"] == "move") {
		// TODO: thread safe (write) // care previous lock
		players[gameId].Speed.X = message.XSpeed
		players[gameId].Speed.Y = message.YSpeed
	//}
	
}

func NewPlayer(id int) {
	// TODO: thread safe (write)
	n := len(players)
	if n == cap(players) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
		
		newSlice := make([]Player, n, 2*n+1)
        copy(newSlice, players)
        players = newSlice
    }
    players = players[0 : n+1]
    players[n] = Player{
		id,
		"Vrixyz",
		Vector2{1, 2},
		Vector2{0, 0},
    }
}

func LostPlayer(id int) {
	// TODO: thread safe (write)
	for i := range players {
		if (players[i].Id == id) {
			n := len(players) - 1
			players[i] = players[n]
			players = players[:n]
		}
	}
}

func Start() {
	maze.Parse("maps/exampleMap.txt")
	// goroutine send data to players
    go func () {
		for {
			// TODO: thread safe (read)
			data := map[string]interface{}{
					"players": players,
					"time": time.Now(),
				}
			buf, err := json.Marshal(data)
			if (err != nil) {
				fmt.Println("Couldn't marshall player", err)
			} else {
				fmt.Println("sending: ", string(buf))
				//fmt.Println(buf)
				udpServer.Broadcast(buf)
			}
			time.Sleep(time.Millisecond * 200)
		}
	}	()
	
	lastTick := time.Now()
	for {
		elapsedTime := time.Now().Sub(lastTick)
		lastTick = time.Now()
		// TODO: thread safe
		for i := range players {
			// TODO: optimize this shit
			fmt.Println("moving: ", players[i])
			players[i].Position.X = players[i].Position.X + (players[i].Speed.X * float64(elapsedTime.Seconds()))
			players[i].Position.Y = players[i].Position.Y + (players[i].Speed.Y * float64(elapsedTime.Seconds()))
		}
		time.Sleep(time.Millisecond * 200)
	}
}