package envr

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
)

// LoadEnv loads env file from directory. Add env file containing folder and file name
func LoadEnv(envDirectory string, filename ...string) error {
	if len(envDirectory) < 1 {
		return errors.New("environment directory not supplied")
	}

	// loads environment files from  director
	err := godotenv.Load(fmt.Sprintf("%v/%v", envDirectory, filename))
	if err != nil {
		return err
	}

	return nil
}
