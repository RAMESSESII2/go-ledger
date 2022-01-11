wait-for "db:3307" -- "$@"

# Watch your .go files and invoke go build if the files changed.
# CompileDaemon --build="RUN CGO_ENABLED=0 go get -v ./..."  --command=./go/bin/server
CompileDaemon --build="go build -o bin/server cmd/server/main.go"  --command=./bin/server