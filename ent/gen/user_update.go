// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"github.com/sirupsenV2/logrus/ent/gen/predicate"
	"github.com/sirupsenV2/logrus/ent/gen/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetWords sets the "words" field.
func (uu *UserUpdate) SetWords(s string) *UserUpdate {
	uu.mutation.SetWords(s)
	return uu
}

// SetNillableWords sets the "words" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWords(s *string) *UserUpdate {
	if s != nil {
		uu.SetWords(*s)
	}
	return uu
}

// ClearWords clears the value of the "words" field.
func (uu *UserUpdate) ClearWords() *UserUpdate {
	uu.mutation.ClearWords()
	return uu
}

// SetNetwork sets the "network" field.
func (uu *UserUpdate) SetNetwork(s string) *UserUpdate {
	uu.mutation.SetNetwork(s)
	return uu
}

// SetNillableNetwork sets the "network" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNetwork(s *string) *UserUpdate {
	if s != nil {
		uu.SetNetwork(*s)
	}
	return uu
}

// ClearNetwork clears the value of the "network" field.
func (uu *UserUpdate) ClearNetwork() *UserUpdate {
	uu.mutation.ClearNetwork()
	return uu
}

// SetAddress sets the "address" field.
func (uu *UserUpdate) SetAddress(s string) *UserUpdate {
	uu.mutation.SetAddress(s)
	return uu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddress(s *string) *UserUpdate {
	if s != nil {
		uu.SetAddress(*s)
	}
	return uu
}

// ClearAddress clears the value of the "address" field.
func (uu *UserUpdate) ClearAddress() *UserUpdate {
	uu.mutation.ClearAddress()
	return uu
}

// SetPrivateKey sets the "private_key" field.
func (uu *UserUpdate) SetPrivateKey(s string) *UserUpdate {
	uu.mutation.SetPrivateKey(s)
	return uu
}

// SetNillablePrivateKey sets the "private_key" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePrivateKey(s *string) *UserUpdate {
	if s != nil {
		uu.SetPrivateKey(*s)
	}
	return uu
}

// ClearPrivateKey clears the value of the "private_key" field.
func (uu *UserUpdate) ClearPrivateKey() *UserUpdate {
	uu.mutation.ClearPrivateKey()
	return uu
}

// SetBalance sets the "balance" field.
func (uu *UserUpdate) SetBalance(d decimal.Decimal) *UserUpdate {
	uu.mutation.SetBalance(d)
	return uu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBalance(d *decimal.Decimal) *UserUpdate {
	if d != nil {
		uu.SetBalance(*d)
	}
	return uu
}

// SetBalanceUpdateTime sets the "balance_update_time" field.
func (uu *UserUpdate) SetBalanceUpdateTime(i int) *UserUpdate {
	uu.mutation.ResetBalanceUpdateTime()
	uu.mutation.SetBalanceUpdateTime(i)
	return uu
}

// SetNillableBalanceUpdateTime sets the "balance_update_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBalanceUpdateTime(i *int) *UserUpdate {
	if i != nil {
		uu.SetBalanceUpdateTime(*i)
	}
	return uu
}

// AddBalanceUpdateTime adds i to the "balance_update_time" field.
func (uu *UserUpdate) AddBalanceUpdateTime(i int) *UserUpdate {
	uu.mutation.AddBalanceUpdateTime(i)
	return uu
}

// ClearBalanceUpdateTime clears the value of the "balance_update_time" field.
func (uu *UserUpdate) ClearBalanceUpdateTime() *UserUpdate {
	uu.mutation.ClearBalanceUpdateTime()
	return uu
}

// SetTokenInfo sets the "token_info" field.
func (uu *UserUpdate) SetTokenInfo(m map[string]interface{}) *UserUpdate {
	uu.mutation.SetTokenInfo(m)
	return uu
}

// ClearTokenInfo clears the value of the "token_info" field.
func (uu *UserUpdate) ClearTokenInfo() *UserUpdate {
	uu.mutation.ClearTokenInfo()
	return uu
}

