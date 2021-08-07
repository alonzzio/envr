package main

import (
	"envr"
	"fmt"
)

const myEnvWD = "/env"

func main() {

	//wd, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err) // error handle
	//}
	//parent := filepath.Dir(wd)
	//envPath := parent + myEnvWD
	//err := envr.LoadEnv("", "configurations.env")
	//if err != nil {
	//	fmt.Println("fhfhfh")
	//	log.Println(err)
	//}
	//
	//fmt.Println("ohohoh")
	//myStrVal, _ := envr.GetString("JWTSecretKey")
	//// use it
	//fmt.Println(myStrVal)
	////fmt.Println(envPath)
	fmt.Println("hi")
	envr.TestFunc()
}
