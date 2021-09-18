FROM golang:alpine3.14 as build
WORKDIR /src
COPY . .
RUN  go build -v  -o /bin/registry-ctl . 

FROM registry:2.7
COPY --from=build /bin/registry-ctl /bin/registry-ctl
COPY ./start.sh /bin/start.sh
#RUN rm -rf /src && mkdir -p /etc/registry-ctl
EXPOSE 5000
ENTRYPOINT ["sh","/bin/start.sh"]
                
