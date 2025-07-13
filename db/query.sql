-- name: CreateArticle :one
INSERT INTO articles (id, author_id, title, body, created_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    NOW()
)
RETURNING *;


-- name: CountTotalArticle :exec
SELECT COUNT(*) as total_article FROM articles; 

-- name: ListArticlesFiltered :many
SELECT a.*
FROM articles a
JOIN authors u ON a.author_id = u.id
WHERE
  ($1 = '' OR (a.title ILIKE '%' || $1 || '%' OR a.body ILIKE '%' || $1 || '%'))
  AND
  ($2 = '' OR u.name ILIKE '%' || $2 || '%' )
ORDER BY a.created_at DESC
LIMIT $3 OFFSET $4;