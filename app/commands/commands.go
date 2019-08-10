package commands

import "strconv"

type Device struct{
	Id int
	Name string
	Value int
	IsTurnedOn bool
}

func (d *Device) String() string  {
	return d.Name + "," + strconv.Itoa(d.Value) + "," + strconv.FormatBool(d.IsTurnedOn)
}

type Command interface {
	Execute() string
}


type ListDevicesCommand struct{
	devices [] Device
}

type SwitchCommand struct{
	devices [] Device
}

type SetCommand struct{
	devices [] Device
}

func (c *ListDevicesCommand) Execute() string {
	
	devicesStr := ""

	for _, device := range c.devices {
		devicesStr += device.String()	
	}

	return devicesStr

}

func (p *SwitchCommand) Execute(params string) string {
	return "react pings"
}

func (p *SetCommand) Execute(params string) string {
	return "react pings"
}