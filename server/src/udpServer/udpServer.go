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
var MessageReceivedCallback func(int, []byte)


/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

var clients []*net.UDPConn
var keepAlive [] time.Time // time last update

var MyAddress string

/// Very important so my sockets stay valid..
func Stop() {
	// TODO: thread safe
	for i := range clients {
		if (clients[i] != nil) {
			clients[i].Close()
			clients[i] = nil
		}
    }
}

func Broadcast(buf []byte) {
	// TODO: thread safe
	for i:= range clients {
		if (clients[i] != nil) {
			fmt.Println("sending: ", i, " for: ", clients[i].RemoteAddr())
			_,err := clients[i].Write(buf)
			if err != nil {
				fmt.Println(buf, err)
			}
		}
	}
}

func extend(slice []*net.UDPConn, element *net.UDPConn) ([]*net.UDPConn, int) {
    n := len(slice)
	placeFound := n
	for i := range clients {
		if clients[i] == nil {
			placeFound = i
			break
		}
	}
	fmt.Println("placeFound:", placeFound)
	if (placeFound == n) {
		n := len(slice)
		if n == cap(slice) {
			fmt.Println("must grow")
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
		keepAlive = keepAlive[0 : n+1]
	}
	fmt.Println(placeFound)
	fmt.Println("keepAlive: ", keepAlive[0])
    slice[placeFound] = element
    return slice, placeFound
}

func deleteClient(id int) {
	LostConnectionCallback(id)
	clients[id].Close()
	clients[id] = nil
}

func maybeNewClient(addr *net.UDPAddr) int {
	// TODO: handle thrade safe
	addrSimple := strings.Split(addr.String(), ":")[0];
	for i:= range clients {
		if (clients[i] != nil) {
			clientSimpleAddr := strings.Split(clients[i].RemoteAddr().String(), ":")[0]
			fmt.Println("addr: ", addrSimple , " ; client[", i, "]: ", clientSimpleAddr)
			if (addrSimple == clientSimpleAddr) {
				fmt.Println("not a new client.")
				keepAlive[i] = time.Now()
				return i // we already saved this client
			}
		}
	}
	fmt.Println("new client!");
	addr.Port = 10002
 
	MyAddr, err := net.ResolveUDPAddr("udp", MyAddress)
	CheckError(err)
	Conn, err := net.DialUDP("udp", MyAddr, addr)
	CheckError(err)
	newLen := 0
	clients, newLen = extend(clients, Conn)
	keepAlive[newLen] = time.Now()
	NewConnectionCallback(newLen)
	return newLen
}

func Start() {
	address, err := ExternalIP()
	CheckError(err)
	MyAddress = address + ":0"
	fmt.Println( "my address: ", MyAddress)
    /* Lets prepare a address at any address at port 10003*/   
    ServerAddr,err := net.ResolveUDPAddr("udp",":10003")
    CheckError(err)
	fmt.Println(ServerAddr)
 
 
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
			id := maybeNewClient(addr); // TODO: goroutine to avoid packet loss
			MessageReceivedCallback(id, buf[0:n])
		}
	}(ServerAddr)
	
	go func() {
		timeout := time.Second * 5
		for {
			now := time.Now()
			// TODO: thread safe
			for i := range clients {
				if (clients[i] != nil) {
					fmt.Println(now.Sub(keepAlive[i]))
					if (now.Sub(keepAlive[i]) > timeout) {
						deleteClient(i)
					}
				}
			}
			time.Sleep(timeout)
		}
	}()
}