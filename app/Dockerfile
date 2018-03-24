FROM golang:latest 

RUN go get github.com/gin-gonic/gin
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/go-sql-driver/mysql

ADD ./src/api /go/src/api 
RUN go install api 
CMD ["api"]