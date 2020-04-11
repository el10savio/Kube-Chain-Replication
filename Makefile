
image-clean:
	@echo "Cleaning GoServer Docker Container"
	docker ps -a | awk '$$2 ~ /goserver/ {print $$1}' | xargs -I {} docker rm -f {}
	
image-build:
	@echo "Building GoServer Docker Image"	
	cd GoServer; docker build -t goserver -f Dockerfile .

image-run:
	@echo "Running GoServer Docker Container"
	docker run -p 8080:8080 -d goserver

proxy-clean:
	@echo "Cleaning GoProxy Docker Container"
	docker ps -a | awk '$$2 ~ /goproxy/ {print $$1}' | xargs -I {} docker rm -f {}
	
proxy-build:
	@echo "Building GoProxy Docker Image"	
	cd GoProxy; docker build -t goproxy -f Dockerfile .

proxy-run:
	@echo "Running GoProxy Docker Container"
	docker run -p 8090:8090 -d goproxy

redis-clean:
	@echo "Cleaning Redis Docker Container"
	docker ps -a | awk '$$2 ~ /redis/ {print $$1}' | xargs -I {} docker rm -f {}
	
redis-run:
	@echo "Running Redis Docker Container"
	docker run -d -p 6379:6379 --name redis redis

redis: redis-run

image: image-build image-run

proxy: proxy-build proxy-run

clean: redis-clean image-clean proxy-clean

all: clean redis image proxy

stop: clean
