package device

import "strconv"

var Devices =  map[int]Device { 1 :Device{ Id : 1, Name : "Lamp" }, 
							   2: Device{ Id : 2, Name : "AirConditioner"} }

type Device struct{
	Id int
	Name string
	Value int
	State int
}



func (d *Device) String() string  {
	return d.Name + "," + strconv.Itoa(d.Value) + "," + strconv.Itoa(d.State)
}


