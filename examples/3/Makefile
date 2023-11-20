SERVER_TARGET=server
SERVER_SRC=server.go

TARGETS=$(SERVER_TARGET)

all: build

build: $(SERVER_TARGET)

$(SERVER_TARGET): $(SERVER_SRC)
	go build -o $(SERVER_TARGET) $(SERVER_SRC)

run:
	./server&
	curl -X POST http://localhost/json -d ''
	curl -X POST http://localhost/xml -d ''

.PHONY: clean
clean:
	rm -f $(TARGETS) *~

