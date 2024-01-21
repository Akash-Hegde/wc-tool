build:
	go build -o ~/bin/ccwc cmd/ccwc/main.go

hello:
	echo "hello"

all: hello build