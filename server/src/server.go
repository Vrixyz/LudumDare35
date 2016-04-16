package main
 
import "udpServer"



func main() {
	
    defer udpServer.Stop()
	udpServer.Start()
}