package main

import "fmt"
import "log"
import "net"
import "strings"
//import d "hello/app/device"
import c "hello/app/connector"
import cmd "hello/app/commands"


/*("fmt"
	    "log"
        "os")*/
func main() {

	commands := InitCommands()

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

			parts := strings.Split(msg,",")
			commandName := parts[0]	
			
			parameter := cmd.CommandParameter{ UserName : connector.UserName , Parts : parts }

			result, _ := commands[commandName].Execute(parameter)
			connector.Send(result)
		})			
	}	


}

func InitCommands() map[string]cmd.Command{
	commands := map[string]cmd.Command{}

	commands["login"] = &cmd.LoginCommand{}

	return commands

}

