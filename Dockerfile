FROM scratch

USER 1001

ADD hello-http-server /

CMD ["/hello-http-server"]