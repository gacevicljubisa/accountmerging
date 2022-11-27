GO ?= go
FILE_PATH ?= accounts.json

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

merge:
	./dist/accountmerging $(FILE_PATH)

FORCE:
