TARGETS=$(SERVER_TARGET) $(CLIENT_TARGET)
SERVER_TARGET=server
SERVER_SRC=server.go
CLIENT_TARGET=client
CLIENT_SRC=client.go

all: build

build: $(SERVER_TARGET) $(CLIENT_TARGET)

$(SERVER_TARGET): $(SERVER_SRC)
	go build -o $(SERVER_TARGET) $(SERVER_SRC)

$(CLIENT_TARGET): $(CLIENT_SRC)
	go build -o $(CLIENT_TARGET) $(CLIENT_SRC)

run:
	./server&
	sleep 1
	./client

.PHONY: clean
clean:
	rm -f $(TARGETS) *~

