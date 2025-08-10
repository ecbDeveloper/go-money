-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS pessoa_fisica (
	id UUID PRIMARY KEY REFERENCES usuario(id),
	renda_mensal DECIMAL NOT NULL,
	idade INT NOT NULL,
	nome_completo VARCHAR(255) NOT NULL,
	celular VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL
);
---- create above / drop below ----
DROP TABLE IF EXISTS pessoa_fisica;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
