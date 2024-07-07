# What is Structs in Go?
A struct in Go is a composite data type that groups together variables under a single name. These variables are known as fields.

# Characteristics of Struct in Go
1. Data modeling: Structs are extensively used to model real-world data and entities in applications.
2. Configuration: Structs are often used to hold configuration data.
3. Design Patterns: Common design patterns like Decorator, Adapter, and Composite can be implemented using structs.
4. Fields: Structs are composed of fields, each of which has a name and a type. Fields can be of any type including other structs, slices, maps, and interfaces.
5. Methods: We can define methods on structs. This allows structs to have behaviours associated with them.
6. Tags: Struct fields can have tags, which are string literals associated with the field for metadata purposes.

# Different Ways to Declare Structs in Go
1. Basic Declarations: 
   type Person struct {
    Name string
    age int
   }

2. Declaring and instantiating together
   var person Person = Person{
    Name: "John",
    Age: 10,
   }
   //or using the shorthand
   person:= Person{
    "Alice", 30
   }

3. Using the new Keyword
   personPtr:=new(Person)
   personPtr.Name="Bob"
   personPtr.Age = 25

4. Anonymous Structs
   var person = struct{
    Name string
    Age int
   }{
    Name: "Charlie",
    Age:40,
   }
    For a one-off struct that doesnot need to be reused, you can declare an anonymous struct

5. Nested Structs
   type Address struct{
    City, State string
   }
   type Person struct{
    Name string
    Age int
    Address Address
   }

6. Embedded (Anonymous Fields)
   type Contact struct{
    Phone, Email string
   }
   type Person struct{
    Name string
    Age int
    Contact //Embedded field
   }

# Common Operations in Go structure
1. Creating an Instance of a Struct
   type Person struct{
    Name string
    Age int
   }
   person1:= Person{
    Name: "Biraj",
    Age: 23,
   }

2. Accessing and Modifying fields
   fmt.Println(person1.Name) //access
   p.Age= 24 //modify

3. Pointers to Structs
   p:= &Person{
    Name:"Biraj",
    Age: 25,
   }

4. Methods on Structs
   func(p Person) Greet()string{
    return "Hello," + p.Name
   }
   We can define methods on structs. Methods have a receiver argument
   that represents teh instance of the struct.

5. Tagging Struct Fields
   type Person struct{
    Name string `json: "name"`
    Age int `json:"age"`
   }
   We can tag struct fields with additional metadata, often used with encoding/decoding libraries.