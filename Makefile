BIN_DIR := bin
SRC_DIR := cmd

# Targets (executables to build)
TARGETS := client proxy server

CLIENT_SRC := ./$(SRC_DIR)/client
PROXY_SRC := ./$(SRC_DIR)/proxy
SERVER_SRC := ./$(SRC_DIR)/server

.PHONY: all
all: $(TARGETS)

client:
	go build -o $(BIN_DIR)/client $(CLIENT_SRC)

proxy:
	go build -o $(BIN_DIR)/proxy $(PROXY_SRC)

server:
	go build -o $(BIN_DIR)/server $(SERVER_SRC)

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: run-client run-proxy run-server
run-client:
	$(BIN_DIR)/client

run-proxy:
	$(BIN_DIR)/proxy

run-server:
	$(BIN_DIR)/server