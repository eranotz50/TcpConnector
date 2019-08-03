package connector

import "fmt"
import "strings"
import "net"
import "log"
import "bufio"


type TcpConnector struct{
	Socket net.Conn 
	User   string
	IsLoggedIn bool
	isRunning bool
}


func (c TcpConnector) ToString() string {
	return c.Socket.RemoteAddr().String();
}

// (string, error)
func(c TcpConnector)  StartRecive() {	

	for {
		
		netData, err := bufio.NewReader(c.Socket).ReadString('\n')	
		
		if(err != nil){
			log.Fatal("Error from TcpConnector.StartRecive()", err)
		}

		messageStr := strings.TrimSpace(string(netData))

		fmt.Printf("%s",messageStr)

		if(!c.isRunning){
			break
		}	
	}
	

	//return netData, err
}

/*
func (c TcpConnector) handleConnection() {
	
	fmt.Printf("Serving %s\n",c.ToString() )
	
	for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
					fmt.Println(err)
					return
			}

			temp := strings.TrimSpace(string(netData))
			if temp == "STOP" {
					break
			}

			result := strconv.Itoa(random()) + "\n"
			c.Write([]byte(string(result)))
	}
	c.Close()
}*/