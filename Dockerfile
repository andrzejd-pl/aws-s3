FROM andrzejd/go-env

RUN go get -u github.com/andrzejd-pl/aws-s3
RUN go install github.com/andrzejd-pl/aws-s3

CMD ["aws-s3"]
