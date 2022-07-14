package main

import "fmt"

func Hello(name string) string {
	result := "Hellso " + name
	return result
}

func main() {
	fmt.Println(Hello("infra-commands"))
}
