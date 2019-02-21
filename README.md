# settings

[![Go Report Card](https://goreportcard.com/badge/github.com/dtylman/settings)](https://goreportcard.com/report/github.com/dtylman/settings)

A simple `go` package for managing configuration files.

Saves & loads configuration from files using simple interfaces.

Supported format: JSON

## Usage
```
go get "github.com/dtylman/settings"
```

```go
import (
	"github.com/dtylman/settings"
)
```

## Example

```go
const confFile = "example.conf"

func main() {
    conf := settings.New()
	err := conf.Load(confFile, settings.FormatJSON)
	if err != nil {
		log.Printf("Cannot load configuration file: %v", err)
	}
	log.Printf("App bind: %s", conf.Get("app.bind", "no bind yet"))
	conf.Set("app.bind", "0.0.0.0:8080")
	err = conf.Save(confFile, settings.FormatJSON)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Saved to ", confFile)
	conf.Print(settings.FormatJSON)
}
```

