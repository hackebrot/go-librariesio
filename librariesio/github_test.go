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
