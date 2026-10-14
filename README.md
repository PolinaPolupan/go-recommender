# Go Recommendation Service

A small recommendation service written in Go.

The service accepts a user ID, ranks candidate items using the user's
interaction history, and returns the top recommendations over HTTP.

## API

### Get recommendations

```text
GET /recommendations?user_id=42&limit=2

## Concurrency experiment

Three approaches were compared for scoring 10,000 items:

| Approach | Time | Allocations |
|---|---:|---:|
| Sequential | ~170 µs | ~1 |
| 10,000 goroutines | ~3.33 ms | ~20,000 |
| 12-worker pool | ~2.78 ms | ~15 |

For this workload, sequential execution was significantly faster because
the scoring operation itself is very cheap.

The concurrent implementations were kept as experiments to study
goroutines, channels, synchronization, and worker pools rather than
being used in the request path.
