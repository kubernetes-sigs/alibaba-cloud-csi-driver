package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ConfigMap map[string]string
}

func (c *Config) Get(configKey, env, defaultValue string) string {
	if value, ok := os.LookupEnv(env); ok {
		return value
	}
	if value, ok := c.ConfigMap[configKey]; ok {
		return value
	}
	return defaultValue
}

func (c *Config) GetBool(configKey, env string, defaultValue bool) bool {
	strValue := os.Getenv(env)
	if strValue == "" {
		strValue = c.ConfigMap[configKey]
	}
	if strValue == "" {
		return defaultValue
	}

	switch strings.ToLower(strValue) {
	case "true", "yes", "enable":
		return true
	case "false", "no", "disable":
		return false
	default:
		panic(fmt.Sprintf("invalid bool value: %s for %s", strValue, configKey))
	}
}

func (c *Config) GetInt(configKey, env string, defaultValue int) int {
	s := c.Get(configKey, env, "")
	if s == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid int value: %s for %s: %v", s, configKey, err))
	}
	return i
}
