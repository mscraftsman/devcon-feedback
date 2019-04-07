VERSION     := 1.0.0
COMMIT      := `git rev-parse HEAD`
DATE        := `date +%FT%T%z`
BUILD_FLAGS := "-X=main.appVersion=$(VERSION) -X=main.appCommit=$(COMMIT) -X=main.appBuilt=$(DATE)"

.PHONY: build

build: clean dist/devcon-feedback

dist/devcon-feedback:
	if [ ! -d dist ]; then mkdir dist; fi;
	go build -ldflags ${BUILD_FLAGS} -o dist/devcon-feedback

clean:
	go clean
	rm dist/devcon-feedback