GENERATOR_HTTP_ADDR := localhost:10000
GENERATOR_GRPC_ADDR := localhost:10001
GENERATOR_DEBUG_ADDR := localhost:10002
PRINTER_HTTP_ADDR := localhost:11000
PRINTER_GRPC_ADDR := localhost:11001
PRINTER_DEBUG_ADDR := localhost:11002

run:
	go run generator/generator-service/generator-server/server_main.go \
		-http.addr $(GENERATOR_HTTP_ADDR) \
		-grpc.addr $(GENERATOR_GRPC_ADDR) \
		-debug.addr $(GENERATOR_DEBUG_ADDR) \
	& go run printer/printer-service/printer-server/server_main.go \
		-http.addr $(PRINTER_HTTP_ADDR) \
		-grpc.addr $(PRINTER_GRPC_ADDR) \
		-debug.addr $(PRINTER_DEBUG_ADDR)

stop:
	-pkill -TERM -f -- "-http.addr $(GENERATOR_HTTP_ADDR)"
	-pkill -TERM -f -- "-http.addr $(PRINTER_HTTP_ADDR)"

.PHONY: run stop

