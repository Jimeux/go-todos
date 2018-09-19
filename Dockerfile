# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.11

ENV APP_NAME=github.com/Jimeux/go-todos
ENV SRC_ROOT=/go/src/$APP_NAME
ENV GIN_MODE=release
ENV VIEW_DIR=$SRC_ROOT/public/views
ENV ASSET_DIR=$SRC_ROOT/public/assets

ENV GO111MODULE=on

WORKDIR $SRC_ROOT

# Copy the local package files to the container's workspace.
COPY . .

# Build the app
RUN go install $SRC_ROOT

# Run the app
ENTRYPOINT /go/bin/go-todos
