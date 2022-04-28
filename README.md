### Example app

Example app to export traces to an HTTP server using the http exporter

Just run
```
go get
go run main.go
```
to export traces to an http server running on `127.0.0.1:4000`

You can use FluentBits HTTP input plugin to start the server in case you are interested in getting the traces to FluentBit using the following config file:

```
[INPUT]
	name http
	host 0.0.0.0
	port 4000

[OUTPUT]
	name stdout
	match *
```

and run `fluent-bit --config http.conf`
