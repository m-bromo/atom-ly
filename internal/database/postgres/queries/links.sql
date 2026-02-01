-- name: Save :one
INSERT INTO links (
    url,
    created_at
) VALUES (
    $1, $2
) RETURNING id;

-- name: GetByID :one
SELECT links.url
FROM links
WHERE id = $1;

-- name: GetIDByUrl :one
SELECT links.id
FROM links
WHERE url = $1;