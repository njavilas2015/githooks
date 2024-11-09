# Usa una imagen base oficial de Go
FROM golang:1.20

# Configura el directorio de trabajo en el contenedor
WORKDIR /workspace

# Instala dependencias adicionales si es necesario (por ejemplo, herramientas de desarrollo)
RUN apt-get update && apt-get install -y git

# Establece el contenedor como no interactivo
ENV DEBIAN_FRONTEND=noninteractive

# Configura Go Modules
ENV GO111MODULE=on

# Abre el puerto para la depuración (opcional)
EXPOSE 8080
