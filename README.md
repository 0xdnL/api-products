# api-products

> example api for playing around

## run

```sh
go run main.go

curl localhost:8080
```


## database

```sh
docker-compose -f docker-compose.yaml up -d
docker-compose -f docker-compose.yaml ps
docker-compose -f docker-compose.yaml exec -it db mysql -uroot -p
```

quick setup for example data

```sql
CREATE DATABASE IF NOT EXISTS learning;
USE learning;
CREATE TABLE data(id INT PRIMARY KEY, data VARCHAR(255));
DESC data;
INSERT INTO data values(1, "abc");
INSERT INTO data values(2, "def");
INSERT INTO data values(3, "ghi");
```
