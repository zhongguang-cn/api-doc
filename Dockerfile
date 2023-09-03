FROM golang:alpine

WORKDIR /work/

COPY . .

ENV LANG C.UTF-8
ENV TZ=Asia/Shanghai

EXPOSE 8081

CMD ["go", "run", "."]