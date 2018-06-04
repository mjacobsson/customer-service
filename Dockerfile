FROM golang:1.10.2-stretch

ENV GOPATH=/home/go
ENV PATH=${PATH}:${GOPATH}/bin

RUN mkdir -p ${GOPATH}/src
RUN mkdir -p ${GOPATH}/bin

ADD . ${GOPATH}/src/github.com/mjacobsson/customer-service
WORKDIR ${GOPATH}/src/github.com/mjacobsson/customer-service

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN go build -o customer-service main.go
CMD ["./customer-service"]

