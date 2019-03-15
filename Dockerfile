FROM alpine:3.8
# ADD ca-certificates.crt /etc/ssl/certs/
ADD main /
CMD ["/main"]
