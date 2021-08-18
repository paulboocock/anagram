.PHONY: gox cli container format lint tidy test-setup test container-release clean

# -----------------------------------------------------------------------------
#  CONSTANTS
# -----------------------------------------------------------------------------

version = `cat VERSION`

go_dirs = `go list ./...`

build_dir  = build
output_dir = $(build_dir)/output

coverage_dir  = $(build_dir)/coverage
coverage_out  = $(coverage_dir)/coverage.out
coverage_html = $(coverage_dir)/coverage.html

container_name = paulboocock/anagram

# -----------------------------------------------------------------------------
#  BUILDING
# -----------------------------------------------------------------------------

gox:
	go get -u github.com/mitchellh/gox
	mkdir -p $(output_dir)

cli: gox
	gox -osarch="linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64" -output=$(output_dir)/anagram_{{.OS}}_{{.Arch}} ./cmd/cli/

container: cli
	docker build -t $(container_name):$(version) .

# -----------------------------------------------------------------------------
#  FORMATTING
# -----------------------------------------------------------------------------

format:
	go fmt $(go_dirs)
	gofmt -s -w .

lint:
	go get -u golang.org/x/lint/golint
	golint $(go_dirs)

tidy:
	go mod tidy

# -----------------------------------------------------------------------------
#  TESTING
# -----------------------------------------------------------------------------

test-setup:
	mkdir -p $(coverage_dir)
	go get -u golang.org/x/tools/cmd/cover

test: test-setup
	go test $(go_dirs) -v -short -covermode=count -coverprofile=$(coverage_out)
	go tool cover -html=$(coverage_out) -o $(coverage_html)
	go tool cover -func=$(coverage_out)

# -----------------------------------------------------------------------------
#  CLEANUP
# -----------------------------------------------------------------------------

clean:
	rm -rf $(build_dir)