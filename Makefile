.DEFAULT_GOAL := build

build:
	make clean
	go build -o bin/concur

clean:
	rm -f bin/concur
	rm -f bin/concur.db
	rm -f bin/config.yaml