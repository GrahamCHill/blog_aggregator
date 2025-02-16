-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
           $1,
           $2,
           $3,
           $4
       )
    RETURNING * ;

-- name: GetUser :one
SELECT *
FROM users
WHERE name = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT name
FROM users;

-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES ($1, $2, $3, $4)
    RETURNING * ;


-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name AS user_name
FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
INSERT INTO feed_follows (user_id, feed_id)
VALUES ($1, $2)
    RETURNING *
    )
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
         INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
         INNER JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows
         INNER JOIN feeds ON feed_follows.feed_id = feeds.id
         INNER JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollowByUserAndURL :exec
DELETE FROM feed_follows f
WHERE f.user_id = $1
AND f.feed_id = (SELECT id FROM feeds WHERE url = $2);

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT id, url
FROM feeds
ORDER BY last_fetched_at NULLS FIRST
    LIMIT 1;
