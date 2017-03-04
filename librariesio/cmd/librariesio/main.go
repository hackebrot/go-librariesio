package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/hackebrot/go-librariesio/librariesio"
)

func loadFromEnv(keys ...string) (map[string]string, error) {
	env := make(map[string]string)

	for _, key := range keys {
		v := os.Getenv(key)
		if v == "" {
			return nil, fmt.Errorf("environment variable %q is required", key)
		}
		env[key] = v
	}

	return env, nil
}

func main() {
	env, err := loadFromEnv("LIBRARIESIO_API_KEY")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	c := librariesio.NewClient(strings.TrimSpace(env["LIBRARIESIO_API_KEY"]))
	project, _, err := c.GetProject("pypi", "cookiecutter")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%+v\n", project)
}
