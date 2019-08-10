package commands

import d "hello/app/device"
import "strings"
import "errors"

import auth "hello/app/auth"
import "strconv"

type Command interface {
	Execute(params[] string) (string,error)
}

type AuthDecoratorCommand struct{
	Command Command
}

type LoginCommand struct {
	Name string
}

type ListDevicesCommand struct{
}

type SwitchCommand struct{
}

type SetCommand struct{

}

func(a *AuthDecoratorCommand) Execute(params[] string) (string,error) {


}

// parts := strings.Split(params, ",")
func (p *LoginCommand) Execute(params[] string) (string,error) {
		
	if len(params) != 1{
		return "",errors.New("LoginCommand should contain only one paramter")
	}

	userName := parts[0]

	auth.Login(userName)

	return userName + " Is Logged in.",nil
}

func (c *ListDevicesCommand) Execute() string {
	
	devicesStr := ""

	for _, device := range d.Devices {
		devicesStr += device.String()	
	}

	return devicesStr

}

func (p *SwitchCommand) Execute(params string) (string,error) {
	
	parts := strings.Split(params, ",")
	
	if len(parts) != 2{
		return "",errors.New("SwitchCommand should contain only 2 paramters")
	}

	state, err := strconv.Atoi(parts[0])
	deviceId, err := strconv.Atoi(parts[1])

	d.Devices[deviceId].State = state
}

func (p *SetCommand) Execute(params string) string {
	return "react pings"
}