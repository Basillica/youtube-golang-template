package helpers

import (
	"os"
	"strconv"
)

// GetEnv demonstrates how to use the Add function.
//
//    fmt.Println(Add(1, 2))
//    // Output: 3
func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}


func GetEnvInt(key, defaultVal string) int {
	var v int
	var err error
	if value, exists := os.LookupEnv(key); exists {
		if v, err = strconv.Atoi(value); err != nil {
			return v
		}
		return -1
	}

	return -1
}

func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}