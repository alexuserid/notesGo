/*For example, an outlet allows plugging many different devices: 
computer, iron, phone charger etc. All because it has universal type of socket (interface).
So interface (set of methods) in Go allows to use many different implementations of structs 
which have the same interface, that includes the same methods.
If no plugs (or structs) would use the same interfaces 
we would have to use different sockets (or sets of methods) 
every time we want to plug in another one device (or struct), 
even if it's completely identical.*/


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
//Thus, we haven’t to write some identical functions to print our d that way. 
//It's really helpful when your project is bigger.
func plug(d Device) {
  println(d.plugIn())
}


//function Plug works with d1 - DeviceComp type variable 
//plug executes following things: 
//println(d1.plugIn()). Here is a variable d1 that has interface that calls method plugIn by dot. 
//Since d1 has type DeviceComp therefore it calls method which return "Device 1"
func main() {
  var d1 Device = new(DeviceComp)
  plug(d1)

  var d2 Device = new(DeviceIron)
  plug(d2)
}






/*------------ Русская версия ---------------*/
/*Например, розетки позволяют подключать к ним самые разные устройства:
компьютеры, утюги, зарядки для телефонов и т.д. Все потому, 
что у розетки универсальный тип разъема (интерфейс).
В Go интерфейс (он же набор методов) позволяет использовать различные реализации структур,
имеющих одинаковый интерфейс, который включает в себя одни методы.
Если бы все наши вилки (или структуры Go) имели разные интерфейсы,
нам бы пришлось использовать разные разъемы розеток
каждый раз, когда мы хотим подключить какое-то другое устройство с вилкой (или структуру в Go),
даже если оно неотличимо от другого.*/


package main 

//создание новой структуры
type DeviceComp struct {}  
//создание нового метода
func (d *DeviceComp) plugIn() string {   
  return "Device 1"
}

//создание еще одной структуры
type DeviceIron struct {}
//создание еще одного метода
func (d *DeviceIron) plugIn() string {
  return "Device 2"
}

//создание интерфейса с набором методов, который могут используют идентичные структуры
type Device interface {
  plugIn() string
}
 
//Создание новой функции, выводящей d с помощью включенных в наш интерфейс методов.
//Таким образом, нам не надо писать кучу одинаковых функций для вывода d таким образом.
//Все это очень кстати когда приходится иметь дело с большими проектами.
func plug(d Device) {
  println(d.plugIn())
}


//Здесь функция Plug работает с d1 - переменной типа DeviceComp
//plug выполняет следующие действия:
//println(d1.plugIn()). Здесь мы видим переменную d1, имеющую интерфейс, вызывающий метод plugIn через точку.
//Так как d1 имеет тип DeviceComp, следовательно вызывается метод, возвращающий "Device 1"
func main() {
  var d1 Device = new(DeviceComp)
  plug(d1)

  var d2 Device = new(DeviceIron)
  plug(d2)
}
