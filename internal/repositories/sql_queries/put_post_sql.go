package sql_queries

const PutPost = `
	UPDATE posts
	SET title = ${title},
		content = ${content},
		author_id = ${author_id},
		image_url = ${image_url},
		updated_at = CURRENT_TIMESTAMP
	WHERE id = ${id};
`
