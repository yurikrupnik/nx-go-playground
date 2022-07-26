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
FROM scratch
WORKDIR /

#COPY --from=build /app /app
COPY ./users-api /app

EXPOSE 8080

#USER nonroot:nonroot

ENTRYPOINT ["/app"]
