FROM golang:1.19

# Copy application data into image
COPY . go/src/service-users-data
WORKDIR go/src/service-users-data

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Install our third-party application for hot-reloading capability.
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build ./cmd/service-users-data" -command="./service-users-data" -directory="."
