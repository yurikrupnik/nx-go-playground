
# Done!
FROM node:16-alpine AS nest-builder
WORKDIR /app

ARG DIST_PATH
RUN test -n "$DIST_PATH" || (echo "DIST_PATH  not set" && false)

ENV NODE_ENV=$NODE_ENV
COPY ./$DIST_PATH .
RUN npm install --production
ENV PORT=3333
EXPOSE ${PORT}

CMD ["node", "main.js"]

# Done
FROM nginx:alpine AS web-builder
WORKDIR /app
ARG DIST_PATH
RUN test -n "$DIST_PATH" || (echo "DIST_PATH  not set" && false)
COPY ./$DIST_PATH /usr/share/nginx/html
ENV PORT=80
EXPOSE ${PORT}
CMD ["nginx", "-g", "daemon off;"]

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
