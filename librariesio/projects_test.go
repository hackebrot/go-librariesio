package librariesio

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hackebrot/go-repr/repr"
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
		t.Errorf("\nExpected %v\nGot %v", repr.Repr(want), repr.Repr(project))
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
		t.Errorf("\nExpected %v\nGot %v", repr.Repr(want), repr.Repr(project))
	}
}

func TestSearch(t *testing.T) {
	server, mux, url := startNewServer()
	client := NewClient(APIKey)
	client.BaseURL = url
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if method := "GET"; method != r.Method {
			t.Errorf("expected HTTP %v request, got %v", method, r.Method)
		}

		if url := r.URL.String(); !strings.Contains(url, "/search") {
			t.Errorf("unexpected URL, got %v", url)
		}

		fmt.Fprintf(w, `[
			{
				"name":"pytest-cookies",
				"keywords": ["testing", "python", "cookiecutter"]
			},
			{
				"name":"pytest",
				"keywords": ["testing", "python"]
			}
		]`)
	})

	projects, _, err := client.Search(context.Background(), "pytest")

	if err != nil {
		t.Fatalf("Search returned unexpected error: %v", err)
	}

	want := []*Project{
		&Project{
			Name: String("pytest-cookies"),
			Keywords: []*string{
				String("testing"),
				String("python"),
				String("cookiecutter"),
			},
		},
		&Project{
			Name: String("pytest"),
			Keywords: []*string{
				String("testing"),
				String("python"),
			},
		},
	}

	if !reflect.DeepEqual(projects, want) {
		t.Errorf("\nExpected %v\nGot %v", repr.Repr(want), repr.Repr(projects))
	}
}
