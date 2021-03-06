package udpServer
 
import (
    "fmt"
    "net"
    "os"
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

var clients []*net.UDPAddr
var keepAlive [] time.Time // time last update

var udpConn *net.UDPConn

var MyAddress string

/// Very important so my sockets stay valid..
func Stop() {
	// TODO: thread safe
	for i := range clients {
		if (clients[i] != nil) {
			//clients[i].Close()
			clients[i] = nil
		}
    }
}

func Broadcast(buf []byte) {
	// TODO: thread safe
	for i:= range clients {
		if (clients[i] != nil) {
			fmt.Println("sending: ", i, " for: ", clients[i])
			_,err := udpConn.WriteTo(buf, clients[i])
			if err != nil {
				fmt.Println(buf, err)
			}
		}
	}
}

func extend(slice []*net.UDPAddr, element *net.UDPAddr) ([]*net.UDPAddr, int) {
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
			newSlice := make([]*net.UDPAddr, n, newCap)
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
	clients[id] = nil
}

func maybeNewClient(addr *net.UDPAddr) int {
	// TODO: handle thrade safe
	//addrSimple := strings.Split(addr.String(), ":")[0];
	for i:= range clients {
		if (clients[i] != nil) {
			//clientSimpleAddr := strings.Split(clients[i].RemoteAddr().String(), ":")[0]
			fmt.Println("addr: ", addr , " ; client[", i, "]: ", clients[i].String())
			if (addr.String() == clients[i].String()) {
				fmt.Println("not a new client.")
				keepAlive[i] = time.Now()
				return i // we already saved this client
			}
		}
	}
	fmt.Println("new client!");

	newLen := 0
	clients, newLen = extend(clients, addr)
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
    ServerAddr,err := net.ResolveUDPAddr("udp4",":10003")
    CheckError(err)
	fmt.Println(ServerAddr)
 
 
	go func(p_serverAddr *net.UDPAddr) {
	    /* Now listen at selected port */
		ServerConn, err := net.ListenUDP("udp4", p_serverAddr)
		CheckError(err)
		udpConn = ServerConn
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