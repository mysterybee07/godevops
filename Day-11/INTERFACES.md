# What is Interfaces in Go?
An interface is defined as a set of method signatures. A method's signature includes its name, parameters, and return types.

# Characteristics of Go Interfaces
1. Type-Agnostic:Interfaces are not tied to any specific type. Any type that implements all the methods in the interface is said to satisfy the interface.
2. Implicit Implementation: Unlike many other languages, Go doesnot require a type to explicitly declare that it implements an interface. If the Type has the necessary methods, it authomatically satisfies the interface.
3. Runtime Polymorphism: Interfaces enable polymorphism in Go. A single interface can be used to represent different types as long as they implement the interface methods.
4. Decoupling: Interfaces help in decoupling the code by not depending on the concrete implementation but rather on the behaviours (methods).
5. Empty Interface (Interface{}): It can hold values of any type, as every type implements at least zero methods. Its commonly used for functions that can accept any type of value.


# Different Ways to Declare Interfaces in Go
1. Single-method Interfaces:
   type Reader interface{
    Read(p []byte)(n int, err error)
   }

2. Embedding Interfaces:
   type Writer interface{
    Write(p []byte)(n int, err error)
   }
   type ReadWriter interface{
    Reader
    Writer
   }
   Go allows interfaces to be embedded within other interfaces. This is a way to compose interfaces from other interfaces.

3. Empty Interface
   type Any Interface{}
   The empty interface specifies zero methods and is satisfied by any type, similar object in Java or interface{} in C#

4. Interface with multiple methods
   type Shape interface{
    Area() float64
    Perimeter() float64
   }

5. Functional Interface (Single method interface)
   type HandlerFunc interface{
    ServeHTTP(w http ResponseWriter, r *http.Request)
   }
   An interface with only one method is often used in go , especially for call back
   functions and implementing functional options.

6. Interface as a Constraints(Generic)
   type Number interface{
    int64| float64
   }
    func Sum[T Number](a,b T)T{
        return a+b
    
   }
   With the introduction of generices in Go, interfaces can be used as a type constraints in generic functions or types.

# Common Operations in Interfaces
1. Normal
   type Shape interface{
    Area() float64
    Perimeter() float64
   }

2. Implementing an Interface
   type Rectangle struct{
    width, height float64
   }
   func(r Rectangle) Area() float64{
    return r.width*r.height
   }
   func (r Rectangle) Perimeter() float64{
    return 2*(r.width+r.height)
   }
   A type implements an interface by implementing all the methods declared in the interface.

3. Type Assertions
    var s Shape = Rectangle{
        Width:5,
        Height:3
    }
    r, ok := s.(Rectangle)
    if ok{
        fmt.Println("Rectangle width", r.width)
    }
    Type assertions are used to retrieve the dynamic value of an interface. This is useful when you need to access the specific type and its methods or properties.

4. Interface with Type Switches
   func do(i interface{}){
    switch v:= i.(type){
        case int:
        fmt.Println("interger",v)
        case string:
        fmt.Println("string",v)
        default:
        fmt.Println("Unknown type")
    }
   }
   Type switches are a way to handle different types storage