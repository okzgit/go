package main

import "fmt"

func main() {
    // declare variables 
    var a, b int
        
    //assign values
    a = 15
    b = 10
    
    //arithmetic operation
    
    //ADDITION
    c := a + b
    fmt.Printf("%d + %d = %d \n",a, b,c)
    
    //SUBTRACTION
    d := a - b
    fmt.Printf("%d - %d = %d \n",a, b,d)
    
    // DIVISION
    e := float32(a) / float32(b)
    fmt.Printf("%d / %d = %.2f \n",a,b,e)
    
    // MULTIPLICATION
    f := a * b
    fmt.Printf("%d * %d = %d \n",a,b,f)
}