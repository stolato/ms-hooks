# --- BUILD STAGE ---
FROM golang:1.24.2 AS builder

# Define o diretório de trabalho dentro do contêiner de build
WORKDIR /app

# Copia os arquivos de módulo para que possam ser cacheados e baixados primeiro
COPY go.mod go.sum ./

# Baixa as dependências. Isso aproveita o cache do Docker.
# Se go.mod/go.sum não mudarem, esta etapa não será executada novamente.
RUN go mod download

# Copia o restante do código-fonte
COPY . .

# Compila a aplicação Go.
# 'CGO_ENABLED=0' é importante para criar um binário estático,
# o que o torna independente de bibliotecas C dinâmicas e adequado para imagens "scratch".
# 'GOOS=linux' e 'GOARCH=arm' são mantidos, assumindo que seu ambiente de destino é ARM Linux.
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# --- FINAL STAGE ---
# Usamos uma imagem base minúscula para o binário final.
# 'scratch' é a menor imagem possível, não contém nada além do que você adiciona.
FROM scratch

# Copia o binário compilado do estágio 'builder' para a imagem final
COPY --from=builder /app/main ./main

# Expor a porta que a aplicação vai escutar
EXPOSE 8081

# Comando para executar a aplicação quando o contêiner iniciar
CMD ["./main"]