FROM golang:1.19-buster
WORKDIR /app
RUN apt -y update
RUN apt -y upgrade
RUN apt install -y python3-pip
RUN apt install -y python3-dev

ENV PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig
ENV DIRECTORY="/app"
ENV PORT=8080

COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN pip3 install -r internal/application/core/api/lib/requirements.txt
RUN GOOS="linux" go test ./...
RUN GOOS="linux" go build cmd/main/main.go

CMD [ "/app/main"]