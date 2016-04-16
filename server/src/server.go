package main
 
import (
	"udpServer"
	"game"
)


func main() {
	udpServer.NewConnectionCallback = game.NewPlayer
	udpServer.LostConnectionCallback = game.LostPlayer
	
	udpServer.Start()
	defer udpServer.Stop()
	
	//defer game.Stop()
	game.Start()
}