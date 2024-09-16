docker-compose:
	docker-compose -f docker-compose.yml up
dockerfile:
	docker build . -t kingseafood
dev:
	go run cmd/dev/main.go