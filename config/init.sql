CREATE TABLE gpdf.`user`
(
    id       INT UNSIGNED auto_increment NOT NULL,
    name     varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    CONSTRAINT user_pk PRIMARY KEY (id),
    CONSTRAINT user_unique UNIQUE KEY (name)
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
