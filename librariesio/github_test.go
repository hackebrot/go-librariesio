package librariesio

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/hackebrot/go-repr/repr"
)

func TestUser(t *testing.T) {
	server, mux, url := startNewServer()
	client := NewClient(APIKey)
	client.BaseURL = url
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if method := "GET"; method != r.Method {
			t.Errorf("expected HTTP %v request, got %v", method, r.Method)
		}

		if url := r.URL.String(); !strings.Contains(url, "/github/hackebrot") {
			t.Errorf("unexpected URL, got %v", url)
		}

		fmt.Fprintf(w, `{
			"login":"hackebrot",
			"name":"Raphael Pierzina",
			"blog":"https://raphael.codes"
		}`)
	})

	user, _, err := client.User(context.Background(), "hackebrot")
	if err != nil {
		t.Fatalf("User returned unexpected error: %v", err)
	}

	want := &User{
		Name:  String("Raphael Pierzina"),
		Login: String("hackebrot"),
		Blog:  String("https://raphael.codes"),
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("\nExpected %v\nGot %v", repr.Repr(want), repr.Repr(user))
	}
}

func TestUserProjects(t *testing.T) {
	server, mux, url := startNewServer()
	client := NewClient(APIKey)
	client.BaseURL = url
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if method := "GET"; method != r.Method {
			t.Errorf("expected HTTP %v request, got %v", method, r.Method)
		}

		if url := r.URL.String(); !strings.Contains(url, "/github/hackebrot/projects") {
			t.Errorf("unexpected URL, got %v", url)
		}

		fmt.Fprintf(w, `[
			{
				"name":"go-repr",
				"keywords": ["go", "golang", "helper"],
				"repository_url": "https://github.com/hackebrot/go-repr"
			},
			{
				"name":"go-librariesio",
				"keywords": ["api-client", "go", "golang"],
				"repository_url": "https://github.com/hackebrot/go-librariesio"
			}
		]`)
	})

	projects, _, err := client.UserProjects(context.Background(), "hackebrot")

	if err != nil {
		t.Fatalf("UserProjects returned unexpected error: %v", err)
	}

	want := []*Project{
		{
			Name: String("go-repr"),
			Keywords: []*string{
				String("go"),
				String("golang"),
				String("helper"),
			},
			RepositoryURL: String("https://github.com/hackebrot/go-repr"),
		},
		{
			Name: String("go-librariesio"),
			Keywords: []*string{
				String("api-client"),
				String("go"),
				String("golang"),
			},
			RepositoryURL: String("https://github.com/hackebrot/go-librariesio"),
		},
	}

	if !reflect.DeepEqual(projects, want) {
		t.Errorf("\nExpected %v\nGot %v", repr.Repr(want), repr.Repr(projects))
	}
}

func TestUserRepositories(t *testing.T) {
	server, mux, url := startNewServer()
	client := NewClient(APIKey)
	client.BaseURL = url
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if method := "GET"; method != r.Method {
			t.Errorf("expected HTTP %v request, got %v", method, r.Method)
		}

		if url := r.URL.String(); !strings.Contains(url, "/github/hackebrot/repositories") {
			t.Errorf("unexpected URL, got %v", url)
		}

		fmt.Fprintf(w, `[
			{
				"full_name": "hackebrot/go-repr",
				"default_branch": "master",
				"language": "Go",
				"keywords": ["go", "golang", "helper"],
				"license": "MIT",
				"pushed_at": "2017-03-18T23:55:35.000Z",
				"private": false
			},
			{
				"full_name": "hackebrot/dotfiles",
				"language": "Shell",
				"pushed_at": "2017-01-22T20:57:47.000Z"
			}
		]`)
	})

	repos, _, err := client.UserRepositories(context.Background(), "hackebrot")

	if err != nil {
		t.Fatalf("UserRepositories returned unexpected error: %v", err)
	}

	want := []*Repository{
		{
			FullName:      String("hackebrot/go-repr"),
			DefaultBranch: String("master"),
			Language:      String("Go"),
			Keywords: []*string{
				String("go"),
				String("golang"),
				String("helper"),
			},
			License:  String("MIT"),
			PushedAt: Time(time.Date(int(2017), time.March, int(18), int(23), int(55), int(35), int(0), time.UTC)),
			Private:  Bool(false),
		},
		{
			FullName: String("hackebrot/dotfiles"),
			Language: String("Shell"),
			PushedAt: Time(time.Date(int(2017), time.January, int(22), int(20), int(57), int(47), int(0), time.UTC)),
		},
	}

	if !reflect.DeepEqual(repos, want) {
		t.Errorf("\nExpected %v\nGot %v", repr.Repr(want), repr.Repr(repos))
	}
}
