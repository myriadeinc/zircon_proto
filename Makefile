
T = $(TAG)

dev:
	docker build -f Dockerfile.dev -t myriadeinc/zircon_proto:dev .

up:
	docker-compose up

build: 
	docker build -f Dockerfile -t myriadeinc/zircon_proto:${T} .