package figma_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	figma "github.com/dataiads/go-figma-api"
	"golang.org/x/oauth2"
)

func TestFilesAPI(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer secret" {
			t.Errorf("Authorization header = %q", got)
		}
		w.Header().Set("Content-Type", "application/json")

		switch r.URL.Path {
		case "/v1/files/file-key/nodes":
			if got := r.URL.Query().Get("ids"); got != "1:2" {
				t.Errorf("ids = %q", got)
			}
			if got := r.URL.Query().Get("geometry"); got != "paths" {
				t.Errorf("geometry = %q", got)
			}
			fmt.Fprint(w, `{"name":"Design","role":"viewer","lastModified":"2026-08-20T00:00:00Z","editorType":"figma","thumbnailUrl":"https://example.com/thumbnail.png","version":"1","nodes":{}}`)
		case "/v1/images/file-key":
			if got := r.URL.Query().Get("ids"); got != "1:2" {
				t.Errorf("ids = %q", got)
			}
			if got := r.URL.Query().Get("format"); got != "png" {
				t.Errorf("format = %q", got)
			}
			if got := r.URL.Query().Get("scale"); got != "1" {
				t.Errorf("scale = %q", got)
			}
			fmt.Fprint(w, `{"err":null,"images":{"1:2":"https://example.com/image.png"}}`)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)

	cfg := figma.NewConfiguration()
	cfg.HTTPClient = server.Client()
	cfg.Servers = figma.ServerConfigurations{{URL: server.URL}}
	client := figma.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), figma.ContextOAuth2, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "secret"}))

	nodes, _, err := client.FilesAPI.GetFileNodes(ctx, "file-key").Ids("1:2").Geometry("paths").Execute()
	if err != nil {
		t.Fatalf("GetFileNodes: %v", err)
	}
	if nodes.Name != "Design" {
		t.Errorf("name = %q", nodes.Name)
	}

	images, _, err := client.FilesAPI.GetImages(ctx, "file-key").Ids("1:2").Format("png").Scale(1).Execute()
	if err != nil {
		t.Fatalf("GetImages: %v", err)
	}
	if got := images.Images["1:2"]; got != "https://example.com/image.png" {
		t.Errorf("image URL = %q", got)
	}
}
