FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY notely /usr/local/bin/
ENV PORT=8080

CMD ["notely"]
