package commands

import d "hello/app/device"
//import "strings"
import "errors"

import c "hello/app/connector"
import "strconv"



type Command interface {
	Execute(c c.TcpConnector,params []string) (string,error)
}

type AuthDecoratorCommand struct{
	Command Command
}

type LoginCommand struct {

}

type ListDevicesCommand struct{
	// todo : validate method in each command
}

type SwitchCommand struct{

}

type SetCommand struct{

}



func(a *AuthDecoratorCommand) Execute(c c.TcpConnector,params []string) (string,error){	
	if c.UserName != "" {
		return a.Command.Execute(c,params)
	}
	
	return "Login Required.", nil
}

// parts := strings.Split(params, ",")
func (p *LoginCommand) Execute(c c.TcpConnector,params []string) (string,error){	
		
	if len(params) != 1{
		return "",errors.New("LoginCommand should contain only one paramter")
	}
	
	c.UserName = params[1]

	return c.UserName + " Is Logged in.",nil
}

func (p *ListDevicesCommand) Execute(c c.TcpConnector,params []string) (string,error){	
	
	devicesStr := ""

	for _, device := range d.Devices {
		devicesStr += device.String()	
	}

	return devicesStr,nil

}

func (p *SwitchCommand) Execute(c c.TcpConnector,params []string) (string,error){	
	
	
	if len(params) != 2{
		return "",errors.New("SwitchCommand should contain only 2 paramters")
	}

	state, err := strconv.Atoi(c.Parts[0])
	if(err != nil){
		return "",err
	}

	deviceId, err := strconv.Atoi(c.Parts[1])
	if(err != nil){
		return "",err
	}

	device, isDeviceExists := d.Devices[deviceId]	
	if isDeviceExists == false{
		return "",errors.New("SwitchCommand could not find device with Id " + string(deviceId))
	}

	device.State = state	

	return device.String(),nil
}

func (p *SetCommand) Execute(c CommandParameter) (string,error) {
	return "react pings",nil
}