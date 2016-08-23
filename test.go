package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("What is your age: ")
    //TODO: check for age to be no more than 3 digits
    age, _ := reader.ReadString('\n')
    fmt.Scanf(age)
    fmt.Print("What is your name: ")
    name, _ := reader.ReadString('\n')
    fmt.Scanf(name)
    fmt.Print("Your name is ",  name, " and you are ", age, " years old")
}



