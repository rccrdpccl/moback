FROM golang:1.9.2

ADD ./ /go/src/github.com/rccrdpccl/moback/slow-backend

RUN cd /go/src/github.com/rccrdpccl/moback/slow-backend &&\
 go install

EXPOSE 8080

CMD ["/go/bin/slow-backend"]
