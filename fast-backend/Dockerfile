FROM golang:1.9.2

ADD ./ /go/src/github.com/rccpccl/moback/fast-backend

RUN cd /go/src/github.com/rccpccl/moback/fast-backend &&\
 go install

EXPOSE 8080

CMD ["/go/bin/fast-backend"]
