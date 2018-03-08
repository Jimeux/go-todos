# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ENV GIN_MODE=release
ENV VIEW_DIR=/go/src/gin-todos/views

# Copy the local package files to the container's workspace.
ADD . /go/src/gin-todos/

# Fetch dependencies
RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-xorm/xorm

# Build the app
RUN go install gin-todos

# Run the app
#ENTRYPOINT /go/bin/gin-todos

#EXPOSE 8080
