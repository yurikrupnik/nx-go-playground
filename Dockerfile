#FROM ubuntu:latest AS ubuntu
#RUN curl -O https://storage.googleapis.com/golang/go1.13.5.linux-amd64.tar.gz
#RUN curl curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
#RUN apk add nodejs-current
#RUN apk update && apk add git
# Install dependencies only when needed
#CMD git checkout

#FROM node:16 AS deps
##RUN #yum install git
## Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
##RUN #apk add --no-cache libc6-compat
#WORKDIR /app
#COPY package*.json ./
#RUN yarn
#
### Rebuild the source code only when needed
#FROM node:16-alpine AS builder
#WORKDIR /app
#COPY --from=deps /app/node_modules ./node_modules
#COPY --from=deps /app/package*.json ./
#
##COPY . .
#COPY nx.json ./
#COPY workspace.json ./
##COPY .eslintignore ./
#COPY .eslintrc.json ./
#COPY ./vite.config.js ./
#COPY ./Makefile ./
#COPY babel.config.json ./
#COPY tsconfig.base.json ./
#COPY jest.*.js ./
###
#COPY apps ./apps
#COPY libs ./libs
##
#RUN yarn test
#RUN yarn build
#RUN yarn lint

#FROM node:16-alpine AS build-assets
#WORKDIR /app
#COPY --from=builder /app/package*.json ./
#COPY --from=builder /app/dist ./
#COPY --from=builder /app/dist ./shit
#RUN yarn docker
#COPY --from=deps /app/package*.json ./
#
#
## todo change to builds!
#FROM node:16-alpine AS runner
#WORKDIR /app
##
#ENV NODE_ENV production
#COPY --from=builder /app/dist ./dist


# Done!
FROM node:16-alpine AS nest-builder
WORKDIR /app

ARG DIST_PATH
RUN test -n "$DIST_PATH" || (echo "DIST_PATH  not set" && false)

ENV NODE_ENV=$NODE_ENV
COPY ./$DIST_PATH .
RUN npm install
ENV PORT=3333
EXPOSE ${PORT}

CMD ["node", "main.js"]

# Done
FROM nginx:alpine AS web-builder
WORKDIR /app
ARG DIST_PATH
RUN test -n "$DIST_PATH" || (echo "DIST_PATH  not set" && false)
COPY ./$DIST_PATH /usr/share/nginx/html
#COPY ./apps/users/client/nginx-config/default.conf /etc/nginx/conf.d/default.conf
ENV PORT=80
EXPOSE ${PORT}
CMD ["nginx", "-g", "daemon off;"]

# Done
FROM gcr.io/distroless/base-debian10 AS go-builder
#FROM golang:1.18 AS go-builder
WORKDIR /
# not using it yet
ARG DIST_PATH
ARG ENTRY_NAME=app
COPY $DIST_PATH ./app
USER nonroot:nonroot
ENTRYPOINT ["/app"]
ENV PORT=8080
EXPOSE ${PORT}
#
#RUN addgroup -g 1001 -S nodejs
#RUN adduser -S nextjs -u 1001
#
## You only need to copy next.config.js if you are NOT using the default configuration
## COPY --from=builder /app/next.config.js ./
#COPY --from=builder /app/public ./public
#COPY --from=builder /app/package.json ./package.json
#
## Automatically leverage output traces to reduce image size
## https://nextjs.org/docs/advanced-features/output-file-tracing
#COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
#COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static
#
#USER nextjs
#
#EXPOSE 3000
#
#ENV PORT 3000
#
## Next.js collects completely anonymous telemetry data about general usage.
## Learn more here: https://nextjs.org/telemetry
## Uncomment the following line in case you want to disable telemetry.
## ENV NEXT_TELEMETRY_DISABLED 1
#
#CMD ["node", "server.js"]


## Go sections
###
### Build
###
#FROM golang:1.16-buster AS go-build
#
#WORKDIR /app
#
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#
#ARG DIST_PATH
#RUN test -n "$DIST_PATH" || (echo "DIST_PATH  not set" && false)
#COPY ./$DIST_PATH ./
#
#RUN #go build -o /go-app
#
###
### Deploy
###
#FROM gcr.io/distroless/base-debian10
#
#WORKDIR /
#
#COPY --from=build /go-app /go-app
#
#EXPOSE 8080
#
#USER nonroot:nonroot
#
#ENTRYPOINT ["/go-app"]
# Install dependencies only when needed
FROM node:16 AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
#RUN #apk add --no-cache libc6-compat
WORKDIR /app
COPY .yarn/releases ./.yarn/releases
COPY package.json ./
COPY yarn.lock ./
COPY .yarnrc.yml ./
RUN yarn
#COPY . .
#RUN yarn test
#RUN yarn lint
#RUN yarn build

# Rebuild the source code only when needed
#FROM node:16-alpine AS builder
#WORKDIR /app
#COPY --from=deps /app/yarn.lock ./
#COPY --from=deps /app/node_modules ./node_modules
#COPY --from=deps /app/package.json ./
#COPY --from=deps /app/.yarn/cache ./.yarn/cache
##
#COPY nx.json ./
#COPY workspace.json ./
#COPY vite.config.js ./
#COPY tsconfig.base.json ./
#COPY Makefile ./
#COPY jest.*.js ./
#COPY go.mod ./
#COPY go.sum ./
#COPY Dockerfile ./
#COPY Dockerfile ./
#COPY babel.config.json ./
##COPY .eslintignore ./
#COPY .eslintrc.json ./
#COPY .yarnrc.yml ./
#COPY .yarn/releases ./.yarn/releases
### docker
#COPY skaffold.yaml ./
#COPY docker-compose.yml ./
##
#COPY apps ./apps
#COPY libs ./libs
#COPY scripts ./scripts
##
##RUN #npx nx run-many --target=build --parallel --max-parallel=4 --prod --all
#RUN #yarn
#RUN #yarn test
#RUN #yarn build
##RUN npx nx run-many --parallel 10 --all --target=build
#
#
#
#FROM node:16-alpine AS runner
#WORKDIR /app
##
#ENV NODE_ENV production
#COPY --from=builder /app/apps ./dist

