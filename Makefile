help:
	echo "Usage: make head, make dd"
head:
	go build -o myhead cmd/head/main.go
dd:
	go build -o mydd cmd/dd/main.go
watch:
	go build -o mywatch cmd/watch/main.go
clean:
	rm mydd myhead
