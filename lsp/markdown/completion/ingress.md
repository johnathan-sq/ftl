Declare an ingress function.

Verbs annotated with `ftl:ingress` will be exposed via HTTP (http is the default ingress type). These endpoints will then be available on one of our default ingress ports (local development defaults to http://localhost:8891).

```go
type GetRequest struct {
	UserID string `json:"userId"`
	PostID string `json:"postId"`
}

type GetResponse struct {
	Message string `json:"msg"`
}

//ftl:ingress GET /http/users/{userId}/posts
func Get(ctx context.Context, req builtin.HttpRequest[GetRequest]) (builtin.HttpResponse[GetResponse, string], error) {
  return builtin.HttpResponse[GetResponse, string]{
    Status:  200,
    Body:    ftl.Some(GetResponse{}),
  }, nil
}
```

See https://tbd54566975.github.io/ftl/docs/reference/ingress/
---
type ${1:Func}Request struct {
	${2:Field} ${3:Type} `json:"${4:field}"`
}

type ${1:Func}Response struct {
	${5:Field} ${6:Type} `json:"${7:field}"`
}

//ftl:ingress ${8:GET} ${9:/url/path}
func ${1:Func}(ctx context.Context, req builtin.HttpRequest[${1:Func}Request]) (builtin.HttpResponse[${1:Func}Response, string], error) {
	${7:// TODO: Implement}
	return builtin.HttpResponse[${1:Func}Response, string]{
		Status: 200,
		Body: ftl.Some(${1:Func}Response{}),
	}, nil
}