# What is Pointers in Go?
A Pointer is a variable that stores the memory address of another variable.

# Characteristics of Go Pointers
1. Efficiency: Using pointers can be more efficient for large structs or arrays, as it avoids copying large amounts of data.
2. Mutability: Pointers allow you to modify the original variable, which is useful in functions where you want to change the input data.
3. Passing Pointers: We can pass pointers to functions to modify the original data or to improve performance by avoiding data copying
4. Type Specific: Pointers in Go are type-specific. A pointer to an int is different from a poijnter to a string.
5. Using & operator: The & operator is used to get the address of a variable. For example, p:= &x assigns the address of x to p
6. Using * operator: The * operator is used to access the value at the address pointed to by the pointer. For example, *p accesses the value pointed by p.


# Different Way to Declare Pointers in Go

1. Using var
   var p *int
   To declare a pointer, you use the * followed by the type of the stored value. The zero value of a pointer is nil 

2. Pointer to an Existing Variable
   var x int = 10
   p:= &x
   We can assign a pointer to the address of an existing variable using the &(addres of) operator.

3. Creating a Pointer with new:
   p:= new(int)
   The new function allocates a memory and returns a pointer to it. This is useful when you want to allocate memory for a variable directly without declaring it first.

4. Pointer to Pointer
   var x int
   var p *int=&x
   var pp **int=&p
   We can have a pointer to another pointer, allowing for multiple levels of indirections

# Common Operations in Pointers
1. Dereferencing a Pointer
   var x int=10
   var p *int= &x
   fmt.Println(*p)
   To access the value that a pointer points to, you use the * operator. This is known as dereferencing a pointer.

2. Changing the value at a Pointer
   *p=20 
   <!--changes the value of x to 20 -->
   We can change the value at the memeory address a pointer is pointing to by dereferencing the pointer and assigning a new value

3. Comparing pointers
   var a, b int
   p1:=&a
   p2:=&b
   p3:=&a
   fmt.Println(p1==p2)
   fmt.Println(p1==p3)
   We can compare two pointers using the == and != operators to check if they point to the same address

4. Passing pointers to functions
   func increment(x *int){
    *x++
   }

   var value int =5
   increment(&value)
   fmt.Println(value)
   We can pass pointers to function. This is often done to modify the value if the argument or to pass large structs without copying them.
