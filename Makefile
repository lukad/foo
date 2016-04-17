default: builddocker

setup:
	go get github.com/op/go-logging
	go get github.com/gorilla/mux

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o helix ./go/src/github.com/lukad/helix

builddocker:
	docker build -t lukad/helix -f ./Dockerfile.build .
	docker run -t lukad/helix /bin/true
	docker cp `docker ps -q -n=1`:/helix .
	chmod 755 ./helix
	docker build --rm=true -t lukad/helix -f Dockerfile.static .

run: builddocker
	docker run \
	-p 8080:8080 durdn/project-name
