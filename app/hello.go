package main

import "fmt"
import "log"
import "net"
import c "hello/app/connector"

/*("fmt"
	    "log"
        "os")*/
func main() {
	fmt.Printf("hello, world\n")

    ln, err := net.Listen("tcp",":9055");
   
	if err != nil{
		log.Fatal(err);		
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error from listener.Accept()", err)
		}

		
		connector := c.TcpConnector{ Socket : conn, User : "Eran", IsLoggedIn : false}
		fmt.Printf(connector.User);

		
	}


	
}






	// Console.ReadKey()
	/*b := make([]byte, 10)
	chr, err := os.Stdin.Read(b);	

	fmt.Printf("%c",chr);

	if  err != nil {
	   log.Fatal(err)
	}*/