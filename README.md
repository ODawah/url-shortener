URL Shortener
---------------------

## Technologies & Language
- Golang
- Chi
- Mongo DB

## Use Cases
- User Send URL to be Shortened
- User hits a shortened URL
- User wants to see metrics for his url

## Constraints
*State Constraints*
- URL max length is 2048 Bytes 
- shortened URL should be less than or equal 75 bytes (average number for URL length)



# Design
![design](design.png)
# System Architecture
The system Consists of the following components

### API
We will use REST API Architecture as it's simple, flexible and scalable since it's stateless and easy to use

### Back End
The Backend will contain three services:
- Shorten the URL
- Redirect from short URL to the original
- Return Analytics for the shortened URL

### Database
For Database, we decided to rely on NoSQL since we have high writes on database for analytics,
and we can't use relational database because storing metrics for urls will result in unmaintainable big tables
so the decision is to use MongoDB

### Caching
In Cache, we will store the URLs with high hit rate, when a user makes a request to a short URL
this URL will be stored in cache for future usages and for the TTL of data in cache will be based on the frequency of usage of URL
since the data rarely or never changes. we chose Redis for caching


## Design Core Components

#### Use Case: User sends a URL to be shortened
- The Client sends a Request to API with long URL to be shortened
- The server recevies the request and send the URL to backend
- Backend Shorten the URL and store the short URL coupled with original in the database
  - URL Shortener function:

    - Shorten the URL with MD5 or SHA1 algorithm
    - Check if the shortened URL collides with another in database but doesn't map to the same original URL
    - double hash the url or choose another algorithm for hashing
- The Server then Returns the shortened URL to the client

#### Use Case : User Hits a shortened URL
- The Client send a request to shortened URL
- The server receives the request and search in cache for the URL
- Search the database for URL if it's not in cache
  - Store the short url mapped with the original in cache
- The database returns the original URL if it exisits 
- update the metrics of the url
- The server then redirect the client to original URL

#### Use Case : User wants to get metrics about URL 
The user might want to know metrics for his shortened url how many times it has been hit
timeline of hits rate and most country hit the URL.
- The client send request to get metrics for URL.
- The server recevies request and search database for url
- The database retrieve the url metrics
- The server returns the metrics to the client

#### Use Case: User want to get time series of URL hit rate
The user want to get the day with the highest hit rate on his url compared with the other days
starts from the day he stored his url until the day he requests to see metrics
- The client send request to get metrics for URL.
- The server receives request and search database for url
- the metrics are retrieved and calculated
- the server returns the metrics to the client

### Shortcoming
- Increase in the data URL time metrics as it increases when time passes

