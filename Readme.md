# Embed a React app in a go web-server

In this example the react app is available under the ui folder. Before being able to serve it we need to build it and we can use for that the react script available

```bash
npm run build
```

this will build the react app into the folder ui/build.

This folder is embedded into our go app using the embed library
```go
var (
	//go:embed ui/build
	web embed.FS
)
```

Now lets build the binary

```
go build -o bin/embedded-react-app main.go
```


and run the embedded app
```bash
./bin/embedded-react-app
Your embedded react app is running on http://localhost:8090
```

