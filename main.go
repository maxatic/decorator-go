package main

import "fmt"

func decprocess(base func(output string)) func(string) {
	return func(output string) {
		fmt.Println("before decoration process \n")
		base(output)
		fmt.Println("after decoration process")
	}
}

func basic(output string) {
	fmt.Println(output)
}

func main() {
	basic("We're Humble \n")
	decprocess(basic)("We're Reactive \n")
}
