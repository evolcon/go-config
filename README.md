GoConfig
========

A simple Go package with built-in validation. It allows configuration loading from multiple sources such as `YAML` files, `.env` files, and environment variables.

This library is just a wrapper around popular libraries such as [godotenv](http://github.com/joho/godotenv), [envconfig](http://github.com/kelseyhightower/envconfig), [yaml](https://github.com/go-yaml/yaml/tree/v3.0.1) and [go-playground/validator](https://github.com/go-playground/validator).

Installation
------------

    go get github.com/evolcon/go-config

🚀 Quick Start
------------

### 1. Define your config files

Add [config.yaml](main/config.yml) and [.env](main/.env) files
```yaml
# config.yaml
server:
  port: 8080
```

```.env
# .env
SERVER_HOST=0.0.0.0
```
#### 2. Define your app config structure and call its filling in [main file](main/example.go)

```go
package main

import (
	"fmt"

	goconfig "github.com/evolcon/go-config"
)

// defining app config
type AppConfig struct {
	Server struct {
		Host string
		Port int
	}

	Debug bool
}

func main() {
	cfg := &AppConfig{}

    // call package initialization and config filling
	goconfig.InitOnce()
	if err := goconfig.Fill(cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}

```

#### 3. Run your application

```go
# Environment variables only (no flags needed) and .env file is readed by default
go run main.go

# Using YAML config and .env
go run main.go --yaml-config=config.yaml

# Defining specific .env file.
go run main.go --env-config=.env

# With environment prefix
go run main.go --env-prefix=MYAPP

# Multiple .env files
go run main.go --env-config=".env,.env.local"
```

📖 API Reference
------------

`InitOnce()` Initializes the package configuration. Must be called once at application startup.

```go
func main() {
    goconfig.InitOnce()
    // ... rest of your application
}
```

`Fill(config any) error` Populates the provided configuration structure with data from all configured sources. Supports validation using [go-playground/validator](https://github.com/go-playground/validator) tags.

Parameters:
- Receive `config`: Pointer to your configuration structure
- Returns `error`: Validation or loading error if any

```go
var cfg MyConfig
if err := goconfig.Fill(&cfg); err != nil {
    // Handle error
}
```

🛠 Configuration Sources
------------

### YAML Files

Define your configuration in YAML format:

```yaml
server:
  host: "localhost"
  port: 8080
  timeout: 30s

database:
  url: "postgres://user:pass@localhost:5432/db"
  max_connections: 100

debug: true
features:
  - "auth"
  - "cache"
```

### Environment Variables & .env Files

```env
SERVER_HOST=0.0.0.0
SERVER_PORT=3000
DATABASE_URL=postgres://user:pass@localhost:5432/prod
DEBUG=true
FEATURES=auth,cache,api
```

### Environment Prefix
Use prefixes to avoid naming conflicts:

```go
// Run with --env-prefix=MYAPP
type Config struct {
    Host string `envconfig:"HOST"` // Becomes MYAPP_HOST
}
```

✅ Validation
------------

The package uses `go-playground/validator`. Add validation tags to your struct fields:

```go
type Config struct {
    Email    string `validate:"required,email"`
    Port     int    `validate:"required,min=1,max=65535"`
    URL      string `validate:"url"`
    Timeout  int    `validate:"gte=0"`
}
```

⚙️ Additional configuration and flexibility
------------

You can also describe your configs in a more precise and flexible way, if the keys in your env and yaml configs differ from the fields in your config structure.

```go

type App struct {
    Name    string `yaml:"name" envconfig:"NAME"`
    Version string `yaml:"version" envconfig:"VERSION"`
}

type Config struct {
    AppConfig App `yaml:"app" envconfig:"APP"`  # redefines key for env and yaml prefix

    Server struct {
        Host string `yaml:"host" envconfig:"SERVER_HOST" validate:"required"`
        Port int    `yaml:"port" envconfig:"SERVER_PORT" validate:"required"`
    } `yaml:"server"` # redefines key for yaml prefix
}
```

🔄 Priority Order
------------

Configuration sources are loaded in this order (later sources override earlier ones):
1. YAML file (if provided via `--yaml-config`)
2. Environment variables (always loaded)
3. .env file(s) (if provided via `--env-config`)

📋 Complete Example
------------

### Project Structure
```
myapp/
├── cmd/app/main.go
├── configs/config.yaml
├── internal
├───── configs
├──────── config.go
├── .env.reference
└── .env.local
```

🎯 Best Practices
------------

1. Call `InitOnce()` early in your application
2. Use meaningful field names in your struct
3. Add validation tags for important fields
4. Use environment prefixes in shared environments
5. Provide sensible defaults in your YAML files
6. Use multiple .env files for different environments

📄 License
===
Apache License - feel free to use in your projects.

---
This package simplifies configuration management in Go applications by providing a clean, unified interface for multiple configuration sources with built-in validation.