FROM alpine

ADD /adoption /adoption

ENTRYPOINT [ "/adoption" ]