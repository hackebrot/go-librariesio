package librariesio

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetProject(t *testing.T) {
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

	project, _, err := client.GetProject(context.Background(), "pypi", "cookiecutter")
	if err != nil {
		t.Fatalf("GetProject returned unexpected error: %v", err)
	}

	name := "cookiecutter"
	want := &Project{Name: &name}

	if !reflect.DeepEqual(project, want) {
		t.Errorf("\nExpected %#v\nGot %#v", want, project)
	}
}
