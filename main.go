package main

import "fmt"

func main() {
	greeting := "hello "
	name := "bashir"
	age := 24
	cousre := "golang"
	calling(greeting, name, cousre, age)
}

func calling(greeting, name, course string, age int) {
	fmt.Println(greeting, "my name is", name, "and my age is", age, "i am learning", course)
}
