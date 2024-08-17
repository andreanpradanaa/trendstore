// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type Chart struct {
	ID        uuid.UUID        `json:"id"`
	UserID    uuid.UUID        `json:"user_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type ChartItem struct {
	ID        uuid.UUID        `json:"id"`
	CartID    uuid.UUID        `json:"cart_id"`
	ProductID uuid.UUID        `json:"product_id"`
	Quantity  int32            `json:"quantity"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type Order struct {
	ID          uuid.UUID        `json:"id"`
	UserID      uuid.UUID        `json:"user_id"`
	TotalAmount pgtype.Numeric   `json:"total_amount"`
	Status      string           `json:"status"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type OrderItem struct {
	ID        uuid.UUID        `json:"id"`
	OrderID   uuid.UUID        `json:"order_id"`
	ProductID uuid.UUID        `json:"product_id"`
	Quantity  int32            `json:"quantity"`
	Price     pgtype.Numeric   `json:"price"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type Product struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       pgtype.Numeric   `json:"price"`
	Stock       int32            `json:"stock"`
	CategoryID  uuid.UUID        `json:"category_id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type User struct {
	ID             uuid.UUID        `json:"id"`
	Name           string           `json:"name"`
	Email          string           `json:"email"`
	HashedPassword string           `json:"hashed_password"`
	Address        pgtype.Text      `json:"address"`
	Phone          string           `json:"phone"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
}
