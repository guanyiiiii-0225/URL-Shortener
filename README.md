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
## 3rd Party Library Usage
+ Web Engineering
  + Use gin for web framework.
+ Database
  + Use PostgreSQL for database.
