package config

import (
	"context"
	"errors"
	"log"
	"reflect"
	"strings"

	"github.com/kaitsubaka/hexagonal-architecture-go/internal/platform/context/base"
	"github.com/spf13/viper"
)

type Configuration struct {
	App      App      `mapstructure:",squash"`
	Postgres Postgres `mapstructure:",squash"`
}

type Postgres struct {
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DBName   string `mapstructure:"POSTGRES_DBNAME"`
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
}

type App struct {
	Env  string `mapstructure:"APP_ENV"`
	Port string `mapstructure:"APP_PORT"`
}

func New(c context.Context) (*Configuration, error) {
	if key, ok := c.Value(base.CtxEnvKey).(string); ok && key == "local" {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
	}
	cfg := new(Configuration)
	// try load settings from env vars
	if err := bindEnv(cfg); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	log.Println(cfg)
	return cfg, nil
}

func (c *Configuration) IsLowEnv() bool {
	switch c.App.Env {
	case "local", "develpment":
		return true
	default:
		return false
	}
}

func bindEnv(cfg any) error {
	// Retrieve the underlying type of variable `i`.
	dataType := reflect.TypeOf(cfg).Elem()
	dataValue := reflect.ValueOf(cfg).Elem()
	// Iterate over each field for the type
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		fieldValue := dataValue.FieldByName(field.Name)
		tag := field.Tag.Get("mapstructure")
		isSquashTag := strings.EqualFold(tag, ",squash")
		if isSquashTag && fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				return errors.New("")
			}
			err := bindEnv(fieldValue.Interface())
			if err != nil {
				return err
			}
			continue
		}
		if isSquashTag && fieldValue.Kind() == reflect.Struct {
			err := bindEnv(fieldValue.Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}
		// Bind the environment variable.
		if err := viper.BindEnv(tag); err != nil {
			return err
		}
	}
	return nil
}
