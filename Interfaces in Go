
/*
For example, there is an outlet that allows to plug in many different devices:
computer, iron, phone charger etc. All because it has universal type of socket (interface).
So interface (set of methods) in Go allows to use many different implementations of structs
which have the same interface, that includes the same methods.
If all plugs (or structs) wouldn't use the same interfaces 
we would have to use different sockets (or sets of methods)
every time we want to plug in another one device (or struct),
even if it's completely idenical.*/

package main 

//creating a new struct
type DeviceComp struct {}  
//creating a new method
func (d *DeviceComp) plugIn() string {   
  return "Device 1"
}

//creating another one struct
type DeviceIron struct {}
//creating another one method
func (d *DeviceIron) plugIn() string {
  return "Device 2"
}

//creating an interface that includes all methods of identical structs
type Device interface {
  plugIn() string
}

//creating a new function that prints d by methods that interface includes
//(or doing something else by interface's methods in another case).
//Thus we haven't to write some identical functions to print our d that way
//It's really helpful when your project is bigger
func plug(d Device) {
  println(d.plugIn())
}


//function Plug works with d1 - DeviceComp type variable
//Plug execute following things:
//почему мы не могли в функции Plug просто написать println(d)? Ведь d в данном случае вроде как имеет тип интерфейса, 
//в котором уже есть функция PlugIn. Почему нам надо ее еще и через точку писать? 
//Следующий комментарий верно отвечает на этот вопрос?
//println(d1.PlugIn()). Here d1 that has intrface that call method PlugIn by dot. 
//Since d1 has type DeviceComp therefore it call method which return "Device 1"
func main() {
  var d1 Device = new(DeviceComp)
  plug(d1)

  var d2 Device = new(DeviceIron)
  plug(d2)
}
