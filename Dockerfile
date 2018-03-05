# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ENV GIN_MODE=release
ENV VIEW_DIR=/go/src/gin-test/views
ENV DATABASE_URL="postgres://localhost:5432/xorm_test?user=deployer&password=pass&sslmode=disable"

# Copy the local package files to the container's workspace.
ADD . /go/src/gin-test/

# Fetch dependencies
RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-xorm/xorm

# Build the app
RUN go install gin-test

# Run the app
ENTRYPOINT /go/bin/gin-test

EXPOSE 8080
