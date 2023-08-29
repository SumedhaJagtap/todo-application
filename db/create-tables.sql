DROP DATABASE IF EXISTS todo;
CREATE DATABASE todo;
use todo;

DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks(
    ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(128) NOT NULL,
    Description VARCHAR(128) NOT NULL,
    Status INT,
    CreatedAt VARCHAR(128) NOT NULL,
    ModifiedAt VARCHAR(128)
)