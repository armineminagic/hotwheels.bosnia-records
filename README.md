# DiecastModelRecords
![diecastapp](https://user-images.githubusercontent.com/37677214/123234543-4fff7480-d4db-11eb-89d7-d97234b92819.png)

Simple website for recording information about my Diecast car models collection

## Database

App is using hwrecords database so we need to create one:

```
CREATE DATABASE hwrecords;
USE hwrecords;
CREATE TABLE hwmodels
(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(64) NOT NULL,
    `set` varchar(64) NOT NULL,
    year int(11) NOT NULL,
    manufacturer varchar(32) NOT NULL,
    model_number int NOT NULL,
    PRIMARY KEY (id)
    UNIQUE KEY namemodnum (name, model_number),
    KEY (name),
    KEY (model_number)
);
```
We use `root` user with `root` password to access the database.
If you don't have the same you should create it and grant all privileges to it.

```
CREATE USER 'root'@'localhost' IDENTIFIED BY 'root';
```
```
GRANT ALL PRIVILEGES ON hwrecords.hwmodels TO 'root'@'localhost';
```
