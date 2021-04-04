package main

import (
	"appointment/entity"
	"appointment/setting"
	"appointment/sqllite-repository"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"strconv"
)

func  init(){
	setting.Setup()
}

func main() {

	r := gin.New()
	r = gin.Default()
	r.POST("/schedule", func(c *gin.Context) {
		var data entity.Appointment

		err := c.ShouldBindWith(&data, binding.JSON)
		if err != nil {
			log.Println(err)
			c.String(500,"invalid input")
		}

		err = sqllite_repository.ScheduleAppointment(data)
		if err!=nil{
			log.Println(err)
			c.String(500,"Couldnt log available time")
		}

		c.String(200,"Successfully logged available time ")
	})

	r.POST("/book", func(c *gin.Context) {
		var data entity.Appointment

		err := c.ShouldBindWith(&data, binding.JSON)
		if err != nil {
			log.Println(err)
			c.String(500,"invalid input")
		}

		err = sqllite_repository.BookApponitment(data)
		if err != nil{
			log.Println(err)
			c.String(500,"Couldnt book appointment")
		}

		c.String(200,"Successfully booked appointment ")
	})

	r.GET("/list", func(c *gin.Context) {

		response ,err := sqllite_repository.GetAppointments()
		if err != nil{
			log.Println(err)
			c.String(500,"Couldnt fetch list of appointments")
		}

		c.JSON(200,response)
	})

	r.DELETE("/cancel", func(c *gin.Context) {
		var data entity.Appointment

		err := c.ShouldBindWith(&data, binding.JSON)
		if err != nil {
			log.Println(err)
			c.String(500,"invalid input")
		}

		err = sqllite_repository.CancelAppointment(data.FirstName)
		if err != nil{
			log.Println(err)
			c.String(500,"Couldnt cancel appointment")
		}
		c.String(200,"Successfully cancelled the appointment")
	})

	r.Run(":" +  strconv.Itoa(setting.ServerSetting.HttpPort))
	}

