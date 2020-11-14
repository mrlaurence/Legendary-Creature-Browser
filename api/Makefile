GO := go
GO_BUILD := $(GO) build

MAIN := cmd/main/main.go
EXEC := api

build:
	$(GO_BUILD) -v -o $(EXEC) $(MAIN)

$(EXEC): build

clean:
	rm $(EXEC)

run: $(EXEC) config.json
	./$(EXEC)

config.json:
	cp default-$@ $@

.PHONY: build clean run