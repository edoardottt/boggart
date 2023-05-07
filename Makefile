SHELL := /bin/bash
BOGGART_IMAGE := boggart_boggart:latest

#---Run testcases---
test:
	@go test -race ./...

#---Run linter---
lint:
	@golangci-lint run

#---Run the service---
up:
	@sudo docker-compose up

#---Stop the service---
stop:
	@sudo docker-compose stop

#---Delete the service---
down:
	@sudo docker-compose down

#---Delete the image created---
clean:
	@sudo docker rmi $(BOGGART_IMAGE)

#---Delete the volume---
cleanvol:
	@sudo rm -rf /boggart-data

#---Prune---
prune:
	@sudo docker system prune -f

#---Restart the service (applying the changes made)---
restart:
	@sudo docker-compose down
	@sudo docker rmi $(BOGGART_IMAGE)
	@sudo docker-compose up