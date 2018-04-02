FROM golang:latest 

RUN go get github.com/gin-gonic/gin

ADD ./tests /go/src/tests 
RUN go install 
CMD ["go test /go/src/tests"]