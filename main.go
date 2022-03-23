package main

import (
	"log"

	"github.com/wryonik/appointment/controllers"
	"github.com/wryonik/appointment/models"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://98e7c538041340539b730bdeb03ae775@o1176298.ingest.sentry.io/6273810",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/prescriptions", controllers.FindPrescriptions)
	r.POST("/prescriptions", controllers.CreatePrescription)
	r.PATCH("/prescriptions", controllers.UpdatePrescription)
	r.DELETE("/prescriptions", controllers.DeletePrescription)

	// Run the server
	r.Run()
}
