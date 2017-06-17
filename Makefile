PROJECT := github.com/vodkabears/yaspell
VERSION := $(shell git describe --tags --abbrev=0)
VERSION_FLAG := $(PROJECT)/config.version=$(VERSION)
DIST_FOLDER := dist

.PHONY: githooks
githooks:
	cp -f githooks/* .git/hooks/

.PHONY: install
install: githooks
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	go get -u github.com/mitchellh/gox
	go get -u github.com/msoap/go-carpet
	go get -u github.com/go-playground/overalls
	go get -u github.com/mattn/goveralls
	go get -u github.com/vodkabears/yaspell
	go install ./...

.PHONY: lint
lint:
	yaspell -opts=IGNORE_UPPERCASE,IGNORE_DIGITS,IGNORE_URLS -dict=.yaspellignore README.md
	gometalinter ./... --enable-all --line-length=100 --vendor --sort=path --sort=line --sort=column --deadline=5m -t -j 1

.PHONY: test
test:
	go test -ldflags "-X ${VERSION_FLAG}" -cover ./...

.PHONY: cover
cover:
	go-carpet

.PHONY: clean
clean:
	rm -rf ${DIST_FOLDER}

.PHONY: build
build: clean
	gox -ldflags "-X ${VERSION_FLAG}" -os="linux darwin windows" -arch="386 amd64" -output="${DIST_FOLDER}/{{.Dir}}_{{.OS}}_{{if eq .Arch \"386\"}}i386{{else}}{{.Arch}}{{end}}" ./...

.PHONY: run
run:
	go run -ldflags "-X ${VERSION_FLAG}" *.go ${ARGS}

.PHONY: coveralls
coveralls:
	overalls -project=$(PROJECT) -covermode=count
	goveralls -coverprofile=overalls.coverprofile -service=travis-ci -repotoken 2zEPspDOcwDzRYBKbDmjk846HOf4ugxlO
