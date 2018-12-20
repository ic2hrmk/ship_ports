package port

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

//
// Default env. file for testing
//
const defaultEnvFile = "test.env"

//
// Env. vars names
//
const (
	portGatewayURL       = "PORT_GATEWAY_URL"
	inconsistentFilePath = "INCONSISTENT_IMPORT_FILE_PATH"
	correctFilePath      = "CORRECT_IMPORT_FILE_PATH"
)

func init() {
	var envFile string

	flag.StringVar(&envFile, "env", defaultEnvFile, "Env. file")
	flag.Parse()

	if err := godotenv.Load(defaultEnvFile); err != nil {
		log.Fatal(err)
	}

	log.Println("ENV FILE LOADED:", envFile)
}
