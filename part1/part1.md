# snippetBox

This project is a journey documenting how I learn web app development with golang step by step.

## Lessons Learnt

### step1: create a module
1. What is a module in golang?
- A module is a group of packages. 
- The requirements of a module is listed in a file go.mod
- go will use go.mod for each installation to build the binary application

2. what does go.mod file contain?
- module path
- go version
- dependencies, use keyword `require`

3. how to create a module?
- run `go mod init <module-path>`
- module path needs to be unique, usually it is your domain name or github repo path
- if you are creating a project which can be downloaded by other people, it is better for your module name to be equal to the path that can be downloaded, like `github.com/yanglyu520/snippetBox`