// SetCreateTime sets the "create_time" field.
func (uu *UserUpdate) SetCreateTime(i int) *UserUpdate {
	uu.mutation.ResetCreateTime()
	uu.mutation.SetCreateTime(i)
	return uu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreateTime(i *int) *UserUpdate {
	if i != nil {
		uu.SetCreateTime(*i)
	}
	return uu
}

// AddCreateTime adds i to the "create_time" field.
func (uu *UserUpdate) AddCreateTime(i int) *UserUpdate {
	uu.mutation.AddCreateTime(i)
	return uu
}

// ClearCreateTime clears the value of the "create_time" field.
func (uu *UserUpdate) ClearCreateTime() *UserUpdate {
	uu.mutation.ClearCreateTime()
	return uu
}

// SetCreateDate sets the "create_date" field.
func (uu *UserUpdate) SetCreateDate(t time.Time) *UserUpdate {
	uu.mutation.SetCreateDate(t)
	return uu
}

// SetNillableCreateDate sets the "create_date" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreateDate(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreateDate(*t)
	}
	return uu
}

// ClearCreateDate clears the value of the "create_date" field.
func (uu *UserUpdate) ClearCreateDate() *UserUpdate {
	uu.mutation.ClearCreateDate()
	return uu
}

// SetIsTransfer sets the "is_transfer" field.
func (uu *UserUpdate) SetIsTransfer(i int8) *UserUpdate {
	uu.mutation.ResetIsTransfer()
	uu.mutation.SetIsTransfer(i)
	return uu
}

// SetNillableIsTransfer sets the "is_transfer" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsTransfer(i *int8) *UserUpdate {
	if i != nil {
		uu.SetIsTransfer(*i)
	}
	return uu
}

// AddIsTransfer adds i to the "is_transfer" field.
func (uu *UserUpdate) AddIsTransfer(i int8) *UserUpdate {
	uu.mutation.AddIsTransfer(i)
	return uu
}

// ClearIsTransfer clears the value of the "is_transfer" field.
func (uu *UserUpdate) ClearIsTransfer() *UserUpdate {
	uu.mutation.ClearIsTransfer()
	return uu
}

// SetTotalTokenValue sets the "total_token_value" field.
func (uu *UserUpdate) SetTotalTokenValue(d decimal.Decimal) *UserUpdate {
	uu.mutation.SetTotalTokenValue(d)
	return uu
}

// SetNillableTotalTokenValue sets the "total_token_value" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTotalTokenValue(d *decimal.Decimal) *UserUpdate {
	if d != nil {
		uu.SetTotalTokenValue(*d)
	}
	return uu
}

// SetTrxMode sets the "trx_mode" field.
func (uu *UserUpdate) SetTrxMode(um user.TrxMode) *UserUpdate {
	uu.mutation.SetTrxMode(um)
	return uu
}

// SetNillableTrxMode sets the "trx_mode" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTrxMode(um *user.TrxMode) *UserUpdate {
	if um != nil {
		uu.SetTrxMode(*um)
	}
	return uu
}

// ClearTrxMode clears the value of the "trx_mode" field.
func (uu *UserUpdate) ClearTrxMode() *UserUpdate {
	uu.mutation.ClearTrxMode()
	return uu
}

// SetTrxAddrType sets the "trx_addr_type" field.
func (uu *UserUpdate) SetTrxAddrType(uat user.TrxAddrType) *UserUpdate {
	uu.mutation.SetTrxAddrType(uat)
	return uu
}

// SetNillableTrxAddrType sets the "trx_addr_type" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTrxAddrType(uat *user.TrxAddrType) *UserUpdate {
	if uat != nil {
		uu.SetTrxAddrType(*uat)
	}
	return uu
}

// ClearTrxAddrType clears the value of the "trx_addr_type" field.
func (uu *UserUpdate) ClearTrxAddrType() *UserUpdate {
	uu.mutation.ClearTrxAddrType()
	return uu
}

// SetTrxPrivAddr sets the "trx_priv_addr" field.
func (uu *UserUpdate) SetTrxPrivAddr(s string) *UserUpdate {
	uu.mutation.SetTrxPrivAddr(s)
	return uu
}

// SetNillableTrxPrivAddr sets the "trx_priv_addr" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTrxPrivAddr(s *string) *UserUpdate {
	if s != nil {
		uu.SetTrxPrivAddr(*s)
	}
	return uu
}

// ClearTrxPrivAddr clears the value of the "trx_priv_addr" field.
func (uu *UserUpdate) ClearTrxPrivAddr() *UserUpdate {
	uu.mutation.ClearTrxPrivAddr()
	return uu
}

