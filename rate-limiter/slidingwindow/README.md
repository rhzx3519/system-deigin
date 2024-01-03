# Sliding Window Logs

As discussed above, the fixed window counter algorithm has a major drawback: it allows more requests to go through at the edge of a window. The sliding window logs algorithms fixes this issue. It works as follows:
- The algorithm keeps track of request timestamps. Timestamp data is usually kept in cache, such as sorted sets of Redis .
- When a new request comes in, remove all the outdated timestamps. Outdated timestamps are defined as those older than the start of the current time window.
Storing the timestamps in sorted order in sorted set enables us to effieciently find the outdated timestamps.
- Add timestamp of the new request to the log.
- If the log size is the same or lower than the allowed count, a request is accepted. Otherwise, it is rejected.

Pros:
- Sliding Window Logs algorithm works flawlessly. Rate limiting implemented by this algorithm is very accurate. In rolling window, requests will not exceed the rate limit.

Cons:
- Sliding Window Logs algorithm consumes a lot of memorybecause even if a request is rejected , it's timestamp will still be stored in memory.


# Sliding Window Counter

The Sliding Window Counter algorithm is a hybrid approach that combines the Fixed Window Counter algorithm and Sliding Window Logs algorithm.

Assume the rate limiter allows a maximum of 100 requests per hour, and there are 84 requests in the previous hour and 36 requests in the current hour. For a new request that arrives at a 25% position in the current hour, the number of requests in the rolling window is calculated using the following formula:


Requests in current window + (Requests in the previous window * overlap percentage of the rolling window and previous window)

Using the above formula, we get (36 + (84 * 75%)) = 99

Since the rate limiter allows 100 requests per hour, the current request will go through.

This algorithm assumes a constant request rate in the (any) previous window, which is not true as there can be request spikes too during a minute and no request during another hour. Hence the result is only an approximated value.


Pros:
- Memory efficient.
- It smoothes out spikes in the traffic because the rate is based on the average rate of the previous window.

Cons:
- It only works for not-so-strict look back window. It is an approximation of the actual rate because it assumes requests in the previous window are evenly distributed. However, this problem may not be as bad as it seems. According to experiments done by Cloudflare, only 0.003% of requests are wrongly allowed or rate limited among 400 million requests.
