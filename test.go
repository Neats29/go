package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    name, _ := reader.ReadString('\n')
    fmt.Scanf(name)
}

//func ask(string name, int age, string favLang){

//}

