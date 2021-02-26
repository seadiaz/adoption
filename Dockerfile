FROM alpine:3.8
RUN adduser -D -u 1000 adoption

FROM scratch
COPY /adoption /adoption
COPY --from=0 /etc/passwd /etc/passwd
USER 1000
ENTRYPOINT [ "/adoption" ]
CMD ["server"]