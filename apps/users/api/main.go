package main

import (
	"fmt"
)

func Hello(name string) string {
	return "Hello " + name
	//return result
}

//func getUsers() {
//
//}

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	fmt.Println(Hello("users-api"))
	//converter := typescriptify.New().
	//	Add(Person{})
	//err := converter.ConvertToFile("ts/models.ts")
	//if err != nil {
	//	panic(err.Error())
	//}
	//router := gin.Default()
	//router.GET("/users", getUsers)
	//// router.POST("/somePost", posting)
	//// router.PUT("/somePut", putting)
	//// router.DELETE("/someDelete", deleting)
	//// router.PATCH("/somePatch", patching)
	//// router.HEAD("/someHead", head)
	//// router.OPTIONS("/someOptions", options)
	//
	//// By default it serves on :8080 unless a
	//// PORT environment variable was defined.
	//router.Run()
}
