.PHONY: gen
gen:
	goa gen github.com/harrytucker/hello-world-goa/design

stubs:
	goa example github.com/harrytucker/hello-world-goa/design