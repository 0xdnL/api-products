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


CREATE DATABASE IF NOT EXISTS inventory;
USE inventory;
CREATE TABLE IF NOT EXISTS products(
  id INT NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  quantity int,
  price float(10,7),
  PRIMARY KEY(id)
  );
INSERT INTO products values(1, "chair", 100, 200.00);
INSERT INTO products values(2, "table", 150, 220.00);
INSERT INTO products values(3, "lamp", 80, 50.00);
```

## todo

- try out ORM
