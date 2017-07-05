FROM golang:1.8
COPY . /go/src/github.com/eleanorrigby/openservicebroker_exper
WORKDIR /go/src/github.com/eleanorrigby/openservicebroker_exper
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /openservicebroker_exper .

FROM ubuntu:16.04
COPY --from=0 /openservicebroker_exper /openservicebroker_exper
ADD charts/jenkins /jenkins
CMD ["/openservicebroker_exper", "-logtostderr"]