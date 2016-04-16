package main
 
import (
	"udpServer"
	"game"
)


func main() {
	
	udpServer.Start()
	defer udpServer.Stop()
	
	//defer game.Stop()
	game.Start()
}