FROM alpine:3.15.4

COPY new.html /
COPY view.html /
COPY guestbook /

WORKDIR /
CMD /guestbook
