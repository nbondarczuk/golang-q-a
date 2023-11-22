TARGET=workers
SRC=workers.go
RFLAGS=-race

all: $(TARGET)

$(TARGET): $(SRC)
	go build -o $(TARGET) $(SRC)

run:
	go run $(RFLAGS) $(SRC)

vet:
	go vet $(SRC)

.PHONY: clean
clean:
	rm -f $(TARGET) *~

