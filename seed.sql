CREATE
DATABASE bird_encyclopedia;

CREATE TABLE birds
(
    id          SERIAL PRIMARY KEY,
    bird        VARCHAR(256),
    description VARCHAR(1024)
);

INSERT INTO birds (bird, description)
VALUES ('pigeon', 'common in cities');
INSERT INTO birds (bird, description)
values ('eagle', 'bird of prey');