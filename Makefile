GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY_NAME=auth-test
SECOND_BINARY_NAME=auth-provider

build: VER=$(shell git rev-parse --short HEAD)
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
	$(GOBUILD) -o $(SECOND_BINARY_NAME) -v auth/main.go

clean:
	$(GOCLEAN)

build-lambda:
	CGO_ENABLED=0 GOOS=linux make build

deploy-lambda-beta: build-lambda
	sls deploy --aws-profile my --stage beta --region us-east-2
	make clean

deploy-lambda-prod: build-lambda
	sls deploy --aws-profile my --stage prod --region us-east-2
	make clean