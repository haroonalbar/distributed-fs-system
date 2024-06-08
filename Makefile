# @ at the beginning to not show the commands running
build:
	@go build -o ./bin/dfs

# added build as dependancy
run: build
	@./bin/dfs

# ./... will look for test files in subdirs recursivly and -v for verbose
test:
	@go test ./... -v
