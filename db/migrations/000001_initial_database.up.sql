-- Tabel authors
CREATE TABLE
    authors (id TEXT PRIMARY KEY, name TEXT);

-- Tabel articles
CREATE TABLE
    articles (
        id TEXT PRIMARY KEY,
        author_id TEXT,
        title TEXT,
        body TEXT,
        created_at TIMESTAMP,
        CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE SET NULL
    );