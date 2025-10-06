ALTER TABLE public.categories
  DROP COLUMN IF EXISTS created_at,
  DROP COLUMN IF EXISTS updated_at;