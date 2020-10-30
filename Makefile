APP=bell
MODULE := github.com/branislavlazic/bell
VERSION := v0.1-alpha

.PHONY: clean binaries test

all: clean test zip

test:
	go test -count=1 -v ./...

clean:
	rm -rf binaries release

zip: release/$(APP)_$(VERSION)_osx_x86_64.tar.gz release/$(APP)_$(VERSION)_windows_x86_64.zip release/$(APP)_$(VERSION)_linux_x86_64.tar.gz release/$(APP)_$(VERSION)_windows_x86_32.zip release/$(APP)_$(VERSION)_linux_x86_32.tar.gz

binaries: binaries/osx_x86_64/$(APP) binaries/windows_x86_64/$(APP).exe binaries/linux_x86_64/$(APP) binaries/windows_x86_32/$(APP).exe binaries/linux_x86_32/$(APP)

release/$(APP)_$(VERSION)_osx_x86_64.tar.gz: binaries/osx_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_osx_x86_64.tar.gz -C binaries/osx_x86_64 $(APP)
	tar cfz release/$(APP)-repl_$(VERSION)_osx_x86_64.tar.gz -C binaries/osx_x86_64 $(APP)-repl

binaries/osx_x86_64/$(APP):
	GOOS=darwin GOARCH=amd64 go build -o binaries/osx_x86_64/$(APP) $(MODULE)/cmd/bell
	GOOS=darwin GOARCH=amd64 go build -o binaries/osx_x86_64/$(APP)-repl $(MODULE)/cmd/repl

release/$(APP)_$(VERSION)_windows_x86_64.zip: binaries/windows_x86_64/$(APP).exe
	mkdir -p release
	cd ./binaries/windows_x86_64 && zip -r -D ../../release/$(APP)_$(VERSION)_windows_x86_64.zip $(APP).exe
	cd ./binaries/windows_x86_64 && zip -r -D ../../release/$(APP)-repl_$(VERSION)_windows_x86_64.zip $(APP)-repl.exe

binaries/windows_x86_64/$(APP).exe:
	GOOS=windows GOARCH=amd64 go build -o binaries/windows_x86_64/$(APP).exe $(MODULE)/cmd/bell
	GOOS=windows GOARCH=amd64 go build -o binaries/windows_x86_64/$(APP)-repl.exe $(MODULE)/cmd/repl

release/$(APP)_$(VERSION)_linux_x86_64.tar.gz: binaries/linux_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_x86_64.tar.gz -C binaries/linux_x86_64 $(APP)
	tar cfz release/$(APP)-repl_$(VERSION)_linux_x86_64.tar.gz -C binaries/linux_x86_64 $(APP)-repl

binaries/linux_x86_64/$(APP):
	GOOS=linux GOARCH=amd64 go build -o binaries/linux_x86_64/$(APP) $(MODULE)/cmd/bell
	GOOS=linux GOARCH=amd64 go build -o binaries/linux_x86_64/$(APP)-repl $(MODULE)/cmd/repl

release/$(APP)_$(VERSION)_windows_x86_32.zip: binaries/windows_x86_32/$(APP).exe
	mkdir -p release
	cd ./binaries/windows_x86_32 && zip -r -D ../../release/$(APP)_$(VERSION)_windows_x86_32.zip $(APP).exe
	cd ./binaries/windows_x86_32 && zip -r -D ../../release/$(APP)-repl_$(VERSION)_windows_x86_32.zip $(APP)-repl.exe

binaries/windows_x86_32/$(APP).exe:
	GOOS=windows GOARCH=386 go build -o binaries/windows_x86_32/$(APP).exe $(MODULE)/cmd/bell
	GOOS=windows GOARCH=386 go build -o binaries/windows_x86_32/$(APP)-repl.exe $(MODULE)/cmd/repl

release/$(APP)_$(VERSION)_linux_x86_32.tar.gz: binaries/linux_x86_32/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_x86_32.tar.gz -C binaries/linux_x86_32 $(APP)
	tar cfz release/$(APP)-repl_$(VERSION)_linux_x86_32.tar.gz -C binaries/linux_x86_32 $(APP)-repl

binaries/linux_x86_32/$(APP):
	GOOS=linux GOARCH=386 go build -o binaries/linux_x86_32/$(APP) $(MODULE)/cmd/bell
	GOOS=linux GOARCH=386 go build -o binaries/linux_x86_32/$(APP)-repl $(MODULE)/cmd/repl