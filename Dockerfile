# Use a imagem oficial do Go como base
FROM golang:1.20

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar o go.mod e o go.sum (se existirem)
COPY go.mod go.sum ./

# Baixar as dependências (caso use go.mod)
RUN go mod download

# Copiar o restante do código
COPY . .

# Build da aplicação
RUN go build -o main .

# Definir a porta que o container vai expor
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]
