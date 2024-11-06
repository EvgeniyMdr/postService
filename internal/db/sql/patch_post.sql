UPDATE posts
SET title = COALESCE($1, title),
    content = COALESCE($2, content),
    author_id = COALESCE($3, author_id),
    image_url = COALESCE($4, image_url),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $5;
