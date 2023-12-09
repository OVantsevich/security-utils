ARG GOLANG_VERSION=1.21-alpine
ARG NODE_VERSION=14

FROM golang:alpine AS tools-builder

RUN apk --no-cache add ca-certificates nmap libxslt

RUN go install github.com/OJ/gobuster/v3@latest

FROM node:${NODE_VERSION} AS react-builder

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH
COPY client/package.json client/package-lock.json ./
RUN npm install --silent
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

RUN go mod vendor
RUN go build -o /bin/main -mod=vendor

FROM tools-builder AS dev

ARG PORT
ENV PORT ${PORT}

RUN mkdir /bin/db

COPY --from=go-builder /bin/main /bin/main
COPY server/wordlist.txt /bin/db/wordlist.txt
COPY --from=react-builder /app/build /bin/static
COPY --from=react-builder /app/build/static /bin/static

EXPOSE 12345

ENTRYPOINT ["/bin/main"]