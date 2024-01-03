# Token bucket

We can break that algorithm into three steps
- Fetch the token
- Access the token
- Update the token

Consider we have planned to limit our API 10 requests/minute. For every unique user, we can track last time they have accessed the API and available token(number) of that user. These things can be stored in the Redis server with the key as user id.
- Fetch the token: So when a request comes from the user, then our rate-limiting algorithm first fetch the user’s token using the user id.
- Access the token: When we have enough tokens for the particular user, then it will allow processing otherwise it blocks the request to hit the API. The last access time and remaining tokens of that specific user will be changed.
- Update Token: Rate limiting algorithm will update the token in the Redis as well. One more thing, the tokens will be restored after the time interval. In our scenario, ten will be updated after the time period as a token value.


Pros:
- Token Bucket algorithm is very simple and easy to implement.
- Token Bucket algorithm is very memory efficient.
- Token Bucket technique allows spike in traffic or burst of traffic. A request goes through as long as there are tokens left. This is super important since traffic burst is not uncommon. One example is events like Amazon Prime Day when traffic spikes for a certain time period.

Cons:
- A race condition, as described above, may cause an issue in a distributed system due to concurrent requests from the same user.
  (We could use redis lock or operate through Lua scripting to avoid this situation.)


