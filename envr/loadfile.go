package envr

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
)

// LoadEnv loads env file from directory. Add env file containing folder and file name
func LoadEnv(envDirectory string, filenames ...string) error {
	if len(envDirectory) < 1 {
		return errors.New("environment directory not supplied")
	}

	var f []string
	for _, file := range filenames {
		file = envDirectory + "/" + file // building the directory path
		f = append(f, file)
		fmt.Println(file)
	}

	// loads environment files from  director
	err := godotenv.Load(f...)
	//err := godotenv.Load(fmt.Sprintf("%v/%v", envDirectory, filename))
	if err != nil {
		return err
	}

	return nil
}
