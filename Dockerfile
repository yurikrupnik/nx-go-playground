
# Done!
FROM node:18-alpine AS nest-builder
WORKDIR /app

ARG DIST_PATH
RUN test -n "$DIST_PATH" || (echo "DIST_PATH  not set" && false)

ENV NODE_ENV=$NODE_ENV
COPY ./$DIST_PATH .
RUN npm install --omit=dev
ENV PORT=3333
EXPOSE ${PORT}

CMD ["node", "main.js"]

# Done
FROM nginx:alpine AS web-builder
#FROM haproxy:alpine AS web-builder
WORKDIR /app
ARG DIST_PATH
RUN test -n "$DIST_PATH" || (echo "DIST_PATH  not set" && false)
COPY ./$DIST_PATH /usr/share/nginx/html
#COPY ./apps/users/client/k8s/base/haproxy.cfg /etc/haproxy/haproxy.cfg
ENV PORT=80
EXPOSE ${PORT}
CMD ["nginx", "-g", "daemon off;"]
#CMD ["haproxy", "start"]

# Done
FROM alpine:latest AS go-builder
WORKDIR /
ARG DIST_PATH
RUN test -n "$DIST_PATH" || (echo "DIST_PATH  not set" && false)
ARG ENTRY_NAME=app
ENV PORT=8080
EXPOSE ${PORT}
COPY $DIST_PATH ./app
EXPOSE 8080
ENTRYPOINT ["/app"]

#FROM golang:1.18-buster AS go-build
#WORKDIR /app
#COPY workspace.json ./
#COPY nx.json ./
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#ARG APP_PATH
#RUN test -n "$APP_PATH" || (echo "DIST_PATH  not set" && false)
#COPY ./$APP_PATH ./$APP_PATH
#COPY ./libs/go ./libs/go
##RUN #go build -a -o app $APP_PATH/main.go
#RUN npx nx run users-api:build
#RUN ls
#RUN ls
#
###
### Deploy
###
##34.7MB
##FROM gcr.io/distroless/base-debian10
##24.4MB
#FROM alpine:latest AS go-alpine
#WORKDIR /
#
#COPY --from=go-build /app /app
##COPY ./fiber-mongo /fiber-mongo
#
#EXPOSE 8080
#
##USER nonroot:nonroot
#
#ENTRYPOINT ["/app"]
