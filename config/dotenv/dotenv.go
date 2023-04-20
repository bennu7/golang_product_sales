package dotenv

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadConfigENV(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		fmt.Println("err load config ", err)
	}
}

func GetKeyString(key string) string {
	return os.Getenv(key)
}
