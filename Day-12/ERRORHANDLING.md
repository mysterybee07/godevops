# How to handle errors in Go?
Unlike many modern programming languages that use exceptions for error handling, GO uses a different approach that relies on returning error values

# The error Type:
1. Go has a built-in interface named error.
2. Its a simple interface with a single method: Error() string.
3. An error variable represents any value that can describe itself as a string.

# Returning Errors:
1. Functions that can fail, return an error value as their last return value.
2. If a function succeeds, the error is nil, otherwise its not-nil

# Error handling in GO
1. Return Error:
   func divide(a, b float64)(float64, error){
    if b==0{
        return 0, errors.New("division by zero")
    }
    return a/b, nil
   }

2. Checking errors:
   result, err := divide(10,0)
   if err != nil{
    fmt.Println(err)
    return
   }
   fmt.Println(result)
   When you call a function that returns an error, you should check the error before proceeding

3. Custom Error Types:
   type MyError struct{
    Msg string
    Code int
   }
   func(e *MyError)Error()string{
    return fmt.Sprintf("%d-%s", e.Code, e.Msg)
   }
   We can define custom error types implementing the error interface.
   This is useful for providing more context or handling specific kinds of errors differently.

4. Panic and Recover:
   func riskyFunction(){
    defer func(){
        if r:recover(); r!=nil{
            fmt.Println("Recovered from:", r)
        }
    }()
    panic("Something bad happened")
   }
   panic is a built-in function that stops the ordinary flow of control and begins panicking.

   recover is a built-in function that regains control that regains control of a panicking goroutines.

   Use panic and recover sparingly there more akin to exceptions but generally used to generally unrecover errors and bug.