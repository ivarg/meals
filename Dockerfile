FROM busybox:buildroot-2014.02
MAINTAINER Ivar Gaitan
COPY . /go/src/meals
WORKDIR /go/src/meals
EXPOSE 8080
ENTRYPOINT ["./meals_linux_amd64"]
