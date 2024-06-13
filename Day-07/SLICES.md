What is Slices in Go?
Slices are a key data tupe in Go, providing a more flexible, dynamic, and powerful alternative to arrays

Characteristics in Go Slices
1. Dynamic Size: A slice is a dynamically sized, flexible view into the elements of an array so Unlike  arrays, slices are dynamic and can grow or shrink.
2. Reference Type: Unlike arrays, slices are reference types. When you pass a slice to a function, any modifications to the slice element will lbe visible outside the function.


Different Ways to Declare Slices in Go:

1. Using the var keyword without initializing it
   var slice [] int
   This declares a slice without initializing it. The slice defaults to nil.

2. Using a Slice Literal:
   mySlice := []int{1, 2, 3}
   This is concise way to create and initialize a slice with specific elements

3. Using the make function:
   mySlice := make([]int, 5) -> length and capacity are 5
   mySlice:= make([]int, 5, 10) ->length is 5 and capacity is 10
   The make function creates a slice with an initial length and capacity. This is useful when you know the number of elements the slice will hold.

4. From an Array:
   arr := [5]int{1,2,3,4}
   mySlice := arr[1:4] ->creates a slice including elements at index 1,2 and 3
   You can create a slice by slicing an array. This creates a slice that reference the array.


Common Operations in an array
1. Length and Capacity:
   Length is the number of elements in the slice. Capacity is the number of elements the slice can grow to.
   slice := [] int{1,2,3,4}
   length := len(slice)
   capacity := cap(slice)
   
2. Appending to Slices:
   var numbers [] int
   numbers = append(numbers, 1) ->appending single elements
   numbers = append(numbers, 2,3,4,5)

   moreNumbers := []int{6,7,8,9,10,11}
   numbers = append(numbers, moreNumbers...)
    
   
3. Slicing a Slice
   subSlice := numbers[1:4] ->here index 1 is inclusive but 4 is exclusive

   <!-- slicing upto specific index -->
   firstThree:=numbers[:3]

   Slicing from specific index to the end
   fromThree:=numbers[2:]
   
4. Delete an element from slice
   <!-- Removing the element at index 2 -->
   numbers = append(numbers[:2], numbers[:3]...)
   
5. Checking if a Slice is Empty
   if len(numbers) == 0{
      fmt.Println("Slice is empty")
   }
   
6. Sorting a Slice
   import"sort"
   <!-- sorting a slice of integers -->
   sort.Ints(numbers)
   <!-- Sorting a slice of strings -->
   sort.Strings(names)
   
7. Finding an Element in a Slice
   for _,v :=range numbers{
      if v==target{
         fmt.Println("Element found")
      }
   }

Iterating through the slices
1. For Loop: Traditional for loop or for-range loop can be used
   numbers := []int{1,2,3,4,5,6,7}
   for i:=0; i<len(numbers); i++{
      fmt.Println(numbers[i])
   }
   Use "len()" function to get the length of arrays

2. Range Loop: More idiomatic in Go
   numbers :=[]int{1,2,3,4}
   for index,value :range numbers{
      fmt.Println(index,value)
   }