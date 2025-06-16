package metadb

const Metatable string = `CREATE TABLE IF NOT EXISTS images_meta (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	format TEXT NOT NULL,
	mime_type TEXT NOT NULL,
	size INTEGER NOT NULL,
	upload_at TEXT NOT NULL
);`
