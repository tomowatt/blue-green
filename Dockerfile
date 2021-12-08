FROM golang:1.17-buster as builder

WORKDIR /build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.release=`git rev-parse --short=7 HEAD`'" -o blue-green .

FROM scratch
WORKDIR /app
COPY --from=builder /build/blue-green blue-green
COPY ./template/ ./template

CMD [ "/app/blue-green" ]
