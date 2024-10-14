# Etapa de build
FROM golang:1.21-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de dependências Go para o container
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia o código-fonte da aplicação Go
COPY *.go ./

# Compila a aplicação
RUN go build -o /my-go-app

# Etapa final (imagem leve)
FROM alpine:latest

# Define o diretório de trabalho da imagem final
WORKDIR /root/

# Copia a aplicação compilada da etapa de build
COPY --from=builder /my-go-app .

# Expõe a porta 8080 para o serviço REST
EXPOSE 8080

# Define o comando padrão para iniciar a aplicação
CMD ["./my-go-app"]
