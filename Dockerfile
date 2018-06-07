FROM scratch
# ADD ca-certificates.crt /etc/ssl/certs/
ADD main /
CMD ["/main"]
