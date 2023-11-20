TARGET=hello-world
SRC=hello-world.go

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

