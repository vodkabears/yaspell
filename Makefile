.PHONY: install
install:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	go get -u github.com/mitchellh/gox

.PHONY: githooks
githooks:
	cp -f githooks/* .git/hooks/

.PHONY: lint
lint:
	gometalinter ./... --enable-all --disable=lll --vendor --tests --sort=path --sort=line --sort=column --deadline=1m

.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build: clean
	mkdir build
	gox -os="linux darwin windows" -arch="386 amd64" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}" ./...

.PHONY: cmd
cmd:
	go run ./cmd/**/*.go ${ARGS}
