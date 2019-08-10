package main

import "fmt"
import "log"
import "net"
//import d "hello/app/device"
import c "hello/app/connector"
//import cmd "hello/app/commands"


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


