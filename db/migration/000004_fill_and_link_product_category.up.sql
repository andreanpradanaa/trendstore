INSERT INTO categories(id, name, description) VALUES (999, 'Uncategorized', 'Default Category')
ON CONFLICT (id) DO NOTHING;

UPDATE products 
SET category_id = 999
WHERE category_id IS NULL;

ALTER TABLE products
    ALTER COLUMN category_id SET NOT NULL,
    ADD CONSTRAINT fk_products_category
        FOREIGN KEY (category_id)
        REFERENCES categories(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE;