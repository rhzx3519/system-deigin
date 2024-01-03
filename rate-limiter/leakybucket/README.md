# Leaky bucket


In this algorithm, we use limited sized queue and process requests at a constant rate from queue in First-In-First-Out(FIFO) manner.

For each request a user makes,

- Check if the queue limit is exceeded.
- If the queue limit has exceeded, the request is dropped.
- Otherwise, we add requests to queue end and handle the incoming request.


Requests are processed at a constant rate from the queue in a FIFO manner(removed from the start of the queue and handled) from a background process. (you can use LPOP command in Redis at a constant rate, for example for 60 req/min, you can remove 1 element per second and handle the removed request)


Pros:
- Memory efficient given the limited queue size.
- Requests are processed at a fized rate, therefore it is suitable for use cases where a stable outflow rate id required.

Cons:
- A burst of traffic fills up the queue with old requests, and if they are not processed in time, recent requests will be rate limited.
