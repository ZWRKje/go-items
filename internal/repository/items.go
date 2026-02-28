package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"simple-api/internal/config"
	"simple-api/internal/service"

	_ "github.com/lib/pq"
)

type ItemsRepository struct {
	db *sql.DB
}

func NewItemsRepository(cfg config.Config) *ItemsRepository {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, "disable",
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Create connection err:", err)
	}

	return &ItemsRepository{db}
}

func (r *ItemsRepository) Create(ctx context.Context, data service.CreateParams) (int, error) {
	var id int
	err := r.db.QueryRowContext(ctx, "INSERT INTO items (name, description, cost) VALUES ($1, $2, $3) RETURNING ID",
		data.Name, data.Description, data.Cost).Scan(&id)

	if err != nil {
		log.Println("Create item error:", err)

		return id, err
	}

	return id, nil
}

func (r *ItemsRepository) Update(ctx context.Context, id int, params service.UpdateParams) error {
	var setStr string
	var args []interface{}

	if params.Cost != nil {
		setStr = fmt.Sprintf("query = %f", *params.Cost)
		args = append(args, params.Cost)
	}

	if params.Description != nil {
		qStr := fmt.Sprintf(", description = %s", *params.Description)
		setStr += qStr
		args = append(args, params.Description)
	}

	if params.Name != nil {
		qStr := fmt.Sprintf(", name = %s", *params.Name)
		setStr += qStr
		args = append(args, params.Name)
	}

	query := fmt.Sprintf("Update items SET %s WHERE id = $%d", setStr, id)

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println("Update item error:", err)

		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *ItemsRepository) Get(ctx context.Context, id int) (service.Item, error) {
	var item service.Item
	err := r.db.QueryRowContext(ctx, `SELECT id, name, description, cost FROM items WHERE id = $%d`, id).
		Scan(&item.Id, &item.Name, &item.Description, &item.Cost)

	if err != nil {
		log.Println("Get item error:", err)

		return item, err
	}

	return item, nil
}

func (r *ItemsRepository) Delete(ctx context.Context, id int) (int, error) {
	var delId int
	err := r.db.QueryRowContext(ctx, `DELETE FROM items WHERE id = $%d`, id).
		Scan(&delId)

	if err != nil {
		log.Println("Delete item error:", err)

		return delId, err
	}

	return delId, nil
}
