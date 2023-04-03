package common

import (
	"fmt"
	"io/ioutil"
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

// MustLoadSecretAsString 读取secret，首先从环境变量中读取；如果没有设置，则从secret文件环境变量指定的文件中读取（考虑Kubernetes部署的场景）。
func MustLoadSecretAsString(secretEnvName string, secretFileNameEnvName string) string {
	v := os.Getenv(secretEnvName)
	if v != "" {
		return v
	}

	secretFileName := os.Getenv(secretFileNameEnvName)
	if secretFileName != "" {
		rawBytes, err := ioutil.ReadFile(secretFileName)
		if err != nil {
			fmt.Printf("failed to read secret file %s, %v", secretFileName, err)
			panic(fmt.Sprintf("failed to read secret %s/%s", secretEnvName, secretFileNameEnvName))
		}
		return string(rawBytes)
	}
	panic(fmt.Sprintf("failed to read secret %s/%s", secretEnvName, secretFileNameEnvName))
}
