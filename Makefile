# General
WORKDIR = $(PWD)

# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test
MARP_VER = v2.4.0

build_linux:
	wget https://github.com/marp-team/marp-cli/releases/download/${MARP_VER}/marp-cli-${MARP_VER}-linux.tar.gz -O ./assets/marp-cli.tar.gz
	tar -xvf ./assets/marp-cli.tar.gz -C ./assets
	chmod +x ./assets/marp
	GOOS=linux GOARCH=amd64 ${GOCMD} build -o cgp_linux ./cmd/cgp
	zip cgp_linux.zip cgp_linux

build_windows:
	wget https://github.com/marp-team/marp-cli/releases/download/${MARP_VER}/marp-cli-${MARP_VER}-windows.tar.gz -O ./assets/marp-cli.tar.gz
	tar -xvf ./assets/marp-cli.tar.gz -C ./assets
	chmod +x ./assets/marp
	GOOS=windows GOARCH=amd64 ${GOCMD} build -o cgp_windows ./cmd/cgp
	zip cgp_windows.zip cgp_windows

build_macos:
	wget https://github.com/marp-team/marp-cli/releases/download/${MARP_VER}/marp-cli-${MARP_VER}-mac.tar.gz -O ./assets/marp-cli.tar.gz
	tar -xvf ./assets/marp-cli.tar.gz -C ./assets
	chmod +x ./assets/marp
	GOOS=darwin GOARCH=amd64 ${GOCMD} build -o cgp_macos ./cmd/cgp
	zip cgp_macos.zip cgp_macos

test:
	$(GOTEST) ./...
