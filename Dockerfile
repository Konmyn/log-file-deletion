FROM harbor.geniusafc.com/docker.io/alpine:3.9.2
ENV TZ Asia/Shanghai
RUN apk add --no-cache tzdata
WORKDIR /app
ADD log-walk log-walk
CMD /app/log-walk
