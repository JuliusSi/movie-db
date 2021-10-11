# Commands
up:
	docker-compose build && docker-compose up -d

down:
	docker-compose down

delete:
	docker stop $$(docker ps -a -q) && docker rm $$(docker ps -a -q) && docker rmi $$(docker images -a -q)