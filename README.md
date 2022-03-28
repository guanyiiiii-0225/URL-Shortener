# URL-Shortener
## Description
The project is a simple URL shortener that allows you to shorten your long URLs.
## Settings
1. Clone this project<br>
```
git clone https://github.com/guanyiiiii-0225/URL-Shortener.git
```
2. Copy `.env.example` file to create `.env` file
3. Create `URL` database in PostgreSQL
4. Edit `.env` file and fill out your PostgreSQL username and password<br>
```
PG_HOST=localhost
PG_PORT=5432
PG_USERNAME=
PG_PASSWORD=
PG_DBNAME=URL
DOMAIN_NAME=localhost
```
5. Start backend service<br>
```
go run .
```
6. Test<br>
you can run `http://localhost:8080/test` to test it.
## Usage
### Upload URL API:
1. "url" requires a valid URL. (Eg: https://www.youtube.com/)
2. "expireAt" requires RFC 3399 UTC format.
```
curl -X POST -H "Content-Type:application/json" http://localhost/api/v1/urls -d '{
    "url": "https://steveyi.net/",
    "expireAt": "2021-02-08T09:20:41Z"
}'
```
+ Response:
1. "id" will generate a 10 character random string.
2. "shortUrl" will generate a short URL.
```
{
    "id": "<url_id>",
    "shortUrl": "http://localhost/<url_id>"
}
```
### Redirect URL API:
```
curl -L -X GET http://localhost/<url_id> => REDIRECT to original URL # HTTP 301
```

### If there are some errors, reponse will be:
```
{
    "status": "failed",
    "error":  "<error msg>"
}
```

# Technical details
## 3rd Party Library Usage
+ Web Engineering
  + Use gin for web framework.
+ Database
  + Use PostgreSQL for database.
## Features
+ Init
  + Check whether the table exists or not. If not, create it.
+ Upload URL API
  + HTTP POST with JSON body.
  + Check whether url is valid or not.
    +  require a valid URL: Use `url.ParseRequestURI` to check.
+ Redirect URL API
  + HTTP GET with URL query.
  + Redirect to orginal url.
