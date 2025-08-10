-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS conta (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	saldo DECIMAL DEFAULT 0,
	owner_id UUID NOT NULL REFERENCES usuario(id)
);
---- create above / drop below ----
DROP TABLE IF EXISTS conta;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
