install:
	go get github.com/aws/aws-lambda-go/events
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/boombuler/barcode

build: install
	mkdir -p functions/bin
	@for dir in `ls functions/src`; do \
		go build -o functions/bin/$$dir ./functions/src/$$dir; \
	done