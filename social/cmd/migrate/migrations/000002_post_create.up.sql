CREATE TABLE if not exists posts
(
    id      SERIAL PRIMARY KEY,
    title   VARCHAR(255),
    content TEXT,
    user_id INTEGER ,
    tags    TEXT[]
);