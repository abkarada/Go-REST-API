package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Person struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	RandomVal string `json:"verfication_code"`
}

var Mails = []Person{
	{ID: 1, Username: "John", Email: "john@gmail.com"},
	{ID: 2, Username: "Jane", Email: "jane@gmail.com"},
	{ID: 3, Username: "Doe", Email: "done@gmail.com"},
}

func getAllPersons(c *gin.Context) {
	c.JSON(http.StatusOK, Mails)
}

func getPersonByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid ID format"})
		return
	}
	for _, person := range Mails {
		if person.ID == idInt {
			c.JSON(http.StatusOK, person)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

func createPerson(c *gin.Context) {
	var newPerson Person
	if err := c.ShouldBind(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
	}
	newPerson.ID = len(Mails) + 1
	Mails = append(Mails, newPerson)
	c.JSON(http.StatusCreated, newPerson)
}

func main() {
	server := gin.Default()
	server.GET("/persons", getAllPersons)
	server.GET("/persons/:id", getPersonByID)
	server.POST("/persons", createPerson)
	server.Run("localhost:8080")
}
