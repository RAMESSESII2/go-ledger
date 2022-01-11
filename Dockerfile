# Official Go image that already has all the tools and packages to compile and run a Go application
FROM golang:1.16-alpine

# This instructs Docker to use this directory as the default destination for all subsequent commands.
WORKDIR /app

COPY . .

# # Download necessary Go modules
# RUN go mod download
RUN CGO_ENABLED=0 go get -v ./...
RUN go env && ls /go/bin
# RUN CGO_ENABLED=0 go build -o /go-ledger-service cmd/server/main.go

# Second stage: start from an empty base image
FROM scratch

# # Copy the binary from the first stage
COPY --from=0 go/bin/server /

# Tell Docker what executable to run by default when starting this container
# CMD ["/go-ledger"]

EXPOSE 9000

ENTRYPOINT ["/server"]

