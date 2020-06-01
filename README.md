# tlsping — a tool for measuring TLS handshake latency

## Overview
`tlsping` is a command line tool to measure the time required to perform [Transport Layer Security (TLS)](https://en.wikipedia.org/wiki/Transport_Layer_Security) connections with a network server.

It concurrently establishes several network connections, performs the TLS handshake on every one of them, measures the time spent handshaking and reports a summary on the observed results.

## How to use

#### Examples of usage:

* Measure the time to establish a TCP connection to host `mail.google.com` port `443` and perform the TLS handshaking:
 
	```bash
	# The hostname 'mail.google.com' resolves to an IPv4 address
	$ tlsping mail.google.com:443      
	tlsping: TLS connection to mail.google.com:443 (216.58.204.133) (10 connections)
	tlsping: min/avg/max/stddev = 95.95ms/96.31ms/96.63ms/218.19µs
	```

* Same measurement as above but connect to host `www.cloudflare.com` when it resolves to to IPv6 address `2606:4700::6811:d209`:

	```bash
	# The hostname 'www.cloudflare.com' resolves to an IPv6 address
	$ tlsping www.cloudflare.com:443 
	tlsping: TLS connection to www.cloudflare.com:443 (2606:4700::6811:d209) (10 connections)
	tlsping: min/avg/max/stddev = 85.36ms/86.63ms/88.98ms/1.14ms
	```
	
*	Measure only the time to establish the TCP connection (i.e. do not perform TLS handshaking) to remote server at IPv6 address `2a00:1450:400a:800::2005` port `443`:

	```bash
	# To specify an IPv6 address and port enclose the IP address in '[' and ']'
	$ tlsping -tcponly [2a00:1450:400a:800::2005]:443
	tlsping: TCP connection to [2a00:1450:400a:800::2005]:443 (2a00:1450:400a:800::2005) (10 connections)
	tlsping: min/avg/max/stddev = 5.85ms/5.97ms/6.08ms/61.55µs
	```

#### Synopsis

```bash
tlsping [-c count] [-tcponly] [-json] [-ca=<file>] [-insecure] <server address>
tlsping -help
tlsping -version
```

#### Getting help

```bahs
tlsping -help
```

## Installation
Download a **binary release** for your target operating system from the [releases page](https://github.com/airnandez/tlsping/releases).

Alternatively, if you prefer to **build from sources**, you need the [Go programming environment](https://golang.org). Do:

```
go get -u github.com/airnandez/tlsping/...
```

## Credits

This tool was developed and is maintained by Fabio Hernandez at [IN2P3 / CNRS computing center](http://cc.in2p3.fr) (Lyon, France).

## License
Copyright 2016-2020 Fabio Hernandez

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
