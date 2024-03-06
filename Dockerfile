FROM --platform=linux/amd64 golang:1.20.0

ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/local/bin/aws-lambda-rie
RUN chmod 755 /usr/local/bin/aws-lambda-rie

COPY . /workspace

WORKDIR /workspace
RUN chmod +x ./entry_script.sh

RUN go build -o lambda/bin/main lambda/main.go
RUN chmod +x lambda/bin/main

ENTRYPOINT [ "./entry_script.sh"," lambda/bin/main" ]