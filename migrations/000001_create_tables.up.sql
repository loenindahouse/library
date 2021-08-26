CREATE TABLE IF NOT EXISTS authors(
    id bigserial PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    user_name VARCHAR NOT NULL,
    specialization VARCHAR NOT NULL
);
CREATE TABLE IF NOT EXISTS reader(
    id bigserial PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    book_list VARCHAR NOT NULL
);
CREATE TABLE IF NOT EXISTS books(
    id bigserial PRIMARY KEY,
    title VARCHAR NOT NULL,
    genre VARCHAR NOT NULL,
    ISBN bigserial UNIQUE,
    AuthorID VARCHAR NOT NULL
);
CREATE TABLE IF NOT EXISTS book_reader(
    book_id int REFERENCES books(id),
    reader_id int REFERENCES reader(id),
    CONSTRAINT id PRIMARY KEY (book_id, reader_id)
);








