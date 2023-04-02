run:
	go run cmd/main.go

test: run logs

logs:
	tail -f logs.txt

cleanlogs:
	rm logs.txt

dockerbuild:
	docker build --tag ipchecker .

dockerrun:
	docker run -d --name ipchecker ipchecker

dockerstop:
	docker stop ipchecker

dockerremove:
	docker rm ipchecker

dockerlogs:
	docker logs -f ipchecker

dockerexec:
	docker exec -it ipchecker /bin/sh
