# h3load

h3load is a lightweight HTTP/3 load testing CLI written in Go, built on top of [quic-go](https://github.com/quic-go/quic-go) and its HTTP/3 implementation.

The project explores HTTP/3, QUIC, concurrency and load generation through a simple tool that is developed incrementally from a minimal CLI into a configurable load tester.

## Overview

h3load keeps three load-testing concepts separate:

- **RPS** controls how frequently new requests are started.
- **Concurrency** limits the number of requests in flight at the same time.
- **Connections** control the number of independent QUIC connections.

HTTP/3 multiplexes multiple request streams over a QUIC connection, so increasing concurrency does not automatically require additional connections.

## Load generation

The initial implementation uses a concurrency-limited model. Controlled request rates can later be introduced independently from concurrency.

Achieved RPS is calculated from completed requests over elapsed time, so it can differ from the configured target rate when response times increase or the target system becomes overloaded.

## HTTP/3 and metrics

h3load uses quic-go for HTTP/3 communication and keeps transport handling separate from load-generation logic.

Each request can produce information such as status code, response size, latency and error type. The project distinguishes between header latency, TTFB and complete request latency, with more detailed measurements introduced progressively.

## Development

The CLI, configuration, load generation, HTTP/3 client and metrics are kept as separate responsibilities. The codebase is intentionally kept small, with new abstractions introduced only when they become necessary.

h3load is intended for systems you own or are explicitly authorized to test. During development, a local HTTP/3 server is recommended.