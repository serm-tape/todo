# todo
## download
download executablable file [here](goo.gl/jGRM9V)
## config
Set your database in config file as in `config/config.json`. Don't forget to create database in your config. Table will be migrated automatically
## run
### windows
todo_windows -config=config.json
### osx
./todo_osx -config=config.json
### linux
./todo_linux -config=config.json
## Build it yourself
- install [golang](https://golang.org/dl/)
- install [godep](https://github.com/tools/godep)
- run `godep restore`
- run `go run main.go` or `go build .` then `./todo -config=config.json`
