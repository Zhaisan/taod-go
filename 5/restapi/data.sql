CREATE TABLE public.author (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book_authors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID NOT NULL,
    author_id UUID NOT NULL,

    CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book(id),
    CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id),
    CONSTRAINT book_author_unique UNIQUE (book_id, author_id)
);

INSERT INTO author (id, name) VALUES ('1dc72b22-7730-45e7-8edb-27d303c2ac1c', 'Zhaiss Sars');
INSERT INTO author (id, name) VALUES ('029a0416-1e95-4cf8-a105-9abb71e5ae05', 'Cristiano Ronaldo');
INSERT INTO author (id, name) VALUES ('a7c09359-935b-4553-9e50-a58423143594', 'Keylor Navas');

INSERT INTO book (id, name)  VALUES ('1dc72b22-7730-45e7-8edb-27d303c2ac1c', 'Мартин Иден');
INSERT INTO book (id, name)  VALUES ('029a0416-1e95-4cf8-a105-9abb71e5ae05', 'Золотой Мяч');
INSERT INTO book (id, name)  VALUES ('a7c09359-935b-4553-9e50-a58423143594', 'Золотые перчатки');

-- Мартин Иден
INSERT INTO book_authors (book_id, author_id) VALUES ('1dc72b22-7730-45e7-8edb-27d303c2ac1c', '1dc72b22-7730-45e7-8edb-27d303c2ac1c');
INSERT INTO book_authors (book_id, author_id) VALUES ('1dc72b22-7730-45e7-8edb-27d303c2ac1c', '029a0416-1e95-4cf8-a105-9abb71e5ae05');

-- Золотой мяч
INSERT INTO book_authors (book_id, author_id) VALUES ('029a0416-1e95-4cf8-a105-9abb71e5ae05', '029a0416-1e95-4cf8-a105-9abb71e5ae05');
INSERT INTO book_authors (book_id, author_id) VALUES ('029a0416-1e95-4cf8-a105-9abb71e5ae05', 'a7c09359-935b-4553-9e50-a58423143594');



SELECT b.id, b.name, array(SELECT ba.author_id FROM book_authors ba WHERE ba.book_id = b.id)
    AS authors FROM book b GROUP BY b.id, b.name;


SELECT a.id, a.name FROM book_authors
    JOIN public.author a on a.id = book_authors.author_id
        WHERE book_id = '029a0416-1e95-4cf8-a105-9abb71e5ae05';