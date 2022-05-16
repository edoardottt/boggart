SHELL := /bin/bash
BOGGART_IMAGE := boggart_boggart:latest

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
	docker rmi $(BOGGART_IMAGE)

#---Delete the volume---
cleanvol:
	sudo rm -rf /boggart-data

#---Prune---
prune:
	docker system prune -f

#---Restart the service (applying the changes made)---
restart:
	docker-compose down
	docker rmi $(BOGGART_IMAGE)
	docker-compose up