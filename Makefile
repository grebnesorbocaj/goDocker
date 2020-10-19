build:
	docker build -t diakou/go-docker .

run:
	docker run --publish 6060:8080 --name testRun --rm diakou/go-docker
