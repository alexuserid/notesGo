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



//Also there is a very good example from gobyexample.com

// _Interfaces_ are named collections of method
// signatures.

package main

import "fmt"
import "math"

// Here's a basic interface for geometric shapes.
type geometry interface {
    area() float64
    perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // The `circle` and `rect` struct types both
    // implement the `geometry` interface so we can use
    // instances of
    // these structs as arguments to `measure`.
    measure(r)
    measure(c)
}
