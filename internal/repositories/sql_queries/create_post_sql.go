package sql_queries

// TODO: Узнать можно ли запросы выность в отдельный пакет
const CreatePost = `
	INSERT INTO posts (id, title, content, author_id, image_url, created_at, updated_at)
	VALUES (${id}, ${title}, ${content}, ${author_id}, ${image_url}, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
`
