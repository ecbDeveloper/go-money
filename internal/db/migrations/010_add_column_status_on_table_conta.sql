-- Write your migrate up statements here
ALTER TABLE conta
ADD COLUMN status INT REFERENCES status_conta(id) DEFAULT 1;
---- create above / drop below ----
ALTER TABLE conta
DROP COLUMN status;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
