package main

import (
	"context"
	"fmt"
	"os"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	project, _, err := c.Project(ctx, "pypi", "cookiecutter")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s\n", *project.Name)
}
