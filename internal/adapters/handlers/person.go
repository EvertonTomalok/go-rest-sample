package handlers

import (
	"net/http"
	"strconv"

	"github.com/evertontomalok/go-rest-sample/internal/app/domain/entities"
	"github.com/evertontomalok/go-rest-sample/internal/ports"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func NewPersonHandler(repo ports.Repository) *PersonHandler {
	return &PersonHandler{repo: repo}
}

type PersonHandler struct {
	repo ports.Repository
}

func (p *PersonHandler) GetPersonById(c *gin.Context) {
	personIdParam := c.Param("personId")
	personId, err := strconv.ParseInt(personIdParam, 10, 64)
	if err != nil {
		log.Error("Error converting string to integer:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	log.Infof("Fetching person id: %d", personId)
	person, found := p.repo.Get(personId)
	if !found {
		c.JSON(http.StatusNotFound, map[string]string{"error": "person id not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (p *PersonHandler) UpdatePerson(c *gin.Context) {
	var person entities.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		log.Error("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := p.repo.Update(person.ID, person); err != nil {
		log.Error("Error upserting person:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upsert person"})
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

	if err := p.repo.Insert(person.ID, person); err != nil {
		log.Error("Error inserting person:", err)
		c.JSON(http.StatusConflict, gin.H{"error": "Person already exists"})
		return
	}

	c.JSON(http.StatusCreated, person)
}

func (p *PersonHandler) GetPersonRoutes() []entities.Route {
	var routes = []entities.Route{
		{
			Path:    "/person/:personId",
			Method:  http.MethodGet,
			Handler: p.GetPersonById,
		},
		{
			Path:    "/person",
			Method:  http.MethodPost,
			Handler: p.InsertPerson,
		},
		{
			Path:    "/person",
			Method:  http.MethodPut,
			Handler: p.UpdatePerson,
		},
	}
	return routes
}
