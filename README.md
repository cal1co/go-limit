# go-limit

This package provides a few rate limiting algorithms that can be used to control the rate of incoming requests or events in your applications. Rate limiting is a crucial technique to prevent resource exhaustion, protect your services from abuse, and ensure fair usage of resources.

## Algorithms Included

This package includes the following rate limiting algorithms:

1. **Token Bucket Rate Limiter (`tokenbucket`)**:
   - The Token Bucket algorithm allows requests to be processed at a fixed rate, making it useful for controlling the rate of incoming requests over time.

2. **Leaky Bucket Rate Limiter (`leakybucket`)**:
   - The Leaky Bucket algorithm allows requests to be processed at a constant rate, useful for smoothing out the rate of requests and preventing bursts of traffic.

3. **Sliding Window Rate Limiter (`slidingwindow`)**:
   - The Sliding Window algorithm allows a fixed number of requests within a specific time window, ideal for maintaining a consistent request rate over time.

4. **Fixed Window Rate Limiter (`fixedwindow`)**:
   - The Fixed Window algorithm allows a fixed number of requests within fixed time windows, useful when you need to enforce a strict limit within specific intervals.

## Usage

Each rate limiting algorithm is provided in its own subpackage within this package. You can use these rate limiters in your Go applications to control the rate of incoming requests or events.

Refer to the documentation and test files within each subpackage for detailed usage examples and instructions on how to use each algorithm.

## Installation

You can include this package in your Go project using `go get`:

```sh
go get github.com/yourusername/rate_limiter
