FROM alpine:latest

RUN mkdir /socle

WORKDIR /socle

COPY bin/app .
COPY socle.yaml . 
COPY tls.yaml .  
COPY *.crt .
COPY *.key .
COPY *.p12 .


CMD [ "/socle/app"]