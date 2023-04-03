run:
	go run cmd/main.go

test: run logs

logs:
	tail -f logs.txt

cleanlogs:
	rm logs.txt

dockerbuild:
	docker build --tag ipchecker .

up:
	docker run -d --name ipchecker ipchecker

down:
	docker stop ipchecker

delete:
	docker rm ipchecker

dl:
	docker logs -f ipchecker

exec:
	docker exec -it ipchecker /bin/sh
