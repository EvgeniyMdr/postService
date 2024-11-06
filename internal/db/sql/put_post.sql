UPDATE posts
SET title = $1,
    content = $2,
    author_id = $3,
    image_url = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $5;