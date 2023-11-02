run-local:
	go run .

docker-build:
	docker build -t "kaizenc/go-resume-api" -f "./Dockerfile" .

docker-run:
	docker run -p 8000:8000 kaizenc/go-resume-api
