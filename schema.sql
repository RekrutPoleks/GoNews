DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL
                   ON DELETE CASCADE
                   ON UPDATE CASCADE,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL
);

INSERT INTO authors (id, name) VALUES (0, 'Дмитрий');
INSERT INTO posts (id, author_id, title, content, created_at) VALUES (0, 0, 'Статья', 'Содержание статьи', 0);

DO
$$
BEGIN
FOR j IN 1..20 LOOP
		INSERT INTO authors (name) VALUES ('Author ' || j);
        FOR i iN 1..100 LOOP
			INSERT INTO posts (author_id, title, content, created_at) VALUES (J, 'Статья '  || i, 'Содержание статьи', round(extract(epoch from now()), 0));
        END LOOP;
END LOOP;
END
$$;