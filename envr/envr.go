package envr

import (
	"errors"
	"os"
	"strconv"
)

// GetInt64  load from env file return as int64
func GetInt64(key string) (int64, error) {
	resultStr := os.Getenv(key)

	if len(resultStr) < 1 {
		return 0, errors.New("empty environment value or key not exists")
	}

	result64, err := strconv.ParseInt(resultStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return result64, nil
}

// GetInt32  load from env file return as int32
func GetInt32(key string) (int32, error) {
	resultStr := os.Getenv(key)

	if len(resultStr) < 1 {
		return 0, errors.New("empty environment value or key not exists")
	}

	result64, err := strconv.ParseInt(resultStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(result64), nil
}

// GetInt  load from env file return as int
func GetInt(key string) (int, error) {
	resultStr := os.Getenv(key)

	if len(resultStr) < 1 {
		return 0, errors.New("empty environment value or key not exists")
	}

	result64, err := strconv.ParseInt(resultStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(result64), nil
}

// GetString  load from env file return as string
func GetString(key string) (string, error) {
	resultStr := os.Getenv(key)
	if len(resultStr) < 1 {
		return "", errors.New("empty environment value or key not exists")
	}
	return resultStr, nil
}
