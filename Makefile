GO ?= go

test:
	$(GO) test -v ./...

run:
	$(GO) run .

binary: dist FORCE
	$(GO) version
ifeq ($(OS),Windows_NT)
	$(GO) build -o dist/accountmerging.exe .
else
	$(GO) build -o dist/accountmerging .
endif

dist:
	mkdir $@

FORCE:
