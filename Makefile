gobuild:
	@go build -v -o paper-id-take-home-test cmd/main.go

gorun:
	make gobuild
	@./paper-id-take-home-test
