build:
	docker build -t diakou/go-docker:1.2 .

run:
	docker run --publish 6060:8080 --name testRun --rm diakou/go-docker:1.2
