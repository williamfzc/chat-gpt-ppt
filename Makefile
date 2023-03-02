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
	${GOCMD} build ./cmd/cgp

test:
	$(GOTEST) ./...
