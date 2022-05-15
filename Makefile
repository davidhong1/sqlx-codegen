.PHONY: test
test:
	CGO_ENABLED=0 go build . && mv sqlx-codegen test/ && cd test/ && ./sqlx-codegen -t User && rm ./sqlx-codegen