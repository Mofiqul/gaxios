# A http client in Golang inspired from nodejs axios

## Installation 

```bash
go get github.com/mofiqul/gaxios
```

## Read the [Documentation](https://pkg.go.dev/github.com/mofiqul/gaxios)

## Example 

```go
type ResponseData struct {
	Message string `json:"message"`
}

res, err := gaxios.Get(fmt.Sprintf("%s/test", server.URL), &gaxios.GAxiosConfig{
	Header: http.Header{
		"Accept": []string{"application/json"},
	},
}) 
defer res.Data.Close()
if err != nil {
	// Handle error
}

resp := &ResponseData{}
_ = json.NewDecoder(res.Data).Decode(resp)
```

Or create a instance with configuation 

```go
var axios = gaxios.New( 
	&gaxios.GAxiosConfig{
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	})

res, err := axios.Get(fmt.Sprintf("%s/test", server.URL))
defer res.Data.Close()
if err != nil {
	// Handle error
}

resp := &ResponseData{}
_ = json.NewDecoder(res.Data).Decode(resp)
```

