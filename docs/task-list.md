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
25. Generate short codes with a simple deterministic approach
26. Discuss collision risk and improve short code generation
27. Add a `GET /{code}` redirect flow using standard library routing constraints
28. Learn why dynamic path routing gets awkward in `net/http`
29. Decide when a third-party router becomes justified
30. Introduce Chi and replace fallback `/` routing with explicit dynamic routes
31. Return correct redirect status codes and `Location` header
32. Add unit tests for the health handler
33. Learn table-driven tests in Go
34. Add unit tests for the shorten handler
35. Add unit tests for the service layer
36. Add unit tests for the repository layer
37. Learn how to use `httptest`
38. Introduce `testify` after standard testing and `httptest`
39. Improve error responses to be consistent
40. Introduce a small response-writing helper
41. Add structured logging with only standard library basics first
42. Add request logging middleware

................................................

43. Introduce middleware chaining
44. Add panic recovery middleware
45. Add request IDs
46. Add timeout handling at the server level
47. Discuss server timeouts: read, write, idle, header
48. Add basic health vs readiness endpoint distinction
49. Add persistent storage design discussion
50. Introduce PostgreSQL concepts before implementation
51. Design the first database schema for shortened URLs
52. Discuss primary keys, unique constraints, and indexes
53. Add SQL migrations strategy
54. Add a PostgreSQL repository implementation
55. Compare in-memory vs PostgreSQL repository behavior
56. Handle duplicate short code conflicts safely
57. Add created-at and updated-at fields
58. Add expiration support for links
59. Discuss TTL and expiry semantics
60. Add analytics event model design
61. Track redirect counts in memory first
62. Discuss race conditions and concurrent access
63. Protect in-memory state with `sync.RWMutex`
64. Learn when mutexes are needed in Go
65. Add background cleanup for expired links
66. Discuss goroutines and lifecycle ownership
67. Add graceful background worker shutdown
68. Add integration tests for HTTP endpoints
69. Add integration tests for repository behavior
70. Introduce Docker for local development
71. Add Docker Compose for app + Postgres
72. Introduce Redis use cases before adding Redis
73. Add a cache design discussion for redirect lookups
74. Add Redis caching for hot short codes
75. Discuss cache-aside and cache invalidation
76. Add rate limiting design discussion
77. Implement simple in-memory rate limiting
78. Compare fixed window vs token bucket
79. Add authentication discussion for admin-only endpoints
80. Add an admin endpoint to inspect stored URLs
81. Add pagination basics
82. Add delete or disable short link capability
83. Discuss soft delete vs hard delete
84. Add observability discussion: logs, metrics, traces
85. Add basic metrics endpoint or counters
86. Discuss production deployment concerns
87. Add GitHub Actions for formatting and tests
88. Add linting after understanding what it catches
89. Add API documentation
90. Write a design doc for the project architecture
91. Prepare interview talking points for each subsystem
92. Review tradeoffs and decide what makes the final project resume-worthy
