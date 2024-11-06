package sql_queries

const GetPosts = `
	SELECT id, title, content, author_id, image_url, created_at, updated_at FROM posts
`
