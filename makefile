.SILENT:test
.SILENT:run-cli
.SILENT:generate

test:
	go test ./... -coverprofile cover.out.temp && cat cover.out.temp | grep -v 'mock_*\|\.pb\.' > cover.out && go tool cover -func cover.out
	go tool cover -html=cover.out -o cover.out.html
	go tool cover -html=cover.out
run-cli:
	go run cmd/cli/app.go
generate:
	go generate ./...
