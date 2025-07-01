package imagemeta

const (
	QueryInsertImageMeta = "INSERT INTO images_meta (id, name, format ,mime_type, size, uploaded_at) values ($1,$2,$3,$4,$5,$6)"
	QueryGetMaxId        = "SELECT MAX(id) FROM images_meta"
	QueryGetRowWithId    = "SELECT * FROM images_meta WHERE id=?"
	QueryDeleteRowWithId = "DELETE FROM images_meta WHERE id=?"
)
