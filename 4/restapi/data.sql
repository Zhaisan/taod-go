CREATE TABLE public.author (
                               id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             name VARCHAR(100) NOT NULL,
                             author_id UUID NOT NULL,
                             CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);

INSERT INTO author (name) VALUES ('Zhaiss Sars'); -- 1dc72b22-7730-45e7-8edb-27d303c2ac1c
INSERT INTO author (name) VALUES ('Cristiano Ronaldo'); -- 029a0416-1e95-4cf8-a105-9abb71e5ae05
INSERT INTO author (name) VALUES ('Keylor Navas'); -- a7c09359-935b-4553-9e50-a58423143594

INSERT INTO book (name, author_id)  VALUES ('Мартин Иден', '1dc72b22-7730-45e7-8edb-27d303c2ac1c');
INSERT INTO book (name, author_id)  VALUES ('Золотой Мяч', '029a0416-1e95-4cf8-a105-9abb71e5ae05');
INSERT INTO book (name, author_id)  VALUES ('Золотые перчатки', 'a7c09359-935b-4553-9e50-a58423143594');





