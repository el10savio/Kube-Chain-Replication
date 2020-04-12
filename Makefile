
server-clean:
	@echo "Cleaning GoServer Docker Container"
	docker ps -a | awk '$$2 ~ /goserver/ {print $$1}' | xargs -I {} docker rm -f {}
	
server-build:
	@echo "Building GoServer Docker Image"	
	cd GoServer; docker build -t goserver -f Dockerfile .

server-run:
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

cluster: 
	@echo "Starting Cluster"
	kubectl apply -f goredis.yaml
	kubectl create clusterrolebinding list-view --clusterrole=view --serviceaccount=craq:default
	kubectl apply -f goproxy.yaml

redis: redis-run

server: server-build server-run

proxy: proxy-build proxy-run

clean: redis-clean server-clean proxy-clean

image: clean redis server proxy

all: image cluster

stop: clean
