CREATE TABLE users (
    id UUID PRIMARY KEY,
    phone VARCHAR(20) NOT NULL,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    gender BOOLEAN NOT NULL,
    birthday DATE NOT NULL,
    subregion VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL
);

CREATE TABLE bot (
    id VARCHAR(20) NOT NULL,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    isforum BOOLEAN NOT NULL,
    username VARCHAR(255) NOT NULL,
    languagecode VARCHAR(255) NOT NULL,
    isbot BOOLEAN NOT NULL,
    ispremium BOOLEAN NOT NULL,
    addedtomenu BOOLEAN NOT NULL,
    usernames VARCHAR[],
    customemojistatus VARCHAR(255) NOT NULL,
    canjoingroups BOOLEAN NOT NULL,
    canreadmessages BOOLEAN NOT NULL,
    supportsinline BOOLEAN NOT NULL
);

CREATE TABLE sms (
    username VARCHAR(255) NOT NULL,
    sms INTEGER NOT NULL
);
