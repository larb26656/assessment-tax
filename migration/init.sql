CREATE TABLE tax_deduction_setting (
    key VARCHAR(255) PRIMARY KEY,
    value FLOAT8
);

INSERT INTO
    tax_deduction_setting ("key", value)
VALUES
    ('personal', 60000),
    ('k-receipt', 50000);

;