# go-qulia

![Coverage](https://img.shields.io/badge/Coverage-97.8%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/qulia/go-qulia.svg)](https://pkg.go.dev/github.com/qulia/go-qulia)

Go data structures, and helper libraries.

# Data Structures

- [Graph](lib/graph)
- [Heap](lib/heap/)
- [Queue](lib/queue/)
- [Set](lib/set)
- [Skiplist](lib/skiplist/)
- [Stack](lib/stack/)
- [Tree](lib/tree/)
  - [BinaryIndexTree](lib/tree/bit.go)
  - [SegmentTree](lib/tree/segment.go)
  - [BinarySearchTree](lib/tree/bst.go)
- [Trie](lib/trie)
- [UnionFind](lib/unionfind/)

# Algo

- RateLimiter
  - [TokenBucket](algo/ratelimiter/tokenbucket/)
  - [LeakyBucket](algo/ratelimiter/leakybucket/)
  - [FixedWindowCounter](algo/ratelimiter/fixedwindowcounter/)
  - [SlidingWindowLog](algo/ratelimiter/slidingwindowlog/)
  - [SlidingWindowCounter](algo/ratelimiter/slidingwindowcounter/)

# Clone
- [Clone](clone/clone.go)

# Middleware

- [RateLimiter](http/server/middleware/ratelimiter)
  
# Concurrency

- [Unique](concurrency/unique/)

# Data Processing

- Windowing
  - [FixedWindow](dataprocessing/window/window.go)
  - [SlidingWindow](dataprocessing/window/window.go)

# Messaging

- [Pub/Sub Broker](messaging/broker/)

# Mock

- [MockTimeProvider](mock/mock_time/provider.go)

---

---

Contributions are welcome!
