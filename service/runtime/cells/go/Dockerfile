FROM golang:1.18

WORKDIR /
COPY entrypoint.sh /
RUN chmod 755 entrypoint.sh
ENTRYPOINT ["bash", "/entrypoint.sh"]
