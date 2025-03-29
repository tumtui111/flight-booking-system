flight-book-system:
	go run main.go flight-book-system


prepare:
	@echo "== install preparation =="
	brew install go
	brew install golangci-lint
	brew upgrade golangci-lint

	@echo "== install gomock =="
	go install github.com/golang/mock/mockgen@latest

ci.lint:
	@echo "== ci.linter =="
	golangci-lint run -v ./... --fix

gen-mock:
	mockery --all --case snake --disable-version-string --keeptree

test:
	go test ./service/... -coverprofile coverage.out -tags dynamic
