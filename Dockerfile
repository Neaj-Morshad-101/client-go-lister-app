FROM ubuntu

COPY ./client-go-lister-app ./client-go-lister-app

ENTRYPOINT ["./client-go-lister-app" ]