package sql_queries

const PatchPost = `
	UPDATE posts
	SET title = COALESCE(${title}, title),
		content = COALESCE(${content}, content),
		author_id = COALESCE(${author_id}, author_id),
		image_url = COALESCE(${image_url}, image_url),
		updated_at = CURRENT_TIMESTAMP
	WHERE id = ${id};

`
