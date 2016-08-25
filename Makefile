GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
LINTIGNOREDEPS='vendor/.+\.go'
TARGET_ONLY_PKGS=$(shell go list ./... 2> /dev/null | grep -v "/misc/" | grep -v "/vendor/")
INTERNAL_BIN=.bin
HAVE_GLIDE:=$(shell which glide)
HAVE_GOLINT:=$(shell which golint)
HAVE_GOTEMPLATES:=$(shell which templates)

.PHONY: unit generate lint vet test golint gotemplates install-deps

init: install-deps

unit: generate lint vet test

generate: gotemplates
	go generate .

lint: golint
	@echo "Invoking linter"
	@lint=`golint ./...`; \
	lint=`echo "$$lint" | grep -E -v -e ${LINTIGNOREDEPS}`; \
	echo "$$lint"; \
	if [ "$$lint" != "" ]; then exit 1; fi

vet:
	@echo "Invoking vet"
	@go tool vet -all -structtags -shadow $(shell ls -d */ | grep -v "misc" | grep -v "vendor")

test:
	@echo "Invoking test"
	@go test $(TARGET_ONLY_PKGS)

install-deps: glide
	@echo "Installing all dependencies"
	@PATH=$(INTERNAL_BIN):$(PATH) glide i

golint:
ifndef HAVE_GOLINT
	@echo "Installing linter"
	@go get -u github.com/golang/lint/golint
endif

gotemplates:
ifndef HAVE_GOTEMPLATES
	@echo "Installing templates"
	@go get -u github.com/cyberdelia/templates
endif

glide:
ifndef HAVE_GLIDE
	@echo "Installing glide"
	@mkdir -p $(INTERNAL_BIN)
	@wget -q -O - https://github.com/Masterminds/glide/releases/download/v0.11.1/glide-v0.11.1-$(GOOS)-$(GOARCH).tar.gz | tar xvz
	@mv $(GOOS)-$(GOARCH)/glide $(INTERNAL_BIN)/glide
	@rm -rf $(GOOS)-$(GOARCH)
endif

