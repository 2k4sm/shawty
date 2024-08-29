# Build Stage
FROM golang:alpine as build

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .

# Run Stage
FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . ./app

COPY --from=build /build/main /app/

WORKDIR /app

EXPOSE 6969

CMD [ "./main" ]

