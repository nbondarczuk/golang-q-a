TARGET=gocrur
SRC=gocrur.go
DB=user.db

all: build

build: $(TARGET)

$(TARGET): $(SRC)
	go build -o $(TARGET) $(SRC)

run: build
	./$(TARGET)

install:
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite

.PHONY: clean
clean:
	rm -f $(TARGET) $(DB) *~

