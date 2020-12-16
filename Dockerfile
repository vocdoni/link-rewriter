## BUILDER

FROM golang:1.15-alpine AS builder

WORKDIR /app
ADD . /app

RUN go build -o rewriter

## RUNNER

FROM alpine:3.12

WORKDIR /app
COPY --from=builder /app/rewriter .
ADD config.yaml .

CMD [ "/app/rewriter" ]
