# Redis Challenge in Go

This project is a Go implementation of the Redis Challenge as described on [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-redis). It aims to replicate the basic functionalities of Redis, including key-value storage, handling concurrent connections, and supporting a subset of the Redis protocol.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.15 or later recommended)
- Redis (optional, for comparison and benchmarking)

### Installing

To get a development environment running, clone the repository and build the project:

```
bash
git clone https://github.com/yourusername/redis-challenge-go.git
cd redis-challenge-go
go build ./cmd/server
```

This will compile the server application. You can start the server with:

```
bash
./server
```

By default, the server listens on port 6379, but you can change this in the configuration if necessary.

## Running the tests

To run the automated tests for this system, use:

```
bash
go test ./...
```


This command will recursively run all tests in the project.

## Usage

Once the server is running, you can connect to it using any Redis client by pointing the client to `localhost:6379`. Basic commands supported include `PING`, `SET`, `GET`, `DEL`, `INCR`, `DECR`, `LPUSH`, `RPUSH`, and `SAVE`.

Example using `redis-cli`:

```
bash
redis-cli -p 6379
127.0.0.1:6379> PING
PONG
```


## Benchmarking

You can benchmark the performance of this Redis server implementation using `redis-benchmark`. For example:

```
bash
redis-benchmark -t set,get -n 100000 -q
```

Compare the output with the performance of the official Redis server to see how this implementation stacks up.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- Thanks to [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-redis) for providing the challenge.
- Salvatore Sanfilippo for creating Redis.