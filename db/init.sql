-- CREATE DATABASE IF NOT EXISTS peya
-- SELECT 'CREATE DATABASE userDB'
--    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'userDB')\gexec

-- Crear la tabla persons
CREATE TABLE persons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    dni INT NOT NULL,
    birthdate DATE NOT NULL
);

-- Crear la tabla users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(250) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    role VARCHAR(50) NOT NULL,
    phone VARCHAR(30)
);

-- Crear la clave foránea de users que referencia a persons
ALTER TABLE users
ADD CONSTRAINT fk_person
FOREIGN KEY (id) REFERENCES persons(id)
ON DELETE CASCADE;