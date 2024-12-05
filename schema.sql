CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,          -- ID auto-incrementável
    name TEXT NOT NULL,             -- Nome do usuário
    birth_date DATE NOT NULL,       -- Data de nascimento
    cpf TEXT UNIQUE NOT NULL,       -- CPF, único para cada usuário
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Data de criação
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Data de última atualização
);

-- Criação de índices para melhorar a performance de buscas
CREATE INDEX idx_user_cpf ON "user" (cpf);
