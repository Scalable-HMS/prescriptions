package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wryonik/appointment/controllers"
	"github.com/wryonik/appointment/models"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Role  string `json:"given_name"`
	Email string `json:"email"`
	Id    string `json:"nickname"`
}

func authMid(c *gin.Context) {

	url := "https://dev-rgmfg73e.us.auth0.com/userinfo"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", c.Request.Header["Authorization"][0])

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := Response{}
	json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Email)
	fmt.Println(response.Role)
	fmt.Println(response.Id)
	c.Params = []gin.Param{
		{
			Key:   "email",
			Value: response.Email,
		},
		{
			Key:   "role",
			Value: response.Role,
		},
		{
			Key:   "id",
			Value: response.Id,
		},
	}
}

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

	secureGroup := r.Group("/secure/", authMid)

	// Routes
	secureGroup.GET("/prescriptions", controllers.FindPrescriptions)
	secureGroup.POST("/prescriptions", controllers.CreatePrescription)
	secureGroup.PATCH("/prescriptions", controllers.UpdatePrescription)
	secureGroup.DELETE("/prescriptions", controllers.DeletePrescription)

	// Run the server
	r.Run()
}
