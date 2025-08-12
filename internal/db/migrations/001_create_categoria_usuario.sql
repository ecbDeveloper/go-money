-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS categoria_cliente (
	id SERIAL PRIMARY KEY,
	categoria VARCHAR(50)
);

INSERT INTO categoria_cliente (categoria) 
VALUES ('Pessoa Física'), ('Pessoa Jurídica');
---- create above / drop below ----
DROP TABLE IF EXISTS categoria_cliente;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
