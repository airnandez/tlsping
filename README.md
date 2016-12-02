# tlsping — a tool for measuring TLS handshake latency

## Overview
`tlsping` is a command line tool to measure the time required to perform [Transport Layer Security (TLS)](https://en.wikipedia.org/wiki/Transport_Layer_Security) connections with a network server.

It concurrently establishes several network connections, performs the TLS handshake on every one of them, measures the time spent and reports a summary on the observed results.

## How to use
Example of usage:

```bash
$ tlsping mail.google.com:443
tlsping: TLS connection to server mail.google.com:443 (10 connections)
tlsping: min/avg/max/stddev = 92.90ms/94.65ms/96.77ms/913.94µs
```

This is the synopsis of the command:

```
tlsping [-c count] [-tcponly] [-json] [-ca=<file>] [-insecure] <server address>
tlsping -help
tlsping -version
```

For getting details on options do `tlsping --help`.

## Installation
Download a **binary release** for your target operating system from the [releases page](https://github.com/airnandez/tlsping/releases).

Alternatively, if you prefer to **build from sources**, you need the [Go programming environment](https://golang.org). Do:

```
go get -u github.com/airnandez/tlsping/...
```

## Credits

This tool was developed and is maintained by Fabio Hernandez at [IN2P3 / CNRS computing center](http://cc.in2p3.fr) (Lyon, France).

## License
Copyright 2016 Fabio Hernandez

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
