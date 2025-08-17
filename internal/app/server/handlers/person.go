package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/evertontomalok/go-rest-sample/internal/app/domain/entities"
	"github.com/evertontomalok/go-rest-sample/internal/ports"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type PersonHandler struct {
	repo ports.Repository
}

func (p *PersonHandler) GetPersonById(c *gin.Context) {
	personId, err := parsePersonId(c)
	if err != nil {
		return
	}
	person, found := p.repo.Get(personId)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "person id not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (p *PersonHandler) DeletePersonById(c *gin.Context) {
	personId, err := parsePersonId(c)
	if err != nil {
		return
	}

	if err := p.repo.Delete(personId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "person id not exist"})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (p *PersonHandler) UpdatePersonById(c *gin.Context) {
	personId, err := parsePersonId(c)
	if err != nil {
		return
	}

	var person entities.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		log.Error("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	person.ID = personId
	if err := p.repo.Update(person); err != nil {
		log.Error("Error upserting person:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (p *PersonHandler) InsertPerson(c *gin.Context) {
	var person entities.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		log.Error("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	personId, err := p.repo.Insert(person)
	if err != nil {
		log.Error("Error inserting person:", err)
		c.JSON(http.StatusConflict, gin.H{"error": "Person already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": personId})
}

func (p *PersonHandler) GetRoutes() []Route {
	var routes = []Route{
		{
			Path:    "/person",
			Method:  http.MethodPost,
			Handler: p.InsertPerson,
		},
		{
			Path:    "/person/:personId",
			Method:  http.MethodGet,
			Handler: p.GetPersonById,
		},
		{
			Path:    "/person/:personId",
			Method:  http.MethodPut,
			Handler: p.UpdatePersonById,
		},
		{
			Path:    "/person/:personId",
			Method:  http.MethodDelete,
			Handler: p.DeletePersonById,
		},
	}
	return routes
}

func parsePersonId(c *gin.Context) (int64, error) {
	personIdParam := c.Param("personId")
	personId, err := strconv.ParseInt(personIdParam, 10, 64)
	if err != nil {
		errMsg := fmt.Errorf("error converting string to integer: %s", err)
		log.Error(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return 0, errMsg
	}
	return personId, nil
}

func NewPersonHandler(repo ports.Repository) *PersonHandler {
	return &PersonHandler{repo: repo}
}
