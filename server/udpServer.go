package main
 
import (
    "fmt"
    "net"
    "os"
	"time"
    "strconv"
)
 
/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
 
func main() {
    /* Lets prepare a address at any address at port 10001*/   
    ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
    CheckError(err)
 
 
	go func(p_serverAddr *net.UDPAddr) {
	    /* Now listen at selected port */
		ServerConn, err := net.ListenUDP("udp", p_serverAddr)
		CheckError(err)
		defer ServerConn.Close()
		buf := make([]byte, 1024)
	 
		for {
			n,addr,err := ServerConn.ReadFromUDP(buf)
			fmt.Println("Received ",string(buf[0:n]), " from ",addr)
	 
			if err != nil {
				fmt.Println("Error: ",err)
			}
			
		}
	}(ServerAddr)
	
	ClientAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10002")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ClientAddr)
    CheckError(err)
 
    defer Conn.Close()
    i := 0
    for {
        msg := strconv.Itoa(i)
        i++
        buf := []byte(msg)
        _,err := Conn.Write(buf)
        if err != nil {
            fmt.Println(msg, err)
        }
        time.Sleep(time.Second * 1)
    }
}