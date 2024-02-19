CREATE TABLE Missao (
    id_missao SERIAL PRIMARY KEY,
    nome_missao VARCHAR(100), 
    rank_missao CHAR
);

CREATE TABLE Aventureiro (
    id_avent SERIAL PRIMARY KEY,
    nome_avent VARCHAR(100),
    rank_avent CHAR
);

CREATE TABLE Aventura (
    id_aventura SERIAL PRIMARY KEY,
    id_missao INT,
    nome_missao VARCHAR(100),
    id_avent INT,
    nome_avent VARCHAR(100)
);
-- Dados iniciais na tabela Missao
INSERT INTO Missao (nome_missao, rank_missao)
VALUES ('Resgate da princesa elfa', 'A');

-- Dados iniciais na tabela Aventureiro
INSERT INTO Aventureiro (nome_avent, rank_avent)
VALUES ('Eragon', 'A');
-- Dados iniciais na tabela Aventura
INSERT INTO Aventura (id_missao, nome_missao, id_avent, nome_avent)
VALUES (1, 'Resgate da princesa elfa', 1, 'Eragon');
CREATE ROLE estagiario with login 'projeto_de_estagio';
GRANT ALL PRIVILEGES ON DATABASE bd_aventura TO estagiario;


