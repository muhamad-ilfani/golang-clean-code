CREATE TABLE IF NOT EXISTS project2.users(
    id BIGSERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password BYTEA NOT NULL,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    created_at timestamp NOT NULL DEFAULT now(),
    modify_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    modify_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    deleted_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT username_uk UNIQUE(user_name)
);

CREATE TABLE IF NOT EXISTS project2.products(
    id BIGSERIAL PRIMARY KEY,
    user_id int8 NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    created_at timestamp NOT NULL DEFAULT now(),
    modify_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    modify_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    deleted_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT user_fk_products FOREIGN KEY (user_id) REFERENCES project2.users (id)
);