// SetTrxPrivPkey sets the "trx_priv_pkey" field.
func (uu *UserUpdate) SetTrxPrivPkey(s string) *UserUpdate {
	uu.mutation.SetTrxPrivPkey(s)
	return uu
}

// SetNillableTrxPrivPkey sets the "trx_priv_pkey" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTrxPrivPkey(s *string) *UserUpdate {
	if s != nil {
		uu.SetTrxPrivPkey(*s)
	}
	return uu
}

// ClearTrxPrivPkey clears the value of the "trx_priv_pkey" field.
func (uu *UserUpdate) ClearTrxPrivPkey() *UserUpdate {
	uu.mutation.ClearTrxPrivPkey()
	return uu
}

// SetAesType sets the "aes_type" field.
func (uu *UserUpdate) SetAesType(i int8) *UserUpdate {
	uu.mutation.ResetAesType()
	uu.mutation.SetAesType(i)
	return uu
}

// SetNillableAesType sets the "aes_type" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAesType(i *int8) *UserUpdate {
	if i != nil {
		uu.SetAesType(*i)
	}
	return uu
}

// AddAesType adds i to the "aes_type" field.
func (uu *UserUpdate) AddAesType(i int8) *UserUpdate {
	uu.mutation.AddAesType(i)
	return uu
}

