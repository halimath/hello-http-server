FROM scratch

USER 1001
EXPOSE 8080

ADD hello-http-server /

CMD ["/hello-http-server"]
