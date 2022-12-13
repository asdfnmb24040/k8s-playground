package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	loadEnvSuccess = InitEnv()
	envFileName    = ".env"
)

func ToString(val int) string {
	return strconv.Itoa(val)
}

func FixStringLen(str string) string {
	limit := 500

	if len(str) > limit {
		return str[0:limit] + "....more"
	}
	return str
}


func InitEnv() bool {
	loadEnvErr := godotenv.Load("./" + envFileName)

	if loadEnvErr != nil {
		PrintObj("load env err. try another locaion")
		loadEnvErr = godotenv.Load("../" + envFileName) // only test use

		if loadEnvErr != nil {
			PrintObj("load env err. fail")
			return false
		}
	}

	PrintObj("load env suceess")
	return true
}

func GetEnv(str string) string {
	if !loadEnvSuccess {
		PrintObj("load env err pls check the .env", "GetEnv err")
	}

	val := os.Getenv(str)
	// PrintObj(val, "GetEnv "+str)
	return val
}


func PrintObj(obj interface{}, params ...string) {
	// print
	json, _ := json.Marshal(obj)
	key := ""

	if len(params) == 1 {
		if params[0] != "" {
			key = params[0]
			fmt.Println("=== " + key + " ===")
		}
	}

	if obj != "" {
		fmt.Println(string(json))
	}
}