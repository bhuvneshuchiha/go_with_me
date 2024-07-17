package main

import (
	"net/http"
	"rest_api/db"
	"rest_api/models"

	"github.com/gin-gonic/gin"
)


func main() {
    db.InitDB()
    server := gin.Default()

    server.GET("/events", getEvents)

    server.Run(":8080") //Local host
}

func getEvents(context *gin.Context) {
    events, err := models.GetAllEvents()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message" : "Try again later"})
        return
    }
    context.JSON(http.StatusOK, events)
}

func CreateEvent(context *gin.Context) {
    var event models.Event
    err := context .ShouldBindJSON(&event)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"Message" : "Could not parse"})
    }
    event.ID = 1
    event.UserID = 1
    err = event.Save()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not create events. Try again later"})
        return
    }
    context.JSON(http.StatusCreated, gin.H{"Message" : "Event created",
    "event" : event})
}



















