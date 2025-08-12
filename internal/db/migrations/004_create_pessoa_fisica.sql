-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS pessoa_fisica (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	id_cliente UUID NOT NULL UNIQUE REFERENCES cliente(id),
	nome_completo VARCHAR(255) NOT NULL,
	data_nascimento TIMESTAMPTZ NOT NULL,
	cpf VARCHAR(20) NOT NULL UNIQUE
);
---- create above / drop below ----
DROP TABLE IF EXISTS pessoa_fisica;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
