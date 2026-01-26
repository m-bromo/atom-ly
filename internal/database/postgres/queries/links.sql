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

-- name: GetByURl :one
SELECT links.url
FROM links
WHERE url = $1;