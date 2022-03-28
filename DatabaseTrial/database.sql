CREATE TABLE Person (
    PersonId int,
    FirstName varchar(100),
    LastName varchar(255)
);

DROP TABLE Person;

INSERT INTO Person (FirstName, LastName)
VALUES ("Golang", "Playground");

INSERT INTO Person (FirstName, LastName)
VALUES ("Hello", "World");

CREATE TABLE Books (
    bookid int,
    BookName TEXT(100),
    BookUrl TEXT(2048),
);