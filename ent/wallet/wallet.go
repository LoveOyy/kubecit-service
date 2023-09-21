// Code generated by ent, DO NOT EDIT.

package wallet

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the wallet type in the database.
	Label = "wallet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGoldLeaf holds the string denoting the gold_leaf field in the database.
	FieldGoldLeaf = "gold_leaf"
	// FieldSilverLeaf holds the string denoting the silver_leaf field in the database.
	FieldSilverLeaf = "silver_leaf"
	// FieldFrozenGoldLeaf holds the string denoting the frozen_gold_leaf field in the database.
	FieldFrozenGoldLeaf = "frozen_gold_leaf"
	// FieldFrozenSilverLeaf holds the string denoting the frozen_silver_leaf field in the database.
	FieldFrozenSilverLeaf = "frozen_silver_leaf"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the wallet in the database.
	Table = "wallets"
)

// Columns holds all SQL columns for wallet fields.
var Columns = []string{
	FieldID,
	FieldGoldLeaf,
	FieldSilverLeaf,
	FieldFrozenGoldLeaf,
	FieldFrozenSilverLeaf,
	FieldUserID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultGoldLeaf holds the default value on creation for the "gold_leaf" field.
	DefaultGoldLeaf int32
	// DefaultSilverLeaf holds the default value on creation for the "silver_leaf" field.
	DefaultSilverLeaf int32
	// DefaultFrozenGoldLeaf holds the default value on creation for the "frozen_gold_leaf" field.
	DefaultFrozenGoldLeaf int32
	// DefaultFrozenSilverLeaf holds the default value on creation for the "frozen_silver_leaf" field.
	DefaultFrozenSilverLeaf int32
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Wallet queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGoldLeaf orders the results by the gold_leaf field.
func ByGoldLeaf(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGoldLeaf, opts...).ToFunc()
}

// BySilverLeaf orders the results by the silver_leaf field.
func BySilverLeaf(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSilverLeaf, opts...).ToFunc()
}

// ByFrozenGoldLeaf orders the results by the frozen_gold_leaf field.
func ByFrozenGoldLeaf(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrozenGoldLeaf, opts...).ToFunc()
}

// ByFrozenSilverLeaf orders the results by the frozen_silver_leaf field.
func ByFrozenSilverLeaf(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrozenSilverLeaf, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
