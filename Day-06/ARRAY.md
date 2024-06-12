What is Arrays in Go?
An array in Go is a numbered sequence of elements of a specific length and type.

Characteristics of Go arrays
1. Fixed Size: Once declared the size of the array cannot be changed.
2. Value Type: Arrays in Go are value types, not reference types. When you assign or pass an array to a new variable, it copies the entire array.
3. Zero value: Uninitialized elements of an array are set tot the zero value of the arrays type.

Different Ways to Declare Arrays in Go?
1. Declaration without Initialization
2. Declaration With Initialization
3. Shorthand Declarations
4. 
5. Initializing With Specific Elements
    arr:= [5]{1:5, 4:6}: the second and fifth elements is 5 and 6.

6. Using the ... operator to infer length: arr:= [...]int{1,2,3,4}
    Here the length of the array is inferred from the number of initialized elements. 
    The array arr will be of the length 5.

7. Arrays of Arrays(Multi-Dimensional arrays): var multiArr[4][5]
    Initializing Multi-Dimensional arrays:
    multiArr:= [2][3]int{1,2},{3,4,5}

8. Arrays with pointers: var arr[5]*int: This declares an array of pointer to integers
9. Arrays of Structs:
type Person struct{
		Name string
		Age int
}
var people [2]Person


Common Operations of Arrays:
1. Length of an Array:
    To get the number of elements in an array use the built-in len() function.
    var myArr [5]int
    length := len(myArr)

2. Copying an Array:
   arrays are value types in Go, so assigning one array to another creates a copy. copyArr ;= myArr
   var myArr = [3]int {1,2,3}
   var copyArray = myArray ->this creates a copy of myArr

3. Modifying elements:
   To change the value of an element access the element by its index and assign a new value
   myArr[0]=100 ->change the first element to 100

4. Passing Arrays to Functions
   We can pass arrays to functions but remember that this will pass a copy of the array, not a referece .
   func processArray(arr [5]int){
    <!--Function code here-->
   }

5. Comparing Arrays
   We can compare two arrays of the same type and length using the == operator
   var arr1 = [3]int{1,2,3}
   var arr2 = [3]int{1,2,3}

   isEqual := arr1==arr2 ->isEqual is true


6. Iterating over Arrays
   1. For Loop: Traditional for loop or for-range loop can be used
   numbers := [5]int{1,2,3,4,5}
   for i:=0; i<len(numbers);i++ {
    fmt.Println(numbers[i])
   }
   Use "len" function to get the length of arrays

   2. Range Loop: More idiomatic in Go
   numbers := [5]int{1,2,3,4,5}
   for index, value := range numbers{
    fmt.Println(index,value)
   }


What are Multidimensional Arrays:
A multidimensional array is essentially an array of arrays.The most common type is a 2D array, often used to represent matrices or grids.

var matrix [3][3] int -> represents it has three arrays and each array contains three elements.