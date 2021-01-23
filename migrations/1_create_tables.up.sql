CREATE TABLE addresses (
    id UUID PRIMARY KEY,
    address_1 TEXT NOT NULL,
    address_2 TEXT,
    city TEXT NOT NULL,
    us_state TEXT NOT NULL,
    zipcode TEXT NOT NULL,
    geocode TEXT
);

CREATE TABLE manufacturers (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    active BOOLEAN,
    address_id UUID REFERENCES addresses (id) ON DELETE CASCADE
);

CREATE TABLE series (
    id UUID PRIMARY KEY,
    manufacturer_id UUID NOT NULL REFERENCES manufacturers (id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    series_number TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE models (
    id UUID PRIMARY KEY,
    manufacturer_id UUID NULL REFERENCES manufacturers (id) ON DELETE CASCADE,
    series_id UUID NOT NULL REFERENCES series (id) ON DELETE CASCADE,
    model_number TEXT NOT NULL,
    image BYTEA
);

CREATE TABLE distributors (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    active BOOLEAN,
    address_id UUID REFERENCES addresses (id) ON DELETE CASCADE
);

CREATE TABLE branches (
    id UUID PRIMARY KEY,
    distributor_id UUID NOT NULL REFERENCES distributors (id) ON DELETE CASCADE,
    active BOOLEAN,
    address_id UUID REFERENCES addresses (id) ON DELETE CASCADE
);

CREATE TABLE contacts (
    id UUID PRIMARY KEY,
    department TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone TEXT,
    email TEXT,
    active BOOLEAN
);
