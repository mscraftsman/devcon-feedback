.PHONY: build

build: clean dist/devcon-feedback

dist/devcon-feedback:
	go clean
	if [ ! -d dist ]; then mkdir dist; fi;
	go build -o dist/devcon-feedback

clean:
	go clean
	if [ -a dist/devcon-feedback ]; then rm dist/devcon-feedback; fi;