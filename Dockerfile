FROM golang:1.18.3
WORKDIR /app
COPY . .
EXPOSE 8080
CMD ./gorun.sh
