run:
	go run cmd/main.go

test: run logs

logs:
	tail -f logs.txt

cleanlogs:
	rm logs.txt