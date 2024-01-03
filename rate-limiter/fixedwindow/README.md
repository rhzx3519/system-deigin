# Fixed Window

For each request a user makes,
1. Check if the user has exceeded the limit in the current window.
2. If the user has exceeded the limit, the request is dropped
3. Otherwise, we increment the counter


Pros
- Easy to implement.

Cons
- A burst of requests at the end of the window causes server handling more requests than the limit since this algorithm will allow requests for both current and next window requests within a short time. For example, for 100 req/min, if the user makes 100 requests at 55 to 60 seconds window and then 100 requests from 0 to 5 seconds in the next window, this algorithm handles all the requests. Thus, ends up handling 200 requests in 10 seconds for this user, which is above the limit of 100 req/min
