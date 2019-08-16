package connector

import "fmt"
import "strings"
import "net"
//import "log"
import "bufio"


type TcpConnector struct{
	Socket net.Conn 
	UserName   string
	isRunning bool
}


func (c TcpConnector) String() string {
	return c.Socket.RemoteAddr().String();
}


func(c TcpConnector)  StartReceive(onPacket func(string)) {	

	c.isRunning = true
	

	//bufio.NewWriter(c.Socket).WriteString(menu)

	for {
				
		netData, err := bufio.NewReader(c.Socket).ReadString('\n')	
				
		if err != nil {
			//log.Fatal("Error from TcpConnector.StartRecive()", err)
			c.isRunning = false
		}else{
			netData = strings.TrimSpace(string(netData))
			onPacket(netData)
		}
					
		if(!c.isRunning){
			fmt.Println("Disconnected -> " +  c.Socket.RemoteAddr().String())
			break
		}	
	}	
}

func(c TcpConnector) Send(msg string){
	buf := 	[]byte(msg)
	c.Socket.Write(buf)
}
