package main

type Config struct {
	Env string `env:"ENV,required"`

	Port int `env:"PORT,required"`
}
