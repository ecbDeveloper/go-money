-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS status_conta( 
	id SERIAL PRIMARY KEY,
	status VARCHAR(50)
);

INSERT INTO status_conta (status)
VALUES ('ativa'), ('desativa');
---- create above / drop below ----
DROP TABLE IF EXISTS status_conta;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
