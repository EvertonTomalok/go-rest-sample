package handlers

import (
	"net/http"
	"strconv"

	"github.com/evertontomalok/go-rest-sample/internal/app/domain/entities"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetPersonById(c *gin.Context) {

	personIdParam := c.Param("personId")
	personId, err := strconv.ParseInt(personIdParam, 10, 64)
	if err != nil {
		log.Error("Error converting string to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	log.Infof("Fetching person id: %d", personId)
	person := entities.Person{
		ID:   personId,
		Name: "My Name",
		Age:  20,
	}
	c.JSON(http.StatusOK, person)
}
