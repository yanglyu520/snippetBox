# snippetBox

This project is a journey documenting how I learn web app development with golang step by step.

## Lessons Learnt

### step1: create a module
1. What is a module in golang?
- A module is a group of packages. 
- The requirements of a module is listed in a file go.mod
- go will use go.mod for each installation to build the binary application

2. What does go.mod file contain?
- module path
- go version
- dependencies, use keyword `require`

3. How to create a module?
- run `go mod init <module-path>`
- module path needs to be unique, usually it is your domain name or github repo path
- if you are creating a project which can be downloaded by other people, it is better for your module name to be equal to the path that can be downloaded, like `github.com/yanglyu520/snippetBox`
  
### step2: Create a baisc server
1. How to create a simplest server with go?
- golang net/http library has already provides a set of libraries to create a web server, creating a server is trivial and can be done with a call to ListenAndServe

2. Task: create the simplest server
- run `go run main.go`, and open `localhost:8080` in browser, you will see `page 404 not found` when this server is up. You expect to see `page does not exists` if this server is not up

3. How does http.ListenAndServe function work?
- In your Go program, http.ListenAndServe is a method that takes two parameters: a string that represents the address to listen on, and a handler.
- The address usually has two parts: a host name (or IP) and a port. **If the network address is empty string "", then the default is all network interfaces at port 80**. 
- The handler is an interface that has a ServeHTTP(ResponseWriter, *Request) method. **If the handler is nil, it will use http.DefaultServeMux.**

```go
// ListenAndServe always returns a non-nil error.
// The handler is typically nil, in which case the DefaultServeMux is used.
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

### step3: configure the server above
1. Task: add more configurations to the server you created in step 2
- We can define the server explicitly with configurations and then calling server.ListenAndServe()

```go
// A Server defines parameters for running an HTTP server.
// The zero value for Server is a valid configuration.
type Server struct {
	Addr string
	Handler Handler // handler to invoke, http.DefaultServeMux if nil
	TLSConfig *tls.Config
	ReadTimeout time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout time.Duration
	MaxHeaderBytes int
    ...
}
```
2. How does server.ListenAndServe() work?
- Note the meaning of `:http`. The ":http" address is a shorthand notation that represents listening on all available network interfaces ("") and using the default HTTP port (80 for HTTP) or the default HTTPS port (443 for HTTPS). By using ":http", the server will listen on all interfaces and use the default port for handling incoming HTTP requests.
- Here is the function for server.ListenAndServe.
```go
// If srv.Addr is blank, ":http" is used.
//
// ListenAndServe always returns a non-nil error. After Shutdown or Close,
// the returned error is ErrServerClosed.
func (srv *Server) ListenAndServe() error {
	if srv.shuttingDown() {
		return ErrServerClosed
	}
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}
```

### step4: add handlers to ListenAndServe Function
1. What are the 3 components of a web api?
- router: servemux, a mapping between the URL patterns and corresponding handlers
- handler or multiple handlers
- web server, you can establish a web server with golang's standard net library

2.

### step5: 

## Task List:
### 1. Print Hello World to the terminal
### 2. Create a basic server
### 3. Configure the server
### 4. Create a webpage that says hello world, use at least 3 methods
### 5. routing
### 6. 

               
## Additional Learning
1. Learn how net/tcp works to spin up a server
2. Learn how to use https instead of http, and how to add cert