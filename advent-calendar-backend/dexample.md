FROM golang

ADD ./ /go/src/app

WORKDIR /go/src/app

RUN go get ./

RUN 

RUN go build -o ./app

CMD ./app