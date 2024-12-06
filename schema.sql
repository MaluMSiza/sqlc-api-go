CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,          
    name TEXT NOT NULL,         
    birth_date DATE NOT NULL,     
    cpf TEXT UNIQUE NOT NULL,       
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   
);

CREATE INDEX idx_user_cpf ON "user" (cpf);
