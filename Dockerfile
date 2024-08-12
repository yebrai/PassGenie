# Configurar el directorio de trabajo
WORKDIR /app

# Copiar el archivo go.mod y go.sum e instalar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente al contenedor
COPY . .

# Construir el binario
RUN go build -o passgenie ./cmd/passgenie

# Etapa de ejecución
FROM debian:bullseye-slim

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/passgenie /usr/local/bin/passgenie

# Configurar el comando predeterminado para ejecutar el binario
ENTRYPOINT ["passgenie"]