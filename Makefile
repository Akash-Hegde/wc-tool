build:
	go build -o ~/bin/ccwc cmd/ccwc/main.go

run:
	 go run cmd/ccwc/main.go -c -w -l -m test.txt

hello:
	echo "hello"

all: hello build