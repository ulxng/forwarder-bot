-- name: GetChatByUserID :one
SELECT chat_id
FROM forward_config
WHERE user_id = ?;

-- name: SaveConfig :exec
INSERT INTO forward_config (user_id, chat_id)
VALUES (?, ?)
ON CONFLICT(user_id) DO UPDATE SET
    chat_id = excluded.chat_id;