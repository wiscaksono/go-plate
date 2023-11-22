.PHONY: setup run build stop

setup:
	docker-compose up --build

run:
	docker-compose up

build:
	docker-compose build

stop:
	docker-compose down
