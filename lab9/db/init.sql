CREATE TABLE person (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT CHECK(age > 0) NOT NULL,
    phone TEXT CHECK (phone ~ '^(\+7|8)[0-9]{10}$')
);

CREATE TABLE sys_user (
    id SERIAL PRIMARY KEY,
    person_id INT REFERENCES person(id) NOT NULL,
    password_hash CHAR(60) NOT NULL
);

CREATE TABLE audit (
    id SERIAL PRIMARY KEY,
    action TEXT,
    target TEXT,
    level TEXT,
    description TEXT,
    timestamp TIMESTAMP DEFAULT NOW()
);