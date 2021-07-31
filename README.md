# WebArticles
A Demo of Web Article Service built with Go

## Made with
- [Echo Framework](https://github.com/labstack/echo)
- using 💜 and 💰


## Requirements
- [Go (>1.12)](https://golang.org)
- [MySQL 8](https://www.mysql.com)
- [Docker (Optional)](https://www.docker.com)
- [Makefile (Optional)](https://makefiletutorial.com)
- [Go Swagger (Optional)](https://github.com/go-swagger/go-swagger)


## Go Swagger (Optional)
- For using swagger you need install [Go Swagger](https://github.com/go-swagger/go-swagger). 
- For installing you can read documentation [Install](https://goswagger.io/install.html).
- For specification can refer to [Specification](https://swagger.io/specification/)
- Online editor is available on [Swagger Editor](https://editor.swagger.io/)

for example you need create swagger docs in
```
/docs
------ /swagger
------------ /assets
------------ docs.json <- will be generated with go swagger
------------ index.html
------------ swagger.yml
```


## Getting Started
This section will guide you to get the project up and running both for development and production.


### Go Mod
This service need to install all of required dependencies
```
$ go mod tidy
```	


### Swagger
We need to generate `docs.json` in order to swagger works properly after `swagger.yml` has been changed
```
$ make swagger service=webarticles
```

Then change env `USE_SWAGGER` to `true`, then we can access the swagger docs at `http://your-domain/docs`


### Run service
using Makefile
```
$ make run service=webarticles
```
or can be manually using `go run`
```
$ go run ./cmd/webarticles
```


### Clear service
```
make clear service=webarticles
```

### Test service
For using script test. We need you initiate git in your work. Because in script test for get `jsonschema root` we check in repository use library :

"github.com/integralist/go-findroot/find"
```
// Repo uses git via the console to locate the top level directory
func Repo() (Stat, error) {
	path, err := rootPath()
	if err != nil {
		return Stat{
			"Unknown",
			"./",
		}, err
	}

	gitRepo, err := exec.Command("basename", path).Output()
	if err != nil {
		return Stat{}, err
	}

	return Stat{
		strings.TrimSpace(string(gitRepo)),
		path,
	}, nil
}
```

Just add `git init` in your repo, and you can run this :)

```
make test 
```
