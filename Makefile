SHELL := /bin/bash

#---Run testcases---
test:
	go test ./...

#---Run linter---
lint:
	golangci-lint run

#---Run the service---
up:
	docker-compose up

#---Stop the service---
stop:
	docker-compose stop

#---Delete the service---
down:
	docker-compose down

#---Delete the image created---
clean:
	docker rmi boggart_boggart:latest

#---Delete the mongo-volume volume---
cleanvol:
	sudo rm -rf /mongo-volume

#---Prune---
prune:
	docker system prune -f

#---Restart the service (applying the changes made)---
restart:
	docker-compose down
	docker rmi boggart_boggart:latest
	docker-compose up