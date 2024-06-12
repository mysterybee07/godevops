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
   
2. Appending to Slices:
   
3. Slicing a Slice
   
4. Delete an element from slice
   
5. Checking if a Slice is Empty
   
6. Sorting a Slice
   
7. Finding an Element in a Slice