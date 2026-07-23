# URL Shortener Learning Task List

1. Start a minimal HTTP server with `net/http`
2. Move the health handler out of `main.go`
3. Add method checking to `/health`
4. Set explicit response headers and status code in `/health`
5. Handle `ListenAndServe()` errors properly
6. Run `gofmt` and learn standard Go formatting
7. Add a root endpoint to explain what the service is
8. Introduce a dedicated server setup function to keep `main.go` small
9. Add basic configuration for the server port
10. Read configuration from environment variables
11. Validate configuration at startup
12. Add graceful shutdown support
13. Understand request lifecycle and `context.Context`
14. Add a `POST /shorten` endpoint skeleton
15. Learn JSON decoding with `encoding/json`
16. Define request and response structs for shortening
17. Add input validation for the original URL
18. Learn URL parsing with `net/url`
19. Define the core URL model in `internal/model`
20. Move shortening logic into `internal/service`
21. Design a service interface and discuss why interfaces matter
22. Add an in-memory repository implementation
23. Define a repository interface in `internal/repository`
24. Wire handler -> service -> repository dependencies

................................................

25. Generate short codes with a simple deterministic approach
26. Discuss collision risk and improve short code generation
27. Add a `GET /{code}` redirect flow using standard library routing constraints
28. Learn why dynamic path routing gets awkward in `net/http`
29. Decide when a third-party router becomes justified
30. Return correct redirect status codes and `Location` header
31. Add unit tests for the health handler
32. Learn table-driven tests in Go
33. Add unit tests for the shorten handler
34. Add unit tests for the service layer
35. Add unit tests for the repository layer
36. Learn how to use `httptest`
37. Improve error responses to be consistent
38. Introduce a small response-writing helper
39. Add structured logging with only standard library basics first
40. Add request logging middleware
41. Introduce middleware chaining
42. Add panic recovery middleware
43. Add request IDs
44. Add timeout handling at the server level
45. Discuss server timeouts: read, write, idle, header
46. Add basic health vs readiness endpoint distinction
47. Add persistent storage design discussion
48. Introduce PostgreSQL concepts before implementation
49. Design the first database schema for shortened URLs
50. Discuss primary keys, unique constraints, and indexes
51. Add SQL migrations strategy
52. Add a PostgreSQL repository implementation
53. Compare in-memory vs PostgreSQL repository behavior
54. Handle duplicate short code conflicts safely
55. Add created-at and updated-at fields
56. Add expiration support for links
57. Discuss TTL and expiry semantics
58. Add analytics event model design
59. Track redirect counts in memory first
60. Discuss race conditions and concurrent access
61. Protect in-memory state with `sync.RWMutex`
62. Learn when mutexes are needed in Go
63. Add background cleanup for expired links
64. Discuss goroutines and lifecycle ownership
65. Add graceful background worker shutdown
66. Add integration tests for HTTP endpoints
67. Add integration tests for repository behavior
68. Introduce Docker for local development
69. Add Docker Compose for app + Postgres
70. Introduce Redis use cases before adding Redis
71. Add a cache design discussion for redirect lookups
72. Add Redis caching for hot short codes
73. Discuss cache-aside and cache invalidation
74. Add rate limiting design discussion
75. Implement simple in-memory rate limiting
76. Compare fixed window vs token bucket
77. Add authentication discussion for admin-only endpoints
78. Add an admin endpoint to inspect stored URLs
79. Add pagination basics
80. Add delete or disable short link capability
81. Discuss soft delete vs hard delete
82. Add observability discussion: logs, metrics, traces
83. Add basic metrics endpoint or counters
84. Discuss production deployment concerns
85. Add GitHub Actions for formatting and tests
86. Add linting after understanding what it catches
87. Add API documentation
88. Write a design doc for the project architecture
89. Prepare interview talking points for each subsystem
90. Review tradeoffs and decide what makes the final project resume-worthy
