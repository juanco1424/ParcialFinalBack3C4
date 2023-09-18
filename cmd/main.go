package main

import (
	"os"
	"parcial/cmd/config"
	"parcial/cmd/handler"
	"parcial/cmd/middleware"
	"parcial/cmd/server/external/database"
	"parcial/internal/appointments"
	"parcial/internal/odontologo"
	"parcial/internal/paciente"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("ENV")

	if env == "" {
		env = "local"
	}

	if env == "local" {
		err := godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	}

	config, err := config.NewConfig(env)
	if err != nil {
		panic(err)
	}
	authMidd := middleware.NewAuth(config.PublicConfig.PublicKey, config.PrivateConfig.SecretKey)

	mysqlDb, err := database.MySQLDatabase()

	if err != nil {
		panic(err)
	}

	storageDentist := database.SqlStoreDentist{DB: mysqlDb}
	storagePatient := database.SqlStorePatient{DB: mysqlDb}
	storageAppointment := database.SqlStoreAppointment{DB: mysqlDb}

	pacientRepo := paciente.Repository{Store: &storagePatient}
	pservice := paciente.Service{Repository: &pacientRepo}
	phandler := handler.PatientHandler{PService: &pservice}

	router := gin.Default()
	patients := router.Group("patient", authMidd.AuthHeader)
	{
		patients.GET(":id", phandler.GetPatientById)
		patients.POST("", phandler.CreatePatient)
		patients.DELETE(":id", phandler.DeletePatient)
		patients.PUT(":id", phandler.UpdatePatient)
	}

	odontologoRepo := odontologo.Repository{Store: &storageDentist}
	dservice := odontologo.Service{Repository: &odontologoRepo}
	dhandler := handler.DentistHandler{IService: &dservice}

	dentists := router.Group("dentist", authMidd.AuthHeader)
	{
		dentists.GET(":id", dhandler.GetDentistById)
		dentists.POST("", dhandler.CreateDentist)
		dentists.DELETE(":id", dhandler.DeleteDentist)
		dentists.PUT(":id", dhandler.UpdateDentist)
	}

	appointmentRepo := appointments.AppointmentRepository{Store: &storageAppointment}
	aservice := appointments.AppointmentService{Repository: &appointmentRepo}
	ahandler := handler.AppointmentHandler{AppointmentService: &aservice}

	appointments := router.Group("appointments", authMidd.AuthHeader)
	{
		appointments.GET("", ahandler.GetAllAppointments)
		appointments.GET(":id", ahandler.GetAppointmentById)
		appointments.GET("patient/:dni", ahandler.GetAppointmentsByDni)
		appointments.POST("", ahandler.CreateAppointment)
		appointments.DELETE(":id", ahandler.DeleteAppointment)
		appointments.PUT(":id", ahandler.UpdateAppointment)
	}
	router.Run()

}
