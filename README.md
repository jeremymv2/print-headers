# print-headers

Small go app that simply prints all HTTP headers from the request.
Useful for introspection of the HTTP request and headers.

The binary is published as a docker image: [jmv2/print-headers](https://hub.docker.com/repository/docker/jmv2/print-headers)

Usage:

```
$ docker run -p 8080:8080 jmv2/print-headers:0.0.3
```

Example response:

```
Hello, here is your request:
GET / HTTP/1.1
Host: demo-ingress.ethzero.cloud
Accept: */*
User-Agent: curl/7.82.0
X-Forwarded-For: 10.7.0.157
X-Forwarded-Host: demo-ingress.ethzero.cloud
X-Forwarded-Port: 80
X-Forwarded-Proto: http
X-Forwarded-Scheme: http
X-Real-Ip: 10.7.0.157
X-Request-Id: a7e82e6eaa704246a6068e70d1c930b4
X-Scheme: http
```
