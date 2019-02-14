CREATE DATABASE sample;

CREATE TABLE sample.users(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(25) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO sample.users (name, age, gender) VALUES ("James Bond", 43, "male");