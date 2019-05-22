build: main.go
	CGO_ENABLED=0 GOOS=linux go build -o main -ldflags '-extldflags "-static"' main.go

run: 
	go run main.go

dbuild: main dockerfile
	docker build -t chatroom .

drun:
	docker run -d -p 8089:8089 chatroom:latest
