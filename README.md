## Loc2Country

Location coordinates (lat/lon) to ISO alpha-3 country code. Responds in microseconds.

## Manual

Input format: latitude, longitude

Output format: 3-letter-ISO-country-code, time-taken-to-respond-in-nanos

## HowTo

1. Run start.sh
2. This will start a TCP server (localhost:3333) by default.
3. Connect to the server by using telnet. (eg: "telnet localhost 3333")
4. Input lat and lon seperated by comma, returns 3 letter country code and time taken to respond in nanoseconds.

## Compiling

To compile, run:

``` bash
go build src/server.go
```

To compile for a Linux machine from Mac, run (with correct architecture):

``` bash
env GOOS=linux GOARCH=amd64 go build src/server.go
```

## Testing

To test, run:

```
go test
```


## Example

Starting the server:

``` bash
$ sh start.sh 
2016/08/18 23:30:07 Creating server with address localhost:3333
2016/08/18 23:30:07 Loading data..
2016/08/18 23:30:13 Loading complete.
2016/08/18 23:30:13 Total Entries: 5235316
2016/08/18 23:30:13 Boot time: 5 seconds
```

``` bash
$ telnet 127.0.0.1 3333
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
12,77
IND,17176
```

## Data

The world boundaries were generated using QGIS, converted to a set of ~350 million geohashes at precision level 6 and then reduced (compressed) to a set of ~5 million geohashes using [georaptor](https://github.com/ashwin711/georaptor). 


## Contributors

Sooraj B - [@soorajb](http://github.com/soorajb)

Ashwin Nair - [@ashwin711](http://github.com/ashwin711)

Harikrishnan Shaji - [@har777](http://github.com/har777)

## License

MIT License
