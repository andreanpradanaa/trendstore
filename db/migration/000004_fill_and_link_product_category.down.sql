ALTER TABLE public.products
  DROP CONSTRAINT IF EXISTS fk_products_category,
  ALTER COLUMN category_id DROP NOT NULL;