// ClearAesType clears the value of the "aes_type" field.
func (uu *UserUpdate) ClearAesType() *UserUpdate {
	uu.mutation.ClearAesType()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.TrxMode(); ok {
		if err := user.TrxModeValidator(v); err != nil {
			return &ValidationError{Name: "trx_mode", err: fmt.Errorf(`gen: validator failed for field "User.trx_mode": %w`, err)}
		}
	}
	if v, ok := uu.mutation.TrxAddrType(); ok {
		if err := user.TrxAddrTypeValidator(v); err != nil {
			return &ValidationError{Name: "trx_addr_type", err: fmt.Errorf(`gen: validator failed for field "User.trx_addr_type": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Words(); ok {
		_spec.SetField(user.FieldWords, field.TypeString, value)
	}
	if uu.mutation.WordsCleared() {
		_spec.ClearField(user.FieldWords, field.TypeString)
	}
	if value, ok := uu.mutation.Network(); ok {
		_spec.SetField(user.FieldNetwork, field.TypeString, value)
	}
	if uu.mutation.NetworkCleared() {
		_spec.ClearField(user.FieldNetwork, field.TypeString)
	}
	if value, ok := uu.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if uu.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := uu.mutation.PrivateKey(); ok {
		_spec.SetField(user.FieldPrivateKey, field.TypeString, value)
	}
	if uu.mutation.PrivateKeyCleared() {
		_spec.ClearField(user.FieldPrivateKey, field.TypeString)
	}
	if value, ok := uu.mutation.Balance(); ok {
		_spec.SetField(user.FieldBalance, field.TypeOther, value)
	}
	if value, ok := uu.mutation.BalanceUpdateTime(); ok {
		_spec.SetField(user.FieldBalanceUpdateTime, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedBalanceUpdateTime(); ok {
		_spec.AddField(user.FieldBalanceUpdateTime, field.TypeInt, value)
	}
	if uu.mutation.BalanceUpdateTimeCleared() {
		_spec.ClearField(user.FieldBalanceUpdateTime, field.TypeInt)
	}
	if value, ok := uu.mutation.TokenInfo(); ok {
		_spec.SetField(user.FieldTokenInfo, field.TypeJSON, value)
	}
	if uu.mutation.TokenInfoCleared() {
		_spec.ClearField(user.FieldTokenInfo, field.TypeJSON)
	}
	if value, ok := uu.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedCreateTime(); ok {
		_spec.AddField(user.FieldCreateTime, field.TypeInt, value)
	}
	if uu.mutation.CreateTimeCleared() {
		_spec.ClearField(user.FieldCreateTime, field.TypeInt)
	}
	if value, ok := uu.mutation.CreateDate(); ok {
		_spec.SetField(user.FieldCreateDate, field.TypeTime, value)
	}
	if uu.mutation.CreateDateCleared() {
		_spec.ClearField(user.FieldCreateDate, field.TypeTime)
	}
	if value, ok := uu.mutation.IsTransfer(); ok {
		_spec.SetField(user.FieldIsTransfer, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.AddedIsTransfer(); ok {
		_spec.AddField(user.FieldIsTransfer, field.TypeInt8, value)
	}
	if uu.mutation.IsTransferCleared() {
		_spec.ClearField(user.FieldIsTransfer, field.TypeInt8)
	}
	if value, ok := uu.mutation.TotalTokenValue(); ok {
		_spec.SetField(user.FieldTotalTokenValue, field.TypeOther, value)
	}
	if value, ok := uu.mutation.TrxMode(); ok {
		_spec.SetField(user.FieldTrxMode, field.TypeEnum, value)
	}
	if uu.mutation.TrxModeCleared() {
		_spec.ClearField(user.FieldTrxMode, field.TypeEnum)
	}
	if value, ok := uu.mutation.TrxAddrType(); ok {
		_spec.SetField(user.FieldTrxAddrType, field.TypeEnum, value)
	}
	if uu.mutation.TrxAddrTypeCleared() {
		_spec.ClearField(user.FieldTrxAddrType, field.TypeEnum)
	}
	if value, ok := uu.mutation.TrxPrivAddr(); ok {
		_spec.SetField(user.FieldTrxPrivAddr, field.TypeString, value)
	}
	if uu.mutation.TrxPrivAddrCleared() {
		_spec.ClearField(user.FieldTrxPrivAddr, field.TypeString)
	}
	if value, ok := uu.mutation.TrxPrivPkey(); ok {
		_spec.SetField(user.FieldTrxPrivPkey, field.TypeString, value)
	}
	if uu.mutation.TrxPrivPkeyCleared() {
		_spec.ClearField(user.FieldTrxPrivPkey, field.TypeString)
	}
	if value, ok := uu.mutation.AesType(); ok {
		_spec.SetField(user.FieldAesType, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.AddedAesType(); ok {
		_spec.AddField(user.FieldAesType, field.TypeInt8, value)
	}
	if uu.mutation.AesTypeCleared() {
		_spec.ClearField(user.FieldAesType, field.TypeInt8)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetWords sets the "words" field.
func (uuo *UserUpdateOne) SetWords(s string) *UserUpdateOne {
	uuo.mutation.SetWords(s)
	return uuo
}

// SetNillableWords sets the "words" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWords(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetWords(*s)
	}
	return uuo
}

// ClearWords clears the value of the "words" field.
func (uuo *UserUpdateOne) ClearWords() *UserUpdateOne {
	uuo.mutation.ClearWords()
	return uuo
}

// SetNetwork sets the "network" field.
func (uuo *UserUpdateOne) SetNetwork(s string) *UserUpdateOne {
	uuo.mutation.SetNetwork(s)
	return uuo
}

// SetNillableNetwork sets the "network" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNetwork(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNetwork(*s)
	}
	return uuo
}

// ClearNetwork clears the value of the "network" field.
func (uuo *UserUpdateOne) ClearNetwork() *UserUpdateOne {
	uuo.mutation.ClearNetwork()
	return uuo
}

// SetAddress sets the "address" field.
func (uuo *UserUpdateOne) SetAddress(s string) *UserUpdateOne {
	uuo.mutation.SetAddress(s)
	return uuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddress(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAddress(*s)
	}
	return uuo
}

// ClearAddress clears the value of the "address" field.
func (uuo *UserUpdateOne) ClearAddress() *UserUpdateOne {
	uuo.mutation.ClearAddress()
	return uuo
}

// SetPrivateKey sets the "private_key" field.
func (uuo *UserUpdateOne) SetPrivateKey(s string) *UserUpdateOne {
	uuo.mutation.SetPrivateKey(s)
	return uuo
}

// SetNillablePrivateKey sets the "private_key" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePrivateKey(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPrivateKey(*s)
	}
	return uuo
}

// ClearPrivateKey clears the value of the "private_key" field.
func (uuo *UserUpdateOne) ClearPrivateKey() *UserUpdateOne {
	uuo.mutation.ClearPrivateKey()
	return uuo
}

// SetBalance sets the "balance" field.
func (uuo *UserUpdateOne) SetBalance(d decimal.Decimal) *UserUpdateOne {
	uuo.mutation.SetBalance(d)
	return uuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBalance(d *decimal.Decimal) *UserUpdateOne {
	if d != nil {
		uuo.SetBalance(*d)
	}
	return uuo
}

// SetBalanceUpdateTime sets the "balance_update_time" field.
func (uuo *UserUpdateOne) SetBalanceUpdateTime(i int) *UserUpdateOne {
	uuo.mutation.ResetBalanceUpdateTime()
	uuo.mutation.SetBalanceUpdateTime(i)
	return uuo
}

// SetNillableBalanceUpdateTime sets the "balance_update_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBalanceUpdateTime(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetBalanceUpdateTime(*i)
	}
	return uuo
}

// AddBalanceUpdateTime adds i to the "balance_update_time" field.
func (uuo *UserUpdateOne) AddBalanceUpdateTime(i int) *UserUpdateOne {
	uuo.mutation.AddBalanceUpdateTime(i)
	return uuo
}

// ClearBalanceUpdateTime clears the value of the "balance_update_time" field.
func (uuo *UserUpdateOne) ClearBalanceUpdateTime() *UserUpdateOne {
	uuo.mutation.ClearBalanceUpdateTime()
	return uuo
}

// SetTokenInfo sets the "token_info" field.
func (uuo *UserUpdateOne) SetTokenInfo(m map[string]interface{}) *UserUpdateOne {
	uuo.mutation.SetTokenInfo(m)
	return uuo
}

// ClearTokenInfo clears the value of the "token_info" field.
func (uuo *UserUpdateOne) ClearTokenInfo() *UserUpdateOne {
	uuo.mutation.ClearTokenInfo()
	return uuo
}

// SetCreateTime sets the "create_time" field.
func (uuo *UserUpdateOne) SetCreateTime(i int) *UserUpdateOne {
	uuo.mutation.ResetCreateTime()
	uuo.mutation.SetCreateTime(i)
	return uuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreateTime(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetCreateTime(*i)
	}
	return uuo
}

// AddCreateTime adds i to the "create_time" field.
func (uuo *UserUpdateOne) AddCreateTime(i int) *UserUpdateOne {
	uuo.mutation.AddCreateTime(i)
	return uuo
}

// ClearCreateTime clears the value of the "create_time" field.
func (uuo *UserUpdateOne) ClearCreateTime() *UserUpdateOne {
	uuo.mutation.ClearCreateTime()
	return uuo
}

// SetCreateDate sets the "create_date" field.
func (uuo *UserUpdateOne) SetCreateDate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreateDate(t)
	return uuo
}

// SetNillableCreateDate sets the "create_date" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreateDate(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreateDate(*t)
	}
	return uuo
}

// ClearCreateDate clears the value of the "create_date" field.
func (uuo *UserUpdateOne) ClearCreateDate() *UserUpdateOne {
	uuo.mutation.ClearCreateDate()
	return uuo
}

// SetIsTransfer sets the "is_transfer" field.
func (uuo *UserUpdateOne) SetIsTransfer(i int8) *UserUpdateOne {
	uuo.mutation.ResetIsTransfer()
	uuo.mutation.SetIsTransfer(i)
	return uuo
}

// SetNillableIsTransfer sets the "is_transfer" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsTransfer(i *int8) *UserUpdateOne {
	if i != nil {
		uuo.SetIsTransfer(*i)
	}
	return uuo
}

// AddIsTransfer adds i to the "is_transfer" field.
func (uuo *UserUpdateOne) AddIsTransfer(i int8) *UserUpdateOne {
	uuo.mutation.AddIsTransfer(i)
	return uuo
}

// ClearIsTransfer clears the value of the "is_transfer" field.
func (uuo *UserUpdateOne) ClearIsTransfer() *UserUpdateOne {
	uuo.mutation.ClearIsTransfer()
	return uuo
}

// SetTotalTokenValue sets the "total_token_value" field.
func (uuo *UserUpdateOne) SetTotalTokenValue(d decimal.Decimal) *UserUpdateOne {
	uuo.mutation.SetTotalTokenValue(d)
	return uuo
}

// SetNillableTotalTokenValue sets the "total_token_value" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTotalTokenValue(d *decimal.Decimal) *UserUpdateOne {
	if d != nil {
		uuo.SetTotalTokenValue(*d)
	}
	return uuo
}

// SetTrxMode sets the "trx_mode" field.
func (uuo *UserUpdateOne) SetTrxMode(um user.TrxMode) *UserUpdateOne {
	uuo.mutation.SetTrxMode(um)
	return uuo
}

// SetNillableTrxMode sets the "trx_mode" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTrxMode(um *user.TrxMode) *UserUpdateOne {
	if um != nil {
		uuo.SetTrxMode(*um)
	}
	return uuo
}

// ClearTrxMode clears the value of the "trx_mode" field.
func (uuo *UserUpdateOne) ClearTrxMode() *UserUpdateOne {
	uuo.mutation.ClearTrxMode()
	return uuo
}

// SetTrxAddrType sets the "trx_addr_type" field.
func (uuo *UserUpdateOne) SetTrxAddrType(uat user.TrxAddrType) *UserUpdateOne {
	uuo.mutation.SetTrxAddrType(uat)
	return uuo
}

// SetNillableTrxAddrType sets the "trx_addr_type" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTrxAddrType(uat *user.TrxAddrType) *UserUpdateOne {
	if uat != nil {
		uuo.SetTrxAddrType(*uat)
	}
	return uuo
}

// ClearTrxAddrType clears the value of the "trx_addr_type" field.
func (uuo *UserUpdateOne) ClearTrxAddrType() *UserUpdateOne {
	uuo.mutation.ClearTrxAddrType()
	return uuo
}

// SetTrxPrivAddr sets the "trx_priv_addr" field.
func (uuo *UserUpdateOne) SetTrxPrivAddr(s string) *UserUpdateOne {
	uuo.mutation.SetTrxPrivAddr(s)
	return uuo
}

// SetNillableTrxPrivAddr sets the "trx_priv_addr" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTrxPrivAddr(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTrxPrivAddr(*s)
	}
	return uuo
}

// ClearTrxPrivAddr clears the value of the "trx_priv_addr" field.
func (uuo *UserUpdateOne) ClearTrxPrivAddr() *UserUpdateOne {
	uuo.mutation.ClearTrxPrivAddr()
	return uuo
}

// SetTrxPrivPkey sets the "trx_priv_pkey" field.
func (uuo *UserUpdateOne) SetTrxPrivPkey(s string) *UserUpdateOne {
	uuo.mutation.SetTrxPrivPkey(s)
	return uuo
}

// SetNillableTrxPrivPkey sets the "trx_priv_pkey" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTrxPrivPkey(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTrxPrivPkey(*s)
	}
	return uuo
}

// ClearTrxPrivPkey clears the value of the "trx_priv_pkey" field.
func (uuo *UserUpdateOne) ClearTrxPrivPkey() *UserUpdateOne {
	uuo.mutation.ClearTrxPrivPkey()
	return uuo
}

// SetAesType sets the "aes_type" field.
func (uuo *UserUpdateOne) SetAesType(i int8) *UserUpdateOne {
	uuo.mutation.ResetAesType()
	uuo.mutation.SetAesType(i)
	return uuo
}

// SetNillableAesType sets the "aes_type" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAesType(i *int8) *UserUpdateOne {
	if i != nil {
		uuo.SetAesType(*i)
	}
	return uuo
}

// AddAesType adds i to the "aes_type" field.
func (uuo *UserUpdateOne) AddAesType(i int8) *UserUpdateOne {
	uuo.mutation.AddAesType(i)
	return uuo
}

// ClearAesType clears the value of the "aes_type" field.
func (uuo *UserUpdateOne) ClearAesType() *UserUpdateOne {
	uuo.mutation.ClearAesType()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.TrxMode(); ok {
		if err := user.TrxModeValidator(v); err != nil {
			return &ValidationError{Name: "trx_mode", err: fmt.Errorf(`gen: validator failed for field "User.trx_mode": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.TrxAddrType(); ok {
		if err := user.TrxAddrTypeValidator(v); err != nil {
			return &ValidationError{Name: "trx_addr_type", err: fmt.Errorf(`gen: validator failed for field "User.trx_addr_type": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Words(); ok {
		_spec.SetField(user.FieldWords, field.TypeString, value)
	}
	if uuo.mutation.WordsCleared() {
		_spec.ClearField(user.FieldWords, field.TypeString)
	}
	if value, ok := uuo.mutation.Network(); ok {
		_spec.SetField(user.FieldNetwork, field.TypeString, value)
	}
	if uuo.mutation.NetworkCleared() {
		_spec.ClearField(user.FieldNetwork, field.TypeString)
	}
	if value, ok := uuo.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if uuo.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := uuo.mutation.PrivateKey(); ok {
		_spec.SetField(user.FieldPrivateKey, field.TypeString, value)
	}
	if uuo.mutation.PrivateKeyCleared() {
		_spec.ClearField(user.FieldPrivateKey, field.TypeString)
	}
	if value, ok := uuo.mutation.Balance(); ok {
		_spec.SetField(user.FieldBalance, field.TypeOther, value)
	}
	if value, ok := uuo.mutation.BalanceUpdateTime(); ok {
		_spec.SetField(user.FieldBalanceUpdateTime, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedBalanceUpdateTime(); ok {
		_spec.AddField(user.FieldBalanceUpdateTime, field.TypeInt, value)
	}
	if uuo.mutation.BalanceUpdateTimeCleared() {
		_spec.ClearField(user.FieldBalanceUpdateTime, field.TypeInt)
	}
	if value, ok := uuo.mutation.TokenInfo(); ok {
		_spec.SetField(user.FieldTokenInfo, field.TypeJSON, value)
	}
	if uuo.mutation.TokenInfoCleared() {
		_spec.ClearField(user.FieldTokenInfo, field.TypeJSON)
	}
	if value, ok := uuo.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedCreateTime(); ok {
		_spec.AddField(user.FieldCreateTime, field.TypeInt, value)
	}
	if uuo.mutation.CreateTimeCleared() {
		_spec.ClearField(user.FieldCreateTime, field.TypeInt)
	}
	if value, ok := uuo.mutation.CreateDate(); ok {
		_spec.SetField(user.FieldCreateDate, field.TypeTime, value)
	}
	if uuo.mutation.CreateDateCleared() {
		_spec.ClearField(user.FieldCreateDate, field.TypeTime)
	}
	if value, ok := uuo.mutation.IsTransfer(); ok {
		_spec.SetField(user.FieldIsTransfer, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.AddedIsTransfer(); ok {
		_spec.AddField(user.FieldIsTransfer, field.TypeInt8, value)
	}
	if uuo.mutation.IsTransferCleared() {
		_spec.ClearField(user.FieldIsTransfer, field.TypeInt8)
	}
	if value, ok := uuo.mutation.TotalTokenValue(); ok {
		_spec.SetField(user.FieldTotalTokenValue, field.TypeOther, value)
	}
	if value, ok := uuo.mutation.TrxMode(); ok {
		_spec.SetField(user.FieldTrxMode, field.TypeEnum, value)
	}
	if uuo.mutation.TrxModeCleared() {
		_spec.ClearField(user.FieldTrxMode, field.TypeEnum)
	}
	if value, ok := uuo.mutation.TrxAddrType(); ok {
		_spec.SetField(user.FieldTrxAddrType, field.TypeEnum, value)
	}
	if uuo.mutation.TrxAddrTypeCleared() {
		_spec.ClearField(user.FieldTrxAddrType, field.TypeEnum)
	}
	if value, ok := uuo.mutation.TrxPrivAddr(); ok {
		_spec.SetField(user.FieldTrxPrivAddr, field.TypeString, value)
	}
	if uuo.mutation.TrxPrivAddrCleared() {
		_spec.ClearField(user.FieldTrxPrivAddr, field.TypeString)
	}
	if value, ok := uuo.mutation.TrxPrivPkey(); ok {
		_spec.SetField(user.FieldTrxPrivPkey, field.TypeString, value)
	}
	if uuo.mutation.TrxPrivPkeyCleared() {
		_spec.ClearField(user.FieldTrxPrivPkey, field.TypeString)
	}
	if value, ok := uuo.mutation.AesType(); ok {
		_spec.SetField(user.FieldAesType, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.AddedAesType(); ok {
		_spec.AddField(user.FieldAesType, field.TypeInt8, value)
	}
	if uuo.mutation.AesTypeCleared() {
		_spec.ClearField(user.FieldAesType, field.TypeInt8)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
