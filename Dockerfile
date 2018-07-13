FROM golang:1.9
RUN apt-get update
RUN apt-get install -y proj-bin libproj-dev
RUN mkdir -p /go/src/Go-GoSAFE.converter
WORKDIR /go/src/Go-GoSAFE.converter
COPY . /go/src/Go-GoSAFE.converter/
RUN go get ./
RUN go build -o converter .
CMD ["/go/src/Go-GoSAFE.converter/converter"]
