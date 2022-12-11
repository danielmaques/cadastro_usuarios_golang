CREATE DATABASE IF NOT EXISTS fit_api;
USE fit_api;
DROP TABLE IF EXISTS usuarios;
CREATE TABLE usuarios(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;