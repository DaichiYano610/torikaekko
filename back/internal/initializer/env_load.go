package initializer

import (
	//"fmt"
	"log"
	//"os"

	"github.com/joho/godotenv"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/*
func main() {
	Env_load()
	message := fmt.Sprintf("db_host=%s db_user=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"))
	fmt.Println(message)
}
*/
