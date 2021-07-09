BINARY := icity

BUILD_DIR   := build
BUILD_FLAGS := -v

CGO_ENABLED := 0
GO111MODULE := on

LDFLAGS += -w -s -buildid=

GO_BUILD = GO111MODULE=$(GO111MODULE) CGO_ENABLED=$(CGO_ENABLED) \
	go build $(BUILD_FLAGS) -ldflags '$(LDFLAGS)' -trimpath

.PHONY: cli

all: cli

cli:
	$(GO_BUILD) -o $(BUILD_DIR)/$(BINARY)

clean:
	rm -rf $(BUILD_DIR)