package device

import "strconv"

var Devices =  map[int]*Device { 1 : &Device{ name : "Lamp" }, 
							    2: &Device{ name : "AirConditioner"} }

type Device struct{
	name string
	Value int
	State int
}

func (d *Device) String() string  {
	return d.name + "," + strconv.Itoa(d.Value) + "," + strconv.Itoa(d.State)
}


