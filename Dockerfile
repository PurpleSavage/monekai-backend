# Stage 1: Build (Compilación)
FROM golang:1.22-alpine AS builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar archivos de gestión de dependencias
# (Asegúrate de que existan en tu carpeta antes de compilar)
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación
# CGO_ENABLED=0 genera un binario estático que no depende de librerías externas
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final (Ejecución)
FROM alpine:latest

WORKDIR /root/

# Copiamos solo el binario construido en la etapa anterior
COPY --from=builder /app/main .

# Exponer el puerto donde corre tu app (cámbialo si usas otro)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]