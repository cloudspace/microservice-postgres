FROM scratch
ADD ./pgquery /pgquery
ENTRYPOINT ["/pgquery"]