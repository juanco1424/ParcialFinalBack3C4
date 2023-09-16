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
	//Ejemplo de implementacion local probar de nuevo con integracion de DB
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
	odontologoRepo := odontologo.Repository{Store: &storage}
	dservice := odontologo.Service{Repository: &odontologoRepo}
	phandler := handler.DentistHandler{IService: &dservice}

	router := gin.Default()
	dentists := router.Group("dentist", authMidd.AuthHeader)
	{
		dentists.GET(":id", phandler.GetDentistById)
		dentists.POST("", phandler.CreateDentist)
		dentists.DELETE(":id", phandler.DeleteDentist)
		dentists.PUT(":id", phandler.UpdateDentist)
	}
	router.Run()
}
