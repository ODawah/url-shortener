URL Shortener
---------------------

## Technologies & Language
- Golang
- Chi
- Mysql
- Redis Caching

## Use Cases
- User Send URL to be Shortened
- User hits a shortened URL
- User wants to see metrics for his url

## Constraints
*State Constraints*
- URL max length is 2048 Bytes 
- shortened URL should be less than or equal 75 bytes (average number for URL length)



# Design
![url-shortener-design](url-shortener-design.png)
# System Architecture
The system Consists of the following components

## API
We will use REST API Architecture as it's simple, flexible and scalable since it's stateless and easy to use
The API will contain three services:
- Shorten the URL
- Redirect from short URL to the original
- Return Analytics for the shortened URL

## Services
### Shorten the URL
The Shorten URL Service will take the long URL and shorten it by using MD5 hashing then choose the first 7 letters
and hash them by base62 to turn it into short alphanumeric representation.
The MD5 hashing will be applied on user IP address and timestamp, we chose those parameters to reduce the collision rate.

After shortening the URL we store it in the database and map it with the original URL. but before storing into the database
we need to check if the generated URL is unique, so we first check the database if the generate URL collides with another already stored
if it collides generate another URL since it's based on time it'll change. if there's no collision store it in the database 
mapped with the original URL.

Why MD5 ?   
we decided to choose MD5 over other better algorithms like sha-256 even if they have higher collision
because it's faster since it returns 128-bit hash which is less than the better algorithms but also with an acceptable
collision rate


### Redirect from short URL to the original
The Service will take the Short URL and search for it in Redis Cache if it exisits redirect the user
the original URL mapped to it. If it doesn't exist in the Redis Cache, so we need to look for it in the database. The Cache will always contain the most frequently hit URLs, so for
the first few hits for the url the service will have to get the URL from Mysql until it got enough hits frequently to be stored in cache.
The Service Will Also Store The metadata for the requests the URL got for metrics calculations

### Return Analytics For The shortened URL
The Analytics returned will be calculated through query language the user will use to get full flexibility of the 
analytics he wants. The analytics for the URL will be calculated from the Requests metadata stored for URL.
We decided to store the metadata for all the requests the URL got to make the analytics flexible for the user
for example the user wants to know The most hits comes from any region ?, 
The hit rate for a specific time interval etc. we don't know what the user wants to see
,so we let him take control of what analytic he wants to see.


## Registration
The App will register users with email and password so every user will see his shortened URLs
and can keep track of its analytics but the for a certain time frame as the time goes the old data got deleted and the new data stored
but if the user pay for a premium account he will get analytics for the URLs in a larger time frame.

### Database
For Database, we decided to rely on MySQL and use Sharding to increase the performance of DB when data gets bigger

#### Schema 
![database-schema](database-schema.png)

our database schema is simple it has 4 entities:
- users 
- subscriptions
- URls
- Requests

we created the table users to separate users URls from each other, users table have column subscription_id
it's a foreign key to know if the user is subscribed to the premium and when it starts and ends. 
the table URLs is for storing the original URL mapped with the short one, the table have user_id column it's a foreign
key for the user that owns this short url.
the last table Requests is for storing the metadata of requests happen to URLs to calculate the metrics for URLs 


### Caching
In our caching strategy, we aim to optimize performance and reduce response time by leveraging a cache to store URLs with a high hit rate. When a user requests a short URL, we employ the following caching approach:

1. Storing URLs with High Hit Rate:

    - The cache is designed to store frequently accessed URLs that exhibit a high hit rate, indicating their popularity among users.
    - By storing these popular URLs in the cache, we can significantly improve response times by retrieving their associated content directly from the cache.

2. Time-to-Live (TTL) Management:

   - Each URL stored in the cache is assigned a Time-to-Live (TTL) value of 30 seconds initially.
   - When a URL is accessed again within its TTL period, its TTL is extended to maintain its presence in the cache.
   - This approach ensures that frequently accessed URLs remain readily available in the cache, minimizing the need to retrieve them from the underlying data source repeatedly.
   
3. Handling Cache Miss:

   - In the event of a cache miss (i.e., the requested URL is not found in the cache), we retrieve the URL's corresponding content from the database.
   - Once fetched from the database, the URL and its associated content are stored in the cache using the same TTL value as before (30 seconds).
   - This ensures that subsequent requests for the same URL can be served directly from the cache, further enhancing response times.

4. Cache Key-Value Structure:

   - The cache utilizes a key-value structure, where the short URL serves as the key, and the original URL is stored as the corresponding value.
   - For example, a key-value pair in the cache might be "bit.ly/sdXTCs": "www.facebook.com".
   - This key-value mapping enables efficient and quick lookups when serving requests for shortened URLs.

## Design Core Components
#### Use Case: User sends a URL to be shortened
- The Client sends a Request to API with long URL to be shortened
- The server recevies the request and send the URL to backend
- Backend Shorten the URL and store the short URL coupled with original in the database
  - URL Shortener function:
    - Shorten the URL with MD5 algorithm
    - Check if the shortened URL collides with another in database
    - Rehash the URL again until no collision happens
- The Server then Returns the shortened URL to the client

#### Use Case : User Hits a shortened URL
- The Client send a request to shortened URL
- The server receives the request and search in cache for the URL
- Search the database for URL if it's not in cache
- The database returns the original URL if it exisits 
- update the metrics of the url
- The server then redirect the client to original URL

#### Use Case : User wants to get metrics about URL 
The user will send the metrics he wants in query language to the service will then
compile the query and call the database to get and calculate the metrics data user requested
and return it.
