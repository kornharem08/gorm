FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod files
COPY go.mod ./
COPY cmd/app/go.mod ./cmd/app/

# Copy source code
COPY . .

# Build the application
WORKDIR /app/cmd/app
RUN go mod tidy
RUN go build -o /app/main .

# Create runtime container
FROM mcr.microsoft.com/mssql/server:2022-latest AS db
ENV ACCEPT_EULA=Y
ENV SA_PASSWORD=YourPassword
ENV MSSQL_PID=Express

# Create the application container
FROM golang:1.20-alpine

WORKDIR /app

# Copy built executable
COPY --from=builder /app/main .

# Set the connection string environment variable
ENV DB_CONNECTION_STRING="sqlserver://sa:YourPassword@db:1433?database=testdb"

EXPOSE 8080

CMD ["./main"] 