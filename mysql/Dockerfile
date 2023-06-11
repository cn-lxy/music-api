FROM mysql:debian

LABEL maintainer="lxy <longxinyu2021@163.com>"
COPY ./music.sql  /tmp/music.sql
RUN mv /tmp/music.sql /docker-entrypoint-initdb.d
RUN rm -rf /tmp/music.sql