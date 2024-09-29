package teamspeak3

import "os"

func getEnv(key string, fallback string) string {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return fallback
}
