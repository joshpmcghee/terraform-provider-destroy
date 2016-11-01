default: build

deps:
	go install github.com/hashicorp/terraform

build:
	go build -o terraform-provider-destroy .

install:
	go build -o /usr/local/bin/terraform-provider-destroy .
	@echo "Ensure that you move the provided terraformrc to ~/.terraformrc or update your own with the provider mapping."
