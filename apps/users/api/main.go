package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	//"@nx-go-playground/my_lib"
)

var mongoUrl string = "mongodb+srv://yurikrupnik:T4eXKj1RBI4VnszC@cluster0.rdmew.mongodb.net/"

func mongoDb() {
	/*
	   Connect to my cluster
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))

	//client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//err = client.Connect(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	log.Println(">>>>>>>>>>all was good")
	//}
	//defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}

func users() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := gin.New()
		r.GET("/ping", func(c *gin.Context) {
			//fmt.Println(c)
			//c.Error("my error")
			//c.AbortWithStatusJSON(400, {"aris": true})
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}

func Hello(name string) string {
	result := "Hello " + name
	//res := my_lib.MyLib("aris")
	return result
}

func getUsers() {
	//return make([]{})
}

//func Logger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t := time.Now()
//
//		// Set example variable
//		c.Set("example", "12345")
//
//		// before request
//
//		c.Next()
//
//		// after request
//		latency := time.Since(t)
//		log.Print("latency", latency)
//
//		// access the status we are sending
//		status := c.Writer.Status()
//		log.Println("<<< My statuses", status)
//	}
//}

func main() {
	// Force log's color
	//mongoDb()
	//gin.ForceConsoleColor()
	r := gin.Default()
	//gin.RouterGroup{}
	//r.Use(Logger())
	//r.Use(users())
	//r.
	//r.Handlers
	//r.GET("/ping", func(c *gin.Context) {
	//	fmt.Println(c)
	//	//c.Error("my error")
	//	//c.AbortWithStatusJSON(400, {"aris": true})
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)
		fmt.Println("Time is: ", time.Now())
		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	r.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
		shit := c.Query("shit")         // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s %s %s", firstname, lastname, lastname, shit)
	})
	r.POST("/users", func(c *gin.Context) {
		//body := c.Request.Body
		jsonData, err := c.GetRawData()
		bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
		jsonBody := string(bodyAsByteArray)
		if err != nil {
			//Handle Error
		}
		log.Println("jsonData: ", jsonData)
		log.Println("jsonBody: ", jsonBody)
		//err = client.Set("id", jsonData, 0).Err()
		//ra := body.Read()
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.String(http.StatusOK, "Hello %s %s %s %s", id, page, name, message)
	})

	r.GET("/", func(c *gin.Context) {
		// If you set TrustedPlatform, ClientIP() will resolve the
		// corresponding header and return IP directly
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
		fmt.Printf("Request url: %s\n", c.Request.URL.Path)
		fmt.Printf("Body: %s\n", c.Request.Body)
	})
	r.POST("/", func(c *gin.Context) {
		// If you set TrustedPlatform, ClientIP() will resolve the
		// corresponding header and return IP directly
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
		fmt.Printf("Request url: %s\n", c.Request.URL.Path)
		fmt.Printf("Body: %s\n", c.Request.Body)
		bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
		jsonBody := string(bodyAsByteArray)
		log.Println("jsonBody: ", jsonBody)
		//log.Println("jsonBody: ", bodyAsByteArray["name"])
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
