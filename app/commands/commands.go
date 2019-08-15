package commands

import d "hello/app/device"
import "errors"
import "strconv"



type Command interface {
	Execute(userName *string,params []string) (string,error)
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



func(a *AuthDecoratorCommand) Execute(userName *string,params []string) (string,error){	
	if *userName != "" {
		return a.Command.Execute(userName,params)
	}
	
	return "Login Required.", nil
}

// parts := strings.Split(params, ",")
func (p *LoginCommand) Execute(userName *string,params []string) (string,error){	
		
	if len(params) != 1{
		return "",errors.New("LoginCommand should contain only one paramter")
	}
	
	*userName = params[0]

	return *userName + " Is Logged in.",nil
}

func (p *ListDevicesCommand) Execute(userName *string,params []string) (string,error){	
	
	devicesStr := ""

	for _, device := range d.Devices {
		devicesStr += device.String()	
	}

	return devicesStr,nil

}

func (p *SwitchCommand) Execute(userName *string,params []string) (string,error){	
	
	
	if len(params) != 2{
		return "",errors.New("SwitchCommand should contain only 2 paramters")
	}

	state, err := strconv.Atoi(params[0])
	if(err != nil){
		return "",err
	}

	deviceId, err := strconv.Atoi(params[1])
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

func (p *SetCommand) Execute(userName *string,params []string) (string,error) {
	return "react pings",nil
}