-- +goose Up
CREATE TABLE IF NOT EXISTS files (
    id TEXT PRIMARY KEY,
    created_at INTEGER DEFAULT (strftime('%s', 'now')),
    updated_at INTEGER DEFAULT (strftime('%s', 'now')),
    filename TEXT NOT NULL,
    size INTEGER,
    thumbnail_url TEXT,
    file_url TEXT NOT NULL,
    file_type TEXT,
    aspect_ratio TEXT,
    title TEXT,
    description TEXT
);

CREATE TABLE IF NOT EXISTS thumbnails (
    id TEXT PRIMARY KEY,
    created_at INTEGER DEFAULT (strftime('%s', 'now')),
    updated_at INTEGER DEFAULT (strftime('%s', 'now')),
    thumbnail_url TEXT NOT NULL,
    file_id TEXT NOT NULL,
    FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_thumbnails_file_id ON thumbnails(file_id);

-- +goose Down
DROP INDEX IF EXISTS idx_thumbnails_file_id;
DROP TABLE IF EXISTS thumbnails;
DROP TABLE IF EXISTS files;
