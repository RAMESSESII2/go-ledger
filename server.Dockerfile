# Official Go image that already has all the tools and packages to compile and run a Go application
FROM golang:1.16-alpine

# This instructs Docker to use this directory as the default destination for all subsequent commands.
WORKDIR /app


# # Download necessary Go modules
COPY ./go.mod go.sum ./
RUN go mod download && go mod verify
# RUN CGO_ENABLED=0 go get -v ./...
# RUN go env && ls /go/bin
# RUN CGO_ENABLED=0 go build -o /go-ledger-service cmd/server/main.go
RUN go get github.com/githubnemo/CompileDaemon

# Second stage: start from an empty base image
# FROM scratch

# # Copy the binary from the first stage
# COPY --from=0 go/bin/server /
COPY . .
COPY ./entrypoint.sh /entrypoint.sh

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for

RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

# Tell Docker what executable to run by default when starting this container
# CMD ["/go-ledger"]
# ENTRYPOINT ["/server"]
ENTRYPOINT [ "sh", "/entrypoint.sh" ]

RUN go env && ls /go/bin