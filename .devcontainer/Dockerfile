FROM docker.io/golang:1.24-alpine
RUN apk update && apk upgrade
RUN apk --no-cache add curl git
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v golang.org/x/tools/cmd/goimports@latest
RUN go install -v golang.org/x/tools/gopls@latest