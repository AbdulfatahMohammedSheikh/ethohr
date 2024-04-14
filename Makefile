run: 
	@air

setup: 
	@go run ./cmd/setup/main.go

tests : 
	@./run.sh

health:
	@go test ./test/checkhealth_test/checkhealth_test.go

clear :
	@ rm -rf ./tmp/

