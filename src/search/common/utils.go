package common

import (
	"fmt"
	"os"
)

// MustGetEnv 获取环境变量，如果没有设置则退出。
func MustGetEnv(envName string, defaultValue string) string {
	v := os.Getenv(envName)
	if v == "" {
		if defaultValue == "" {
			panic(fmt.Sprintf("environment variable %q not set", envName))
		}
		v = defaultValue
	}
	return v
}
