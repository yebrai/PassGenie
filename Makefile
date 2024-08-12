# Nombre del binario
BINARY_NAME=passgenie
# Nombre de la imagen de Docker
DOCKER_IMAGE=passgenie:latest

# Construir el binario
build:
	go build -o $(BINARY_NAME) ./cmd/passgenie

clean:
	rm -f $(BINARY_NAME)


# Ejecutar las pruebas
test:
	go test ./internal/password/...
	go test ./internal/utils/...

# Construir la imagen de Docker
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Ejecutar el contenedor de Docker
docker-run:
	docker run --rm $(DOCKER_IMAGE)
