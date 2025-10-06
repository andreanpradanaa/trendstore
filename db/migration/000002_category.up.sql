CREATE TABLE public.categories (
    id          SERIAL PRIMARY KEY,         
    name        varchar(50)  NOT NULL,
    description varchar(255) NOT NULL,
    CONSTRAINT categories_category_name_key UNIQUE (name)
);