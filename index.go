package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID        int64  `gorm:"AUTO_INCREMENT;primary_key`
	Firstname string `gorm:"size:255`
	Lastname  string `gorm:"size:255`
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func konek() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@/suko_orm")
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"title": "Ini Judul",
		})
	})

	router.Use(Cors())

	v1 := router.Group("api/v1")
	{
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.POST("/users", PostUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
		v1.OPTIONS("/users", OptionsUser)     // POST
		v1.OPTIONS("/users/:id", OptionsUser) // PUT, DELETE
	}
	router.Run(":9090")
}

func PostUser(c *gin.Context) {
	db, _ := konek()
	cek := db.HasTable(&User{})
	if cek == false {
		db.CreateTable(&User{})
	}
	var user User
	c.Bind(&user)
	log.Println(&user)
	db.Create(&user)
}

func GetUsers(c *gin.Context) {
	var users []User
	db, _ := konek()
	getdata := db.Find(&users)

	if getdata != nil {
		c.JSON(200, users)
	} else {
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}

	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	db, _ := konek()
	data := db.Where("id = ?", id).Find(&user)
	if data != nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	db, _ := konek()
	db.Model(&user).Where("id = ?", id).Update(&user)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	db, _ := konek()
	db.Where("id = ?", id).Delete(User{})

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
