#FROM golang:1.18-buster AS build
#
#WORKDIR /app
#
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#
#COPY ./apps/users/api .
#
#RUN go build -o /app

##
## Deploy
##
#34.7MB
#FROM gcr.io/distroless/base-debian10
#24.4MB
# with goreleaser 18MB
#FROM alpine:latest AS go-alpine
#12.7MB
FROM scratch AS go-builder
WORKDIR /
ARG DIST_PATH
RUN test -n "$DIST_PATH" || (echo "DIST_PATH  not set" && false)
ARG ENTRY_NAME=app
ENV PORT=8080
EXPOSE ${PORT}
COPY $DIST_PATH ./app
EXPOSE ${PORT}
ENTRYPOINT ["/app"]
