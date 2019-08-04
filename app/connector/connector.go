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


func (c TcpConnector) String() string {
	return c.Socket.RemoteAddr().String();
}

func BuildMenu() string {
	return "(1) Login \n(2) List Devices \n(3) Switch On\\Off \n (4) Set"
}

type packet func(string)

// (string, error)
func(c TcpConnector)  StartReceive(onPacket packet) {	

	c.isRunning = true

	var menu = BuildMenu();
	bufio.NewWriter(c.Socket).WriteString(menu)

	for {
				
		netData, err := bufio.NewReader(c.Socket).ReadString('\n')	
				
		if(err != nil){
			log.Fatal("Error from TcpConnector.StartRecive()", err)
		}
			
		netData = strings.TrimSpace(string(netData))
		onPacket(netData)

		//fmt.Println(netData)

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