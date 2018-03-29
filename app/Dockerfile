FROM golang:latest 

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get google.golang.org/api/drive/v3
RUN go get golang.org/x/oauth2/...
RUN go get github.com/go-sql-driver/mysql

ADD ./src/api /go/src/api 
RUN go install api
CMD ["api"]