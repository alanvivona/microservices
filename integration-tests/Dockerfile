FROM golang:latest 

ADD ./tests /go/src/tests 
RUN go install 
CMD ["go test /go/src/tests"]