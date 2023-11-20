ARG GOLANG_VERSION=1.21-alpine

FROM golang:${GOLANG_VERSION} AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY . .

RUN go install github.com/OJ/gobuster/v3@latest

RUN go mod vendor
RUN go build -o /bin/main -mod=vendor

FROM scratch AS dev

ARG PORT
ENV PORT ${PORT}

COPY --from=build /bin/main /bin/main

EXPOSE ${PORT}
ENTRYPOINT ["/bin/main"]
