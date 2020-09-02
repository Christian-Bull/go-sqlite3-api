# get golang image
FROM arm32v7/golang:1.15-alpine

# install git - package dependencies
RUN apk update && apk add git

# create app dir
RUN mkdir /app

# add everything to that dir
ADD cmd /app

# change working directory
WORKDIR /app

# go get
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/mattn/go-sqlite3

# build app
RUN go build -o main ./go-api/.

# Expose port
EXPOSE 2000

# run it yo
CMD ["/app/main"]