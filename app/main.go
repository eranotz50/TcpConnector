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

	sss := "login eran"
	p := strings.Split(sss," ")

	fmt.Printf(p[0])

	menu := BuildMenu();
	//menuPacket := []byte(menu)

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
		
		connector := c.TcpConnector{ Socket : conn}

		fmt.Println(connector.String())
						
		connector.Send(menu)	

		go connector.StartReceive(func(msg string){

			fmt.Printf(msg)

			parameters := strings.Split(msg," ")
			commandName := parameters[0]	
						
			parameters = remove(parameters,0)
			command := commands[commandName]
			result, err := command.Execute(&connector.UserName,parameters)
			
			if err == nil{
				connector.Send(result)	
			}else{
				fmt.Println(err)
			}							
		})			
	}	
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
  }

func BuildMenu() string {
	return "(1) Login \n\r(2) List Devices \n\r(3) Switch On\\Off \n\r(4) Set\n\r"
}


func InitCommands() map[string]cmd.Command{
	commands := map[string]cmd.Command{}

	commands["login"] = &cmd.LoginCommand{}
	commands["listdevices"] = &cmd.ListDevicesCommand{}
		
	commands["switch"] = &cmd.AuthDecoratorCommand { Command : &cmd.SwitchCommand{}}  
	commands["set"] = &cmd.AuthDecoratorCommand { Command : &cmd.SetCommand{}}  

	return commands

}

