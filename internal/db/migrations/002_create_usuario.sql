-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS usuario (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	categoria INT REFERENCES categoria(id)
);
---- create above / drop below ----
DROP TABLE IF EXISTS usuario;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
