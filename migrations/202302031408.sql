CREATE TABLE IF NOT EXISTS  users (
    id    serial PRIMARY KEY,

    login  varchar(30) NOT NULL UNIQUE,
    email varchar(30) NOT NULL UNIQUE,
    password varchar(100) NOT NULL,

    created_at date DEFAULT CURRENT_DATE NOT NULL, // ебаный верблюжий шрифт 
    updated_at date,
    deleted_at date
);

CREATE TABLE IF NOT EXISTS posts (
    id    serial PRIMARY KEY,// id можно ебануть проще а не записью из скрижали

    title  varchar(30) NOT NULL,
    content text NOT NULL,
    image varchar(100),
    viewCount int NOT NULL DEFAULT 0,// NOT NULL DEFAULT 0 )))))))))))))))))))))))))))))))))))))))))))))
    likes int NOT NULL DEFAULT 0, // почему viewCount но просто likes?
    author_id int references users NOT NULL , // 2 ебаных аналогичных референса

    dateCreate date DEFAULT CURRENT_DATE NOT NULL,
    dateUpdate date,
    dateRemove date,

    FOREIGN KEY (authorId) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS tokens (
   id    serial PRIMARY KEY, // ебало аксесс токенов в бд?

    value varchar(255) UNIQUE,
    userId int REFERENCES users NOT NULL,
    dateCreate date DEFAULT CURRENT_DATE,

    FOREIGN KEY(userId) REFERENCES users (id)
);

CREATE UNIQUE INDEX tokenIndex ON tokens (value);
