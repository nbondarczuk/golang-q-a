TARGET=workers
SRC=workers.go

all: $(TARGET)

$(TARGET): $(SRC)
	go build -o $(TARGET) $(SRC)

run:
	go run $(SRC)

vet:
	go vet $(SRC)

.PHONY: clean
clean:
	rm -f $(TARGET) *~

