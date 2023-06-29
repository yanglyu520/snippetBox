# snippetBox

This project is a journey documenting how I learn web app development with golang step by step.

## Part 1: Inital setup for the project

## Task List:
1. Print Hello World to the terminal
2. Create a basic server
3. Configure the server
4. Create a webpage that says hello world, using either DefaultServeMux, or self-defined ServeMux
5. Create 2 more paths and handlers using above method
6. Update our application so that /snippet/create route only responds to HTTP request `POST`
7. 


## Lessons Learnt

### step 1: create a module
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
  
### step 2: Create a baisc server
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

### step 3: configure the server above
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

### step 4: add handlers to ListenAndServe Function
1. What are the 3 components of a web api?
- router: servemux, a mapping between the URL patterns and corresponding handlers
- handler or multiple handlers
- web server, you can establish a web server with golang's standard net library

2. Note: go's servemux treats `/` like catch-all, so at the moment all the http requests to our server will be handled by our home handler function, regardless of its url path

3. Explain how `host:port` work for http.ListenAndServe first parameter?
-  if you omit the host, and put `:4000`, then the server will listen on all your computer's available network interfaces
-  in other golang project, you might come across `:http`, or `:http-alt`
-  if you use a named port, then go will attempt to look up the relevant port number from your `/etc/services` file when starting the server, or will return an error if a match cannot be found

4. Explain how `go run xxx` works?
- `go run` is a shortcut that creates the executable binary in your temporary foler `/tmp` and runs the binary in one step
- it accepts a space separated go files, the path to a specific package, or the full module path

### step 5: Adding more paths and corresponding handlers
1. Task: Add more paths and corresponding handlers in the step4's code
2. Note: go's servemux supports 2 different types of URL patterns: fixed paths and subtree path
3. Explain what is fixed path and subtreepath?
- `fixed path` does not end with `/`, and these are only matched, then the handlers will be called
- `subtree path` ends with `/`, these are matched whenever **the start** of a request url path matches the subtree path
4. Task: Add code so that if you dont want `/` to act like catch-all, and receive 404 page not found 
5. Why not user defaultServeMux for the multiple handlers you code before?
if I use defaultServeMux, the code might be looking like this:
```go
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet/view", snippetView)
    http.HandleFunc("/snippet/create", snippetCreate)
	
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	
	log.Fatal(err)
}
```
The logic is very similar, http.ListenAndServe will create a defaultServeMux, just an instance of the ServeMux we definded.
But it is not recommended to use defaultServeMux for production code, because DefaultServeMux is a global variable and any 
package can access it and register a route - including 3rd party packages.
So for the sake of security, it is recommended to avoid using defaultServeMux.
6. Servemux features and quirks?
- in go's servemux, longer URL patterns takes precedence over the shorter ones, so if a servemux contains multiple
patterns that match a request, it will always dispatch the request to the handler corresponding the longest pattern
- request url paths are automatically sanitized. Ex, if a path contains `.` or `..` or repeated slashes, the user will be 
directed to a cleaned up url
- if a subtree path has been registered and a request is received for the subtree path without trailing slash,
then the user will automatically be sent a 301 permanent redirect to the subtree path with the slash added. Ex:
`/foo` will be redirected to `/foo/` 
- go's servemux does not support routing based on the request method
- go's servemux does not support clean url with variables in them
- go's servemux does not support regex-based patterns

### step6: Customise HTTP headers
1. Note:
- It is only possible to call `w.WriteHeader()` **only once per response**, and after the status code has been written, it cannot be changed.
- if `w.WriteHeader()` is not called before `w.Write()`, then it will automatically send a `200 OK` status code to the user.
Therefore, it is important to call  `w.WriteHeader()` before `w.Write()`.
- make sure you call `w.Header().Set("headerName", "headerValue")` before `w.WriteHeader()` and `w.Write()`, otherwise it will have no effect on the headers that the users receive.

2.  How to send a non-200 status code and a plain-text response body in one function?
- We can use `http.Error(w, "message", 405)` instead of using `w.WriteHeader()` and `w.Write()`
- we are passing responseWriter to another function that sends a response to the user. It is rare to use `w.WriteHeader()` and `w.Write()` methods directly.

3. Can we use `net/http` constants instead of `405`?
- we can use `http.MethodPost` instead of `POST`
- we can use `http.StatusMethodNotAllowed` instead of the integer `405`

4. What are go's system-generated headers?
- `Date`, `Content-Length`, `Content-Type`
- Note that go will attempt to set the correct one for you but content sniffing the response body with `httpDetectContentType()` function.  `httpDetectContentType()` generally works well, except that it cannot distinguish JSON from plaintext, so by default, JSON response will be sent with a `Content-Type: text/plain; charset=utf-8` header. You can prevent this happening by settting the correct header manually like so:
```go
w.Header().Set("Content-Type", "application/json")
```
- If this function cannot guess the content type, go will fall back to setting the header `Content-Type: application/octet-stream`.

5. What is difference between `Set(), Add(), Del(), Get() and Values()` methods on the header map?
- w.Header().Set() adds a new header to the response header map
- w.Header().Del() does not remove system-generated headers, to suppress these, you need to access the underlying header map directly, and set the vaule to nil like so
```go
w.Header()["Date"] = nil

```

6. Header Canaonicalization?
- When we are using those methods on header map, the name of the header will always be canonicalized using `textproto.CanonicalMiMEHeaderKey() function`. This has the practical implicationthat when calling these methods, the header name is case-insensitive.

7. How to avoid this header canonicalization?
- `w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}`
- Note if a HTTP/2 connection is being used, go will always automatically convert the header names and values to lowercase as per HTTP/2 specifications


### step 7: add url query string

### step 8: reorganize files above with golang project organization style

### step 9: html templating

### step 10: serving static files



               
## Additional Learning
1. Learn how net/tcp works to spin up a server
2. Learn how to use https instead of http, and how to add cert
3. Learn more about when to use which specific Http response status for web api, like 301(permanent redirect), 405(method not allowed)
4. 