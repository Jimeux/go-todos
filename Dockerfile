# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ENV APP_NAME=gin-todos
ENV GIN_MODE=release
ENV VIEW_DIR=/go/src/$APP_NAME/public/views
ENV ASSET_DIR=/go/src/$APP_NAME/public/assets

# Fetch dependencies
RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-xorm/xorm
RUN go get github.com/garyburd/redigo/redis
RUN go get github.com/fluent/fluent-logger-golang/fluent

# Copy the local package files to the container's workspace.
ADD . /go/src/$APP_NAME/

# Build the app
RUN go install $APP_NAME

# Run the app
ENTRYPOINT /go/bin/$APP_NAME
