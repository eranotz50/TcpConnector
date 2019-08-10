package connector

//import "fmt"
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

func BuildMenu() string {
	return "(1) Login \n\r(2) List Devices \n\r(3) Switch On\\Off \n\r(4) Set"
}


func(c TcpConnector)  StartReceive(onPacket func(string)) {	

	c.isRunning = true
	
	var menu = BuildMenu();
    buf := 	[]byte(menu)
	
	c.Socket.Write(buf)
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
			break
		}	
	}	
}

func(c TcpConnector) Send(msg string){
	buf := 	[]byte(msg)
	c.Socket.Write(buf)
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