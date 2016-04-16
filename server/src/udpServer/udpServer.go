package udpServer
 
import (
    "fmt"
    "net"
    "os"
	"strings"
	"time"
)

var NewConnectionCallback func(int)
var LostConnectionCallback func(int)


/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

var clients []*net.UDPConn
var keepAlive [] time.Time // time last update


/// Very important so my sockets stay valid..
func Stop() {
	// TODO: thread safe
	for i := range clients {
        clients[i].Close()
    }
}

func Broadcast(buf []byte) {
	// TODO: thread safe
	for i:= range clients {
		fmt.Println("sending: ", i)
		_,err := clients[i].Write(buf)
		if err != nil {
			fmt.Println(buf, err)
		}
	}
}

func extend(slice []*net.UDPConn, element *net.UDPConn) ([]*net.UDPConn, int) {
    n := len(slice)
	placeFound := 0//range clients
	//for placeFound {
	//	if clients[i] == nil {
	//		break
	//	}
	//}
	fmt.Println("placeFound:", placeFound)
	if (placeFound == n) {
		n := len(slice)
		if n == cap(slice) {
			// Slice is full; must grow.
			// We double its size and add 1, so if the size is zero we still grow.
			newCap := 2*n+1
			newSlice := make([]*net.UDPConn, n, newCap)
			copy(newSlice, slice)
			slice = newSlice
			newKeepAlive := make([]time.Time, n, newCap)
			copy(newKeepAlive, keepAlive)
			keepAlive = newKeepAlive
		}
		slice = slice[0 : n+1]
		
    slice[placeFound] = element
    return slice, n
}

func deleteClient(int id) {
	
}

func maybeNewClient(addr *net.UDPAddr) {
	// TODO: handle thrade safe
	addrSimple := strings.Split(addr.String(), ":")[0];
	for i:= range clients {
		clientSimpleAddr := strings.Split(clients[i].RemoteAddr().String(), ":")[0]
		fmt.Println("addr: ", addrSimple , " ; client[", i, "]: ", clientSimpleAddr)
		if (addrSimple == clientSimpleAddr) {
			fmt.Println("not a new client.")
			// TODO: refresh the keep-alive
			return // we already saved this client
		}
	}
	fmt.Println("new client!");
	addr.Port = 10002
 
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)
	Conn, err := net.DialUDP("udp", LocalAddr, addr)
	CheckError(err)
	newLen := 0
	clients, newLen = extend(clients, Conn)
	keepAlive[newLen] = time.Now()
	NewConnectionCallback(newLen)
}

func Start() {
	
    /* Lets prepare a address at any address at port 10003*/   
    ServerAddr,err := net.ResolveUDPAddr("udp",":10003")
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
			maybeNewClient(addr); // TODO: goroutine to avoid packet loss
		}
	}(ServerAddr)
	
	go func() {
		timeout := time.Second * 30
		for {
			now := time.Now()
			// TODO: thread safe
			for i := range clients {
				if (now.Sub(keepAlive[i]) > timeout) {
					deleteClient(i)
				}
			}
			time.Sleep(timeout)
		}
	}()
}