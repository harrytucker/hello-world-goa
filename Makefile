.PHONY: gen
# generate new backend code when you have updated the design
gen:
	goa gen github.com/harrytucker/hello-world-goa/design

# generate handler stubs for new services
stubs:
	goa example github.com/harrytucker/hello-world-goa/design

# build server and cli client binaries
build:
	go build ./cmd/hello_world && go build ./cmd/hello_world-cli

# remove generated code
clean:
	rm -rf gen
	rm -rf cmd
