FROM scratch
LABEL maintainer="Chris Collins <collins.christopher@gmail.com>"

COPY pkg/* /
CMD [ "/gorestapp" ]
