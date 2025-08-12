-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS pessoa_juridica (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	id_cliente UUID NOT NULL UNIQUE REFERENCES cliente(id),
	data_criacao TIMESTAMPTZ NOT NULL,
	nome_fantasia VARCHAR(50) NOT NULL,
	cnpj VARCHAR(50) NOT NULL UNIQUE
);
---- create above / drop below ----
DROP TABLE IF EXISTS pessoa_juridica;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
