package common

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Config struct {
	App  AppConfig
	HTTP HTTPConfig

	kv map[string]string
}

type AppConfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"production"`
}

type HTTPConfig struct {
	Port int `env:"PORT" envDefault:"8080"`
	Test int `env:"TEST"`
}

func NewConfig() *Config {
	cfg := &Config{
		kv: make(map[string]string),
	}

	cfg.init()
	return cfg
}

func (cfg *Config) String() string {
	if cfg == nil {
		return ""
	}

	var sb strings.Builder
	for k, v := range cfg.kv {
		_, _ = fmt.Fprintf(&sb, "%s=%s\n", k, v)
	}
	return sb.String()
}

func (cfg *Config) init() {
	cfg.parseEnv(reflect.ValueOf(cfg).Elem())
}

func (cfg *Config) parseEnv(value reflect.Value) {
	typeOfValue := value.Type()

	for i := range value.NumField() {
		field := value.Field(i)
		structField := typeOfValue.Field(i)

		if field.Kind() == reflect.Struct {
			cfg.parseEnv(field)
			continue
		}

		envKey, ok := structField.Tag.Lookup("env")
		if !ok || !field.CanSet() {
			continue
		}

		envValue, ok := os.LookupEnv(envKey)
		if !ok || strings.TrimSpace(envValue) == "" {
			envValue, ok = structField.Tag.Lookup("envDefault")
			if !ok {
				continue
			}
		}

		if err := parse(field, envValue); err != nil {
			log.Fatalf("解析环境变量失败，键=%s ，值=%q，类型=%s，错误=%v", envKey, envValue, field.Kind(), err)
		}

		cfg.kv[envKey] = fmt.Sprint(field.Interface())
	}
}

func parse(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)

	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(intValue))

	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)

	default:
		return fmt.Errorf("不支持的类型: %s", field.Kind())
	}

	return nil
}
