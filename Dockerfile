FROM golang:1.14.2-alpine AS build
ARG dir=/todo
ADD . ${dir}
RUN apk update && \
    apk add --virtual build-dependencies build-base git && \
    cd ${dir} && \
    make install && \
    make builds

# final stage
FROM alpine:3.7
ARG dir=/todo
WORKDIR /app

COPY --from=build ${dir}/api /app/