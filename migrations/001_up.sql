CREATE TABLE products
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    mark INT          NOT NULL CHECK (mark >= 1 AND mark <= 10)
);

CREATE TABLE categories
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products_categories
(
    product_id  INT REFERENCES products (id),
    category_id INT REFERENCES categories (id),
    PRIMARY KEY (product_id, category_id)
);

INSERT INTO products (id, name, mark)
values (default, 'Роберт Мартин - Чистый код', 8),
       (default, 'Роберт Мартин - Чистая архитектура', 10),
       (default, 'Стив Макконнелл - Идеальный код', 9),
       (default, 'Лопата', 8),
       (default, 'Грабли', 9),
       (default, 'Baldur’s Gate III', 8),
       (default, 'Elden Ring', 10),
       (default, 'Ведьмак', 9);

INSERT INTO categories (id, name)
values (default, 'Книги'),
       (default, 'Инструменты'),
       (default, 'Игры');

INSERT INTO products_categories (product_id, category_id)
VALUES (1, 1),
       (2, 1),
       (3, 1),
       (4, 2),
       (5, 2),
       (6, 3),
       (7, 3),
       (8, 1),
       (8, 3);