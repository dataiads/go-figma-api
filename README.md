# go-figma-api

Generated Go client for the two Figma REST API operations used by Dataiads:

- `GET /v1/files/{file_key}/nodes`
- `GET /v1/images/{file_key}`

The package is generated from the bundled Figma OpenAPI specification v0.42.0.

```go
cfg := figma.NewConfiguration()
cfg.HTTPClient = httpClient
client := figma.NewAPIClient(cfg)

ctx := context.WithValue(
	context.Background(),
	figma.ContextOAuth2,
	oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}),
)

nodes, response, err := client.FilesAPI.
	GetFileNodes(ctx, fileKey).
	Ids(nodeID).
	Geometry("paths").
	Execute()
```

Use `ContextAPIKeys` with the `PersonalAccessToken` key when calling Figma with a personal access token instead of OAuth.

Run `./generate.sh` to regenerate the client from [`api/openapi.yaml`](api/openapi.yaml). Generated Go files must not be edited manually.
