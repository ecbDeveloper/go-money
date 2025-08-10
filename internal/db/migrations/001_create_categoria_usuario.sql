-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS categoria_usuario (
	id SERIAL PRIMARY KEY,
	categoria VARCHAR(50)
);

INSERT INTO categoria_usuario (categoria) 
VALUES ('Pessoa Física'), ('Pessoa Jurídica');
---- create above / drop below ----
DROP TABLE IF EXISTS categoria_usuario;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
