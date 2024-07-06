What is Maps in Go?
Maps in Go are a key data structure used for storing and retrieving data efficiently. They are similar to dictionaries and hash tables in other programming languages.

Characteristics of Go Maps
1. Efficiency: Maps provide efficient lookup, insertion and deletion operations.
2. flexibility: They can be used to represent complex data structures.
3. Dynamic sizing: Maps grow as needed, making them flexile for data that varies in size.
4. Reference Type: Maps are reference types. Passing a map to a function allows the function to modify the original map.

Different Ways to Declare Maps in Go
1. Using the var keyword without initialzing it
   var m map [keytype]valueType
   e.g. var m map [string]int

2. Using a map literal
   m:=map[string]int{
   "apple":5;
   "ball":6;
   }

3. Using the make Function
   m := make(map[string]int)
   <!-- This make function is the most common way to create a map. It initializes the map and allocates memory for it -->

4. Using the make Function with Specific size
   m := make(map[keyType]valueType, initialSize)
   <!-- If you know the number of element that will be inserted into the map, you can specify an initial size to optimize memory allocation -->

5. Using structs as Map Values
   Maps can also have struct types as values, allowing for more complex data structures:
   type Product struct{
    Name string
    Price float64
   }
   products:= make(map[int]Product)

6. Nested Maps
   company := make(map[string] map[int] string) -> first one is a map with key string and value of type map
   This creates a map with string keys, where each key maps to another map with integer keys and string values.

#  Common Operations Using map

1. Check Length
   m:= map[keytype] valueType{key1: value1, key2: value2}
   len(m)

2. Adding Elements
   mapVariable[key]=value

3. Retrieving Elements
   value := mapVariable[key]
   If you access a key that doesnot exist, you get the zero value of the value type

4. Checking existence
   value, ok:= mapVariable[key]
   ok is true if the key exists false otherwise.

5. Deleting Elements
   delete(mapVariable, key)

6. Iterating Over Maps
   for key, value := range mapVariable{
      <!-- Code -->
   }