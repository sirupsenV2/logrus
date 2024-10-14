// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWords holds the string denoting the words field in the database.
	FieldWords = "words"
	// FieldNetwork holds the string denoting the network field in the database.
	FieldNetwork = "network"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPrivateKey holds the string denoting the private_key field in the database.
	FieldPrivateKey = "private_key"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldBalanceUpdateTime holds the string denoting the balance_update_time field in the database.
	FieldBalanceUpdateTime = "balance_update_time"
	// FieldTokenInfo holds the string denoting the token_info field in the database.
	FieldTokenInfo = "token_info"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldCreateDate holds the string denoting the create_date field in the database.
	FieldCreateDate = "create_date"
	// FieldIsTransfer holds the string denoting the is_transfer field in the database.
	FieldIsTransfer = "is_transfer"
	// FieldTotalTokenValue holds the string denoting the total_token_value field in the database.
	FieldTotalTokenValue = "total_token_value"
	// FieldTrxMode holds the string denoting the trx_mode field in the database.
	FieldTrxMode = "trx_mode"
	// FieldTrxAddrType holds the string denoting the trx_addr_type field in the database.
	FieldTrxAddrType = "trx_addr_type"
	// FieldTrxPrivAddr holds the string denoting the trx_priv_addr field in the database.
	FieldTrxPrivAddr = "trx_priv_addr"
	// FieldTrxPrivPkey holds the string denoting the trx_priv_pkey field in the database.
	FieldTrxPrivPkey = "trx_priv_pkey"
	// FieldAesType holds the string denoting the aes_type field in the database.
	FieldAesType = "aes_type"
	// Table holds the table name of the user in the database.
	Table = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldWords,
	FieldNetwork,
	FieldAddress,
	FieldPrivateKey,
	FieldBalance,
	FieldBalanceUpdateTime,
	FieldTokenInfo,
	FieldCreateTime,
	FieldCreateDate,
	FieldIsTransfer,
	FieldTotalTokenValue,
	FieldTrxMode,
	FieldTrxAddrType,
	FieldTrxPrivAddr,
	FieldTrxPrivPkey,
	FieldAesType,
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
	// DefaultWords holds the default value on creation for the "words" field.
	DefaultWords string
	// DefaultNetwork holds the default value on creation for the "network" field.
	DefaultNetwork string
	// DefaultAddress holds the default value on creation for the "address" field.
	DefaultAddress string
	// DefaultPrivateKey holds the default value on creation for the "private_key" field.
	DefaultPrivateKey string
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance decimal.Decimal
	// DefaultBalanceUpdateTime holds the default value on creation for the "balance_update_time" field.
	DefaultBalanceUpdateTime int
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime int
	// DefaultIsTransfer holds the default value on creation for the "is_transfer" field.
	DefaultIsTransfer int8
	// DefaultTotalTokenValue holds the default value on creation for the "total_token_value" field.
	DefaultTotalTokenValue decimal.Decimal
	// DefaultTrxPrivAddr holds the default value on creation for the "trx_priv_addr" field.
	DefaultTrxPrivAddr string
	// DefaultTrxPrivPkey holds the default value on creation for the "trx_priv_pkey" field.
	DefaultTrxPrivPkey string
	// DefaultAesType holds the default value on creation for the "aes_type" field.
	DefaultAesType int8
)

// TrxMode defines the type for the "trx_mode" enum field.
type TrxMode string

// TrxModeTransfer is the default value of the TrxMode enum.
const DefaultTrxMode = TrxModeTransfer

// TrxMode values.
const (
	TrxModeTransfer TrxMode = "transfer"
	TrxModeLock     TrxMode = "lock"
)

func (tm TrxMode) String() string {
	return string(tm)
}

// TrxModeValidator is a validator for the "trx_mode" field enum values. It is called by the builders before save.
func TrxModeValidator(tm TrxMode) error {
	switch tm {
	case TrxModeTransfer, TrxModeLock:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for trx_mode field: %q", tm)
	}
}

// TrxAddrType defines the type for the "trx_addr_type" enum field.
type TrxAddrType string

// TrxAddrTypeSingle is the default value of the TrxAddrType enum.
const DefaultTrxAddrType = TrxAddrTypeSingle

// TrxAddrType values.
const (
	TrxAddrTypeSingle TrxAddrType = "single"
	TrxAddrTypeMulti  TrxAddrType = "multi"
)

func (tat TrxAddrType) String() string {
	return string(tat)
}

// TrxAddrTypeValidator is a validator for the "trx_addr_type" field enum values. It is called by the builders before save.
func TrxAddrTypeValidator(tat TrxAddrType) error {
	switch tat {
	case TrxAddrTypeSingle, TrxAddrTypeMulti:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for trx_addr_type field: %q", tat)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByWords orders the results by the words field.
func ByWords(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWords, opts...).ToFunc()
}

// ByNetwork orders the results by the network field.
func ByNetwork(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNetwork, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByPrivateKey orders the results by the private_key field.
func ByPrivateKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivateKey, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}

// ByBalanceUpdateTime orders the results by the balance_update_time field.
func ByBalanceUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalanceUpdateTime, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByCreateDate orders the results by the create_date field.
func ByCreateDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateDate, opts...).ToFunc()
}

// ByIsTransfer orders the results by the is_transfer field.
func ByIsTransfer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsTransfer, opts...).ToFunc()
}

// ByTotalTokenValue orders the results by the total_token_value field.
func ByTotalTokenValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalTokenValue, opts...).ToFunc()
}

// ByTrxMode orders the results by the trx_mode field.
func ByTrxMode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrxMode, opts...).ToFunc()
}

// ByTrxAddrType orders the results by the trx_addr_type field.
func ByTrxAddrType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrxAddrType, opts...).ToFunc()
}

// ByTrxPrivAddr orders the results by the trx_priv_addr field.
func ByTrxPrivAddr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrxPrivAddr, opts...).ToFunc()
}

// ByTrxPrivPkey orders the results by the trx_priv_pkey field.
func ByTrxPrivPkey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrxPrivPkey, opts...).ToFunc()
}

// ByAesType orders the results by the aes_type field.
func ByAesType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAesType, opts...).ToFunc()
}
