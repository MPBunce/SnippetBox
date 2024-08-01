-- Create a `snippets` table.
CREATE TABLE snippets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);

-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.',
    CURRENT_TIMESTAMP,
    DATETIME(CURRENT_TIMESTAMP, '+365 days')
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n N',
    CURRENT_TIMESTAMP,
    DATETIME(CURRENT_TIMESTAMP, '+365 days')
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning \n the mirror I stare into\nshows my father''s face.\n',
    CURRENT_TIMESTAMP,
    DATETIME(CURRENT_TIMESTAMP, '+7 days')
);


SELECT * FROM snippets

SELECT id, title, content, created, expires
FROM snippets
WHERE expires > strftime('%Y-%m-%d %H:%M:%S', 'now')
  AND id = 2;


SELECT id, title, content, created, expires FROM snippets
WHERE datetime(expires) > datetime('now') ORDER BY created DESC LIMIT 10

CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

CREATE UNIQUE INDEX users_uc_email ON users (email);

