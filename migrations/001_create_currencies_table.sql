CREATE TABLE currencies (
                            id serial PRIMARY KEY,
                            code VARCHAR ( 255 ) UNIQUE NOT NULL,
                            full_name VARCHAR ( 255 ) NOT NULL,
                            sign VARCHAR ( 255 ) UNIQUE NOT NULL,
                            created_at timestamp with time zone,
                            updated_at timestamp with time zone,
                            deleted_at timestamp with time zone
)