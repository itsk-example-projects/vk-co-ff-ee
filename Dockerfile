FROM golang:1.25-alpine AS builder

WORKDIR /build

RUN apk add --no-cache upx

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -v -ldflags="-s -w" -o vk-co-ff-ee cmd/main.go

RUN upx --best --lzma vk-co-ff-ee

FROM scratch

COPY --from=builder /build/vk-co-ff-ee ./
COPY --from=builder /build/config.yaml ./

EXPOSE 8080

CMD ["./vk-co-ff-ee"]
