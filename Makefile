gitBranch = $(shell git rev-parse --abbrev-ref HEAD)
gitTime = $(shell date +'%FT%T%:z')

build:
	echo ">>> git branch: ${gitBranch}, git time: ${gitTime}"
	npm run build
	mkdir -p target
	go build -o target/main main.go

run-go:
	mkdir -p target
	go build -o target/main main.go
	./target/main

run-react:
	npm run local

