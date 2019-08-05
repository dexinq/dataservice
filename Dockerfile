FROM alpine
ADD dataservice-srv /dataservice-srv
ENTRYPOINT [ "/dataservice-srv" ]
