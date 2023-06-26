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

2. Explain how this simplest server works?
- run `go run main.go`, and open `localhost:8080` in browser, you will see `page 404 not found` when this server is up. You expect to see `page does not exists` if this server is not up

3. How does http.ListenAndServe function work?
- In your Go program, http.ListenAndServe is a method that takes two parameters: a string that represents the address to listen on, and a handler.
- The address usually has two parts: a host name (or IP) and a port. The hostname is optional and if you don't specify one, the server will listen on all available interfaces. In your case, "localhost:8080" means that your server will run locally (on your own machine) and listen on port 8080.
- The handler is an interface that has a ServeHTTP(ResponseWriter, *Request) method. The http.ListenAndServe function will pass all requests it receives to this handler. If the handler is nil, it will use http.DefaultServeMux. 
- A ServeMux is essentially an HTTP request router (or multiplexer) that matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.
                                                              
### step3: configure the server above

### step4: add handlers to ListenAndServe Function

### step5: 

## Action Steps:
### 1. Print Hello World to the terminal
### 2. Create a basic server
### 3. Modify the basic server
### 4. Create a webpage that says hello world, use at least 3 methods
### 5. routing
### 6. 

               
## Additional Learning
1. Learn how net/tcp works to spin up a server
2. Learn how to use https instead of http, and how to add cert