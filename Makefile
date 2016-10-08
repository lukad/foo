default: builddocker

setup:
	go get bitbucket.org/chrj/smtpd
	go get github.com/gorilla/mux
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/Sirupsen/logrus

js:
	cd ./go/src/github.com/lukad/helix/frontend; node_modules/brunch/bin/brunch build -p

assets: js
	go-bindata -o ./go/src/github.com/lukad/helix/web/assets/assets.go \
		-pkg assets \
		-prefix ./go/src/github.com/lukad/helix/frontend/public \
		/go/src/github.com/lukad/helix/frontend/public

buildgo: assets
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
