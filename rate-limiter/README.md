# Rate limiter 

# Token Bucket Algorithm

The Token Bucket throttling algorithm works as follows: 
A token bucket is a container that has pre-defined capacity. 
Tokens are put in the bucket at preset rates periodically. 
Once the bucket is full, no more tokens are added.

> [Distributed API Rate Limiter](https://systemsdesign.cloud/SystemDesign/RateLimiter)


# Leaky Buckey Algorithm

Similar to the token bucket, leaky bucket also has a bucket with a finite
capacity for each client. However, instead of tokens, it is filled with 
requests from that client. Requests are taken out of the bucket and 
processed at a constant rate. If the rate at which requests arrive is 
greater than the rate at which requests are processed, the bucket will 
fill up and further requests will be dropped until there is space in the 
bucket.

# Fixed Window Counter Algorithm


Fixed window counter algorithm divides the timeline into fixed-size 
windows and assign a counter to each window. Each request, based on its
arriving time, is mapped to a window. If the counter in the window has
reached the limit, requests falling in this window should be rejected.
For example, if we set the window size to 1 minute. Then the windows are
[00:00, 00:01), [00:01, 00:02), ...[23:59, 00:00).

# Sliding Window Logs Algorithm

As discussed above, the fixed window counter algorithm has a major 
drawback: it allows more requests to go through at the edge of a window.
The sliding window logs algorithms fixes this issue. It works as follows:

- The algorithm keeps track of request timestamps. Timestamp data is usually kept in cache,
  such as [sorted sets of Redis](https://engineering.classdojo.com/blog/2015/02/06/rolling-rate-limiter/) .
- When a new request comes in, remove all the outdated timestamps. 
  Outdated timestamps are defined as those older than the start of the 
  current time window.
  Storing the timestamps in sorted order in sorted set enables us to effieciently find the outdated timestamps.
- Add timestamp of the new request to the log.
- If the log size is the same or lower than the allowed count, a request is accepted. Otherwise, it is rejected.


# Sliding Window Counter algorithm

The Sliding Window Counter algorithm is a hybrid approach that combines the Fixed Window Counter algorithm and
Sliding Window Logs algorithm.

We can explain the above calculation in a more lucid way: Assume the rate
limiter allows a maximum of 100 requests per hour, and there are 84 
requests in the previous hour and 36 requests in the current hour. 
For a new request that arrives at a 25% position in the current hour, 
the number of requests in the rolling window is calculated using the 
following formula:


Requests in current window + (Requests in the previous window * 
overlap percentage of the rolling window and previous window)

Using the above formula, we get (36 + (84 * 75%)) = 99


