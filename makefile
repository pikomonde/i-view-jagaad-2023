.SILENT:test
.SILENT:run-cli

test:
	go test ./... -coverprofile cover.out.temp && cat cover.out.temp | grep -v 'mock_*\|\.pb\.' > cover.out && go tool cover -func cover.out
run-cli:
	go run cmd/cli/app.go
