package commands

import d "hello/app/device"
//import "strings"
import "errors"

import auth "hello/app/auth"
import "strconv"


type CommandParameter struct{  
	UserName string
	Parts []string	
}

type Command interface {
	Execute(c CommandParameter) (string,error)
}

type AuthDecoratorCommand struct{
	Command Command
}

type LoginCommand struct {
}

type ListDevicesCommand struct{

}

type SwitchCommand struct{

}

type SetCommand struct{

}



func(a *AuthDecoratorCommand) Execute(params CommandParameter) (string,error) {

	if(auth.CheckLogin(params.UserName)){
		return a.Command.Execute(params)
	}

	return "Login Required",nil
}

// parts := strings.Split(params, ",")
func (p *LoginCommand) Execute(c CommandParameter) (string,error) {
		
	if len(c.Parts) != 1{
		return "",errors.New("LoginCommand should contain only one paramter")
	}
	
	auth.Login(c.UserName)

	return c.UserName + " Is Logged in.",nil
}

func (p *ListDevicesCommand) Execute(c CommandParameter) (string,error) {
	
	devicesStr := ""

	for _, device := range d.Devices {
		devicesStr += device.String()	
	}

	return devicesStr,nil

}

func (p *SwitchCommand) Execute(c CommandParameter) (string,error) {
	
	
	if len(c.Parts) != 2{
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