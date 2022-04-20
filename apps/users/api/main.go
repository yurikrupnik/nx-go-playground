package main

import (
	"fmt"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

//func getUsers() {
//
//}

func main() {
	fmt.Println(Hello("users-api"))

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
