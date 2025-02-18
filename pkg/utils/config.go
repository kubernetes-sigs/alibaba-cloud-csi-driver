package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const DefaultConfigPath = "/etc/csi-plugin/config"

type Config struct {
	Path string
}

func DefaultConfig() Config {
	return Config{Path: DefaultConfigPath}
}

func (c *Config) Get(configKey, env, defaultValue string) string {
	if env != "" {
		if value, ok := os.LookupEnv(env); ok {
			return value
		}
	}
	if configKey != "" {
		value, err := os.ReadFile(filepath.Join(c.Path, configKey))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return defaultValue
			}
			panic(fmt.Sprintf("failed to reading config file for %q: %v", configKey, err))
		}
		return string(value)
	}
	return defaultValue
}

func (c *Config) GetBool(configKey, env string, defaultValue bool) bool {
	s := c.Get(configKey, env, "")
	if s == "" {
		return defaultValue
	}

	switch strings.ToLower(s) {
	case "true", "yes", "enable", "enabled":
		return true
	case "false", "no", "disable", "disabled":
		return false
	default:
		v, err := strconv.ParseBool(s)
		if err != nil {
			panic(fmt.Sprintf("invalid bool value: %s for %s", s, configKey))
		}
		return v
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
