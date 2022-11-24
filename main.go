package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id    int64
	Name  string `json:"student_name"`
	Age   int
	Level int
}

type Tabler interface {
	TableName() string
}

func (Student) TableName() string {
	return "student"
}
func main() {
	db := setDbConfigs()
	setRouterConfigs(db)
}

func setDbConfigs() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func setRouterConfigs(db *gorm.DB) {
	router := gin.Default()
	router.GET("/students", getAllStudents(db))
	router.Run("Localhost:8080")
}

func getAllStudents(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, getALlStudentsFromDb(db))
	}
	return gin.HandlerFunc(fn)
}

func getALlStudentsFromDb(db *gorm.DB) []Student {
	var students []Student
	result := db.Find(&students)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return students
}
