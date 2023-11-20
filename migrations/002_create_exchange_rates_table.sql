CREATE TABLE exchange_rates (
                                id serial PRIMARY KEY,
                                base_currency_id int NOT NULL,
                                target_currency_id int NOT NULL,
                                rate NUMERIC NOT NULL,
                                FOREIGN KEY (base_currency_id)
                                    REFERENCES currencies (id),
                                FOREIGN KEY (target_currency_id)
                                    REFERENCES currencies (id)
);
