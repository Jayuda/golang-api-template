FROM golang:1.17
RUN mkdir /app 
ADD . /app/
WORKDIR /app

CMD ["/app/binary"]