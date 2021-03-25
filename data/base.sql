--- POSTGRESQL BÁSE TABLE

-- Start create new schema
CREATE SCHEMA IF NOT EXISTS basic_crud;

-- ###########################################################################
-- Book table basic configurate

CREATE TABLE basic_crud."book" (
    id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(255),
    subtitle TEXT,
    price INTEGER,
    publication_date TIMESTAMPTZ,
    image BYTEA,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


INSERT INTO basic_crud."book" (title, subtitle, price) 
VALUES ('Start book', 'Book of start database migration', 2000);

-- SELECT * FROM basic_crud."book";

-- ###########################################################################
-- User table basic configurate

CREATE TABLE basic_crud."user" (
    id SERIAL NOT NULL PRIMARY KEY,
	"name" varchar NOT NULL,
    admin boolean NOT NULL, 
	nickname varchar NOT NULL,
	email text UNIQUE,
	passworld varchar NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


INSERT INTO basic_crud."user" ("name", nickname, email, admin, passworld)
VALUES('root', 'root', 'root@root.com', true, md5('masteruser'));

-- SELECT "name", nickname, age, email
-- FROM basic_crud."user";

-- DELETE FROM basic_crud."user"
-- WHERE email='';


-- SELECT * FROM basic_crud."user";