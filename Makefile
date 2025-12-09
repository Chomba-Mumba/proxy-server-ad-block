## run: starts http services

.PHONY: run-containers
run-containers:
	docker run --rm -d -p 9001:80 --name server1 #imagename
	docker run --rm -d -p 9002:80 --name server2 #imagename
	docker run --rm -d -p 9003:80 --name server3 #imagename

##stop: stops all http services
.PHONY: stop
stop:
	docker stop server1
	docker stop server2
	docker stop server3

## run: starts http services
run-proxy-server:
	go run server/cmd/main.go