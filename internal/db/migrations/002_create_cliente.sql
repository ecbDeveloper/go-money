-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS cliente (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	categoria_cliente INT REFERENCES categoria_cliente(id),
	password BYTEA NOT NULL,
	telefone VARCHAR(50) NOT NULL, 
	email VARCHAR(50) NOT NULL UNIQUE,
	data_cadastro TIMESTAMPTZ DEFAULT now()
);
---- create above / drop below ----
DROP TABLE IF EXISTS cliente;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
