FROM golang:1.27-alpine AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -trimpath -ldflags="-s -w" -o /out/server .

FROM alpine:3.22
RUN addgroup -S app && adduser -S -G app app

WORKDIR /app

COPY --from=builder /out/server ./server

USER app

ENV PORT=8080
ENV APP_MODE=production

EXPOSE 8080

ENTRYPOINT ["/app/server"]