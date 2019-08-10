package main

import "fmt"
import "log"
import "net"
import c "hello/app/connector"
import cmd "hello/app/commands"


/*("fmt"
	    "log"
        "os")*/
func main() {
	fmt.Printf("main\n")

    ln, err := net.Listen("tcp",":9055");
   
	if err != nil{
		log.Fatal(err);		
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error from listener.Accept()", err)
		}
		
		connector := c.TcpConnector{ Socket : conn, IsLoggedIn : false}

		fmt.Println(connector.String())


		go connector.StartReceive(func(msg string){
			fmt.Println("msg -> " + msg)
		})			
	}	
}

func GetDevices() []cmd.Device{
	return []cmd.Device{cmd.Device{ Id : 1, Name : "Lamp" } }
}






	// Console.ReadKey()
	/*b := make([]byte, 10)
	chr, err := os.Stdin.Read(b);	

	fmt.Printf("%c",chr);

	if  err != nil {
	   log.Fatal(err)
	}*/