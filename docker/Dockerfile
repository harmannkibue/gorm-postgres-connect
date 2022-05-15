FROM golang:1.16 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

COPY ../go.mod ./
COPY ../go.sum ./

RUN go mod download

COPY ./* ./

CMD ["air"]
