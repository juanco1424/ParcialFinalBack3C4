package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Ejemplo de implementacion local probar de nuevo con integracion de DB
	/*env := os.Getenv("ENV")
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

	mysqlDb, err := database.NewMySQLDatabase()
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
	router.Run()*/
}
