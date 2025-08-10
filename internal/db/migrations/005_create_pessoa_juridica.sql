-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS pessoa_juridica (
	id UUID PRIMARY KEY REFERENCES pessoa(id),
	faturamento DECIMAL NOT NULL,
	idade INT NOT NULL,
	nome_fantasia VARCHAR(255) NOT NULL,
	celular VARCHAR(255) NOT NULL,
	email_corporativo VARCHAR(255) NOT NULL
);
---- create above / drop below ----
DROP TABLE IF EXISTS pessoa_juridica;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
