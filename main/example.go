package main

import (
	"fmt"

	goconfig "github.com/evolcon/go-config"
)

type AppConfig struct {
	Server struct {
		Host string
		Port int
	}

	Debug bool
}

func main() {
	cfg := &AppConfig{}

	goconfig.InitOnce()
	if err := goconfig.Fill(cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
