package sqllite_repository

import (
	"appointment/entity"
	"appointment/setting"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

func ScheduleAppointment(data entity.Appointment) (error){

	database,err:= sql.Open(setting.ServerSetting.DriverName,setting.ServerSetting.Database)
	if err!=nil {
		log.Fatalf(fmt.Sprintf("Unable to connect to sqlite3 database  due to the following error : %v ", err))
	}
	month:=time.Month(data.Month)
	avaialbleTime := time.Date(data.Year, month, data.Day, data.Hour, data.Minute, data.Second, data.NanoSecond, time.UTC)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS doctor (id INTEGER PRIMARY KEY, Available_time TEXT)")
	statement.Exec()
	statement,_ = database.Prepare("INSERT INTO doctor (Available_time) VALUES (?)")
	statement.Exec(avaialbleTime.String())
	return  nil
}

func BookApponitment(data entity.Appointment) (error){

	database,err:= sql.Open(setting.ServerSetting.DriverName,setting.ServerSetting.Database)
	if err!=nil {
		log.Fatalf(fmt.Sprintf("Unable to connect to sqlite3 database  due to the following error : %v ", err))
	}

	month:=time.Month(data.Month)
	avaialbleTime := time.Date(data.Year, month, data.Day, data.Hour, data.Minute, data.Second, data.NanoSecond, time.UTC)
	availableEndtime := avaialbleTime.Add(time.Minute*15)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS doctor_appointments (id INTEGER PRIMARY KEY, first_name NAME , last_name NAME , Available_time TEXT , Available_endTime TEXT)")
	statement.Exec()
	statement,_ = database.Prepare("INSERT INTO doctor_appointments (first_name , last_name , Available_time ,  Available_endTime ) VALUES (?, ?, ?, ?)")
	statement.Exec(data.FirstName, data.LastName , avaialbleTime.String(), availableEndtime.String())
	return  nil
}

func GetAppointments() ([]map[string]string, error) {
	var response  []map[string]string
	database,err:= sql.Open(setting.ServerSetting.DriverName,setting.ServerSetting.Database)
	if err!=nil {
		log.Fatalf(fmt.Sprintf("Unable to connect to sqlite3 database  due to the following error : %v ", err))
	}
	rows, _ := database.Query(`SELECT  Available_time , Available_endTime FROM doctor_appointments`)
	var Availabletime string
	var Available_endtime string
	for rows.Next() {
		result :=make(map[string]string)
		rows.Scan(&Availabletime, &Available_endtime)
		result["slot start time"] =  Availabletime
		result["slot end time"] = Available_endtime
		response = append(response,result)
	}

	return response , nil
}

func CancelAppointment(FirstName string) (error){

	database,err:= sql.Open(setting.ServerSetting.DriverName,setting.ServerSetting.Database)
	if err!=nil {
		log.Fatalf(fmt.Sprintf("Unable to connect to sqlite3 database  due to the following error : %v ", err))
	}
    _,_ = database.Exec(fmt.Sprintf(`DELETE from doctor_appointments where first_name = "%s"`,FirstName))
	return  nil
}
