package main

import (

	"os"
	"parcial/cmd/config"
	"parcial/cmd/handler"
	"parcial/cmd/middleware"
	"parcial/cmd/sever/external/database"
	"parcial/internal/odontologo"

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

	storage := database.SqlStore{DB: mysqlDb}

	pacientRepo := paciente.Repository{Store: &storage}
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

	odontologoRepo := odontologo.Repository{Store: &storage}
	dservice := odontologo.Service{Repository: &odontologoRepo}
	dhandler := handler.DentistHandler{IService: &dservice}

	router := gin.Default()
	dentists := router.Group("dentist", authMidd.AuthHeader)
	{
		dentists.GET(":id", dhandler.GetDentistById)
		dentists.POST("", dhandler.CreateDentist)
		dentists.DELETE(":id", dhandler.DeleteDentist)
		dentists.PUT(":id", dhandler.UpdateDentist)
	}
	router.Run()

}
