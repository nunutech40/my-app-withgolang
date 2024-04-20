
# RUN and connect to MYSQL

## run docker mysql
```bash
docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=myadmin12345678 -p 3306:3306 -v /Users/nununugraha/Documents/Programming/LearnGo/my-app/mysqldb:/var/lib/mysql mysql
```

## Re-run docker already created
```bash
docker start mysql-container
```

## Stop docker mysql
```bash
docker stop mysql-container
```
## cek docker running
```bash
docker ps -a
```

## delete container 
```bash
docker rm mysql-container
```

## connect to mysql
```bash
docker exec -it mysql-container mysql -u root -p
```

## After connected
```bash
SHOW DATABASES;

CREATE DATABASE exampledb;

USE exampledb;

SHOW TABLES;

CREATE TABLE exampletable (
    id INT AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    primary key(id)
);

INSERT INTO exampledb (name) VALUES ('Test Values');

SELECT * FROM exampletable;


CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY ,
    username VARCHAR(50) UNIQUE NOT NULL ,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP NULL DEFAULT NULL,
    role ENUM('user', 'admin') DEFAULT 'user',
    status ENUM('active', 'inactive') DEFAULT 'active'
) ENGINE=InnoDB;

INSERT INTO users (username, password_hash, email, role) VALUES 
('user', 'hash_of_userpass123', 'user@example.com', 'user'),
('admin', 'hash_of_adminpass123', 'admin@example.com', 'admin');


```