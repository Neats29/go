package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var age string
    var name string
    fmt.Print("What is your age: ")
    age = read(age)
    for age > "130" {
      fmt.Print("Age invalid, please enter your age...")
      age = read(age)
    }
    fmt.Scanf(age)
  
    fmt.Print("What is your name: ")
    name = read(name)
    fmt.Scanf(name)
    fmt.Print("Your name is ",  name, " and you are ", age, " years old")
}


func read(answer string) string {
    reader := bufio.NewReader(os.Stdin)
    answer, _ = reader.ReadString('\n')
    return answer
}



