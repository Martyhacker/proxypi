# proxypi

Proxypi is the proxysite and proxy api

## Installation

Written in go version 1.18

```bash
go run server.go
```

## Server Health
You can know is server running and healthy
```
http://localhost:3031 GET
```
## Proxy API
Send POST request to link below with "uri" field
```
http://localhost:3031/uri POST
```
## Proxy Site
Send your url like shown below as query parameter
```
http://localhost:3031/site/?q=http://stackoverflow.com GET
```

## Contributing
This project in development now.

## Developer
umyt.dev@gmail.com