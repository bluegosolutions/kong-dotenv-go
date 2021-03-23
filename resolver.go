package kongdotenv

import (
	"io"

	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
)

// ENVFile returns a kong.Resolver that retrieves values from a .env file source.
//
// ENVFile resolves only flags with `env:"X"` tag.
func ENVFile(r io.Reader) (kong.Resolver, error) {
	values, err := godotenv.Parse(r)
	if err != nil {
		return nil, err
	}

	var f kong.ResolverFunc = func(context *kong.Context, parent *kong.Path, flag *kong.Flag) (interface{}, error) {
		if flag.Value != nil || flag.Env == "" {
			return nil, nil
		}

		raw, ok := values[flag.Env]
		if !ok {
			return nil, nil
		}
		return raw, nil
	}

	return f, nil
}
