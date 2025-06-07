-- name: GetFornecedoresById :one
SELECT * FROM fornecedores_dados WHERE user_id = $1;