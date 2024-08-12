# Nombre del binario
BINARY_NAME=passgenie

# Construir el binario
build:
	go build -o $(BINARY_NAME) ./cmd/passgenie

# Limpiar archivos generados
clean:
	rm -f $(BINARY_NAME)

# Ejecutar las pruebas
test:
	go test ./internal/password/...

# Ejecutar el binario
run:
	./$(BINARY_NAME)
