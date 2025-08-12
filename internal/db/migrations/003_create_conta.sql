-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS conta (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	saldo  DECIMAL(15,2) DEFAULT 0,
	id_cliente UUID NOT NULL REFERENCES cliente(id) ON DELETE CASCADE,
	data_abertura TIMESTAMPTZ DEFAULT now()
);
---- create above / drop below ----
DROP TABLE IF EXISTS conta;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
