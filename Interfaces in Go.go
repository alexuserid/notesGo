/*For example, an outlet allows plugging many different devices: 
computer, iron, phone charger etc. All because it has universal type of socket (interface).
So interface (set of methods) in Go allows to use many different implementations of structs 
which have the same interface, that includes the same methods.
If no plugs (or structs) would use the same interfaces 
we would have to use different sockets (or sets of methods) 
every time we want to plug in another one device (or struct), 
even if it's completely identical.*/


package main 

type DeviceComp struct {}  
func (d *DeviceComp) plugIn() string {   
  return "Device 1"
}

type DeviceIron struct {}
func (d *DeviceIron) plugIn() string {
  return "Device 2"
}

type Device interface {
  plugIn() string
}

func plug(d Device) {
  println(d.plugIn())
}


func main() {
  var d1 Device = new(DeviceComp)
  plug(d1)

  var d2 Device = new(DeviceIron)
  plug(d2)
}


