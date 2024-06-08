# @ at the beginning to not show the commands running
build:
	@go build

# added build as dependancy
run: build
	@./distributed-fs-system

# ./... will look for test files in subdirs recursivly and -v for verbose
test:
	@go test ./... -v
