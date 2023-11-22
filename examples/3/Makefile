TARGETS=$(SERVER_TARGET)
SERVER_TARGET=server
SERVER_SRC=server.go

all: build

build: $(TARGETS)

$(SERVER_TARGET): $(SERVER_SRC)
	go build -o $(SERVER_TARGET) $(SERVER_SRC)

run:
	./server&

integration-test: build run wait integration-test-json2xml integration-test-xml2json stop

wait:
	@sleep 1

stop:
	@killall server

integration-test-json2xml:
	@printf '{"id":1,"name":"xxx","email":"yyy"} -> '
	@curl http://localhost/json2xml -d '{"id":1,"name":"xxx","email":"yyy"}'
	@echo

integration-test-xml2json:
	@printf '<User><ID>1</ID><Name>xxx</Name><Email>yyy</Email></User> -> '
	@curl http://localhost/xml2json -d '<User><ID>1</ID><Name>xxx</Name><Email>yyy</Email></User>'
	@echo

unit-test:
	go test -v -cover -count=1 .

.PHONY: clean
clean:
	rm -f $(TARGETS) *~

