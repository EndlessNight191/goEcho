CREATE TABLE IF NOT EXISTS  users (
    id    serial PRIMARY KEY,

    login  varchar(30) NOT NULL UNIQUE,
    email varchar(30) NOT NULL UNIQUE,
    password varchar(100) NOT NULL,

    created_at date DEFAULT CURRENT_DATE NOT NULL,
    updated_at date DEFAULT CURRENT_DATE NOT NULL,
    deleted_at date
);

CREATE TABLE IF NOT EXISTS posts (
    id    serial PRIMARY KEY,

    title  varchar(30) NOT NULL,
    content text NOT NULL,
    image varchar(100),
    view_count int DEFAULT 0,
    likes_count int DEFAULT 0,
    author_id int references users(id) ON DELETE CASCADE NOT NULL,

    created_at date DEFAULT CURRENT_DATE NOT NULL,
    updated_at date DEFAULT CURRENT_DATE NOT NULL,
    deleted_at date

);

CREATE TABLE IF NOT EXISTS tokens (
    id    serial PRIMARY KEY,

    value varchar(255) UNIQUE,
    user_id int references users(id) ON DELETE CASCADE NOT NULL,
    created_at date DEFAULT CURRENT_DATE NOT NULL
);


