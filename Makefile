GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
LINTIGNOREDEPS='vendor/.+\.go'
TARGET_ONLY_PKGS=$(shell go list ./... 2> /dev/null | grep -v "/misc/" | grep -v "/vendor/")
INTERNAL_BIN=.bin
HAVE_GOLINT:=$(shell which golint)
HAVE_GOCYCLO:=$(shell which gocyclo)
HAVE_GOTEMPLATES:=$(shell which templates)

.PHONY: unit generate lint vet cyclo test golint gocyclo gotemplates

unit: generate lint vet cyclo test

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

cyclo: gocyclo
	@echo "Collecting cyclomatic complexities"
	@gocyclo -over 30 .

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

gocyclo:
ifndef HAVE_GOCYCLO
	@echo "Installing gocyclo"
	@go get -u github.com/fzipp/gocyclo
endif

gotemplates:
ifndef HAVE_GOTEMPLATES
	@echo "Installing templates"
	@go get -u github.com/cyberdelia/templates
endif
