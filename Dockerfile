FROM golang:1.20.3-alpine3.17 AS build
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk add --update make
RUN go mod download
RUN make

FROM alpine:3.14
EXPOSE 8881
COPY --from=build /app/bin/dataload /dataload
COPY .env .env
CMD [ "/dataload"]
