# ENV File Resolver for [kong](https://github.com/alecthomas/kong)

## Usage

```go
package main

import (
	"github.com/alecthomas/kong"
	dotenv "github.com/bluegosolutions/kong-dotenv-go"
)

type App struct {
	EnvFile kong.ConfigFlag `kong:"optional,name=env-file,help='Path to .env file'"`
}

func main() {
	var app App
	ctx := kong.Parse(&app, kong.Configuration(dotenv.ENVFile))
	ctx.FatalIfErrorf(ctx.Run())
}
```
