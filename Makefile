# General
WORKDIR = $(PWD)

# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test

build_linux:
	GOOS=linux GOARCH=amd64 ${GOCMD} build -o cgp_linux ./cmd/cgp
	zip cgp_linux.zip cgp_linux

build_windows:
	GOOS=windows GOARCH=amd64 ${GOCMD} build -o cgp_windows.exe ./cmd/cgp
	zip cgp_windows.zip cgp_windows.exe

build_macos:
	GOOS=darwin GOARCH=amd64 ${GOCMD} build -o cgp_macos ./cmd/cgp
	zip cgp_macos.zip cgp_macos

test:
	$(GOTEST) ./...
