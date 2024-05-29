FROM ubuntu:latest

# Essential for using tls
RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

# grpc port
EXPOSE 50051

ADD build/userService /app/userService
RUN ls -l

CMD /app/userService
