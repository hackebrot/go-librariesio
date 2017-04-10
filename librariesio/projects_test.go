package librariesio

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestProject(t *testing.T) {
	server, mux, url := startNewServer()
	client := NewClient(APIKey)
	client.BaseURL = url
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if method := "GET"; method != r.Method {
			t.Errorf("expected HTTP %v request, got %v", method, r.Method)
		}
		fmt.Fprintf(w, `{"name":"cookiecutter"}`)
	})

	project, _, err := client.Project(context.Background(), "pypi", "cookiecutter")
	if err != nil {
		t.Fatalf("Project returned unexpected error: %v", err)
	}

	name := "cookiecutter"
	want := &Project{Name: &name}

	if !reflect.DeepEqual(project, want) {
		t.Errorf("\nExpected %#v\nGot %#v", want, project)
	}
}

func TestProjectDeps(t *testing.T) {
	server, mux, url := startNewServer()
	client := NewClient(APIKey)
	client.BaseURL = url
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if method := "GET"; method != r.Method {
			t.Errorf("expected HTTP %v request, got %v", method, r.Method)
		}

		if url := r.URL.String(); !strings.Contains(url, "/dependencies") {
			t.Errorf("unexpected URL, got %v", url)
		}

		fmt.Fprintf(w, `{
			"name":"ava",
			"dependencies": [
				{"name": "mocha", "latest": "3.2.0", "outdated": false},
				{"name": "chalk", "latest": "1.1.3", "outdated": true}
			]
		}`)
	})

	project, _, err := client.ProjectDeps(context.Background(), "npm", "ava", "latest")

	if err != nil {
		t.Fatalf("ProjectDeps returned unexpected error: %v", err)
	}

	want := &Project{
		Name: String("ava"),
		Dependencies: []*ProjectDependency{
			{
				Latest:   String("3.2.0"),
				Name:     String("mocha"),
				Outdated: Bool(false),
			},
			{
				Latest:   String("1.1.3"),
				Name:     String("chalk"),
				Outdated: Bool(true),
			},
		},
	}

	if !reflect.DeepEqual(project, want) {
		t.Errorf("\nExpected %#v\nGot %#v", want, project)
	}
}
