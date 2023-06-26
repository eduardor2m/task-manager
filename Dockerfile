# Use a imagem oficial do Golang como base
FROM golang:1.20.4

# Define o diretório de trabalho dentro do container
WORKDIR /go/src/app

# Copia o código fonte para o diretório de trabalho
COPY . .

# Baixa as dependências do Go
RUN go mod download

# Compila o código Go
RUN go build -o app ./src/api/app/main.go

# Expõe a porta que a aplicação Go escuta
EXPOSE 9090

# Comando para executar a aplicação quando o container for iniciado
CMD ["./app"]
