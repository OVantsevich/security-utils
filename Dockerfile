ARG GOLANG_VERSION=1.21-alpine
ARG NODE_VERSION=14

FROM alpine AS tools-builder

RUN apk --no-cache add ca-certificates nmap libxslt
RUN apk add --update --no-cache vim git make musl-dev go curl
RUN export GOPATH=/root/go
RUN export PATH=${GOPATH}/bin:/usr/local/go/bin:$PATH
RUN export GOBIN=$GOROOT/bin
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
RUN export GO111MODULE=on
RUN go version
RUN go install github.com/OJ/gobuster/v3@latest

FROM node:${NODE_VERSION} AS react-builder

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH
COPY client/package.json client/package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent
RUN npm install react-modal
COPY client/src ./src
COPY client/public ./public

RUN npm run build

FROM golang:${GOLANG_VERSION} AS go-builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY server .
RUN go mod download


COPY --from=react-builder /app/build ./static

RUN go mod vendor
RUN go build -o /bin/main -mod=vendor

FROM tools-builder AS dev

ARG PORT
ENV PORT ${PORT}

COPY --from=go-builder /bin/main /bin/main
COPY --from=go-builder /app/static /bin/static

EXPOSE 12345
EXPOSE 3000

ENTRYPOINT ["/bin/main"]
