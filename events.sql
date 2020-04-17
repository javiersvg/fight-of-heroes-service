CREATE TABLE EVENTS (
    EVENT_TYPE Varchar(255),
    AGGREGATE_ID Varchar(255),
    EVENT_DATA Varchar(255)
);

INSERT INTO EVENTS VALUES ('FightCreated', 'b837e4a9-458e-465e-9eb3-29029961aea8', '["SuperMan","Wonder Woman"]');
INSERT INTO EVENTS VALUES ('HeroesUpdated', 'b837e4a9-458e-465e-9eb3-29029961aea8', '["Superman","Wonder Woman"]');