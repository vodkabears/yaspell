.PHONY: githooks
githooks:
	cp -f githooks/* .git/hooks/

.PHONY: install
install: githooks
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	go get -u github.com/mitchellh/gox
	go get -u github.com/msoap/go-carpet

.PHONY: lint
lint:
	gometalinter ./... --enable-all --line-length=100 --vendor --tests --sort=path --sort=line --sort=column --deadline=2m

.PHONY: test
test:
	go test -cover ./...

.PHONY: cover
cover:
	go-carpet

.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build: clean
	mkdir build
	gox -os="linux darwin windows" -arch="386 amd64" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}" ./...

.PHONY: run
run:
	go run *.go ${ARGS}
