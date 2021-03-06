// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  account role
 * @Author: Gong Zibin
 * @Date:   2017-12-12
 * @Last Modified by:
 * @Last Modified time:
 */

package role

import (
	"github.com/bottos-project/bottos/common"
	"github.com/bottos-project/bottos/common/types"
	"github.com/bottos-project/bottos/db"
)

//Role is a part of DDD, according to DDD, all the data we domain them in Role
type Role struct {
	Db *db.DBService
}

//RoleInterface is the interface that provided by Role
type RoleInterface interface {
	SetChainState(value *ChainState) error
	GetChainState() (*ChainState, error)
	SetCoreState(value *CoreState) error
	GetCoreState() (*CoreState, error)

	SetAccount(name string, value *Account) error
	GetAccount(name string) (*Account, error)
	IsAccountExist(name string) bool
	SetBalance(accountName string, value *Balance) error
	GetBalance(accountName string) (*Balance, error)
	SetStakedBalance(accountName string, value *StakedBalance) error
	GetStakedBalance(accountName string) (*StakedBalance, error)
	SetTransferCredit(name string, value *TransferCredit) error
	GetTransferCredit(name string, spender string) (*TransferCredit, error)
	DeleteTransferCredit(name string, spender string) error

	SetDelegate(accountName string, value *Delegate) error
	GetDelegateByAccountName(name string) (*Delegate, error)
	GetDelegateBySignKey(key string) (*Delegate, error)
	GetCandidateBySlot(slotNum uint64) (string, error)
	GetDelegateParticipationRate() uint64

	SetScheduleDelegate(value *ScheduleDelegate) error
	GetScheduleDelegate() (*ScheduleDelegate, error)

	CreateDelegateVotes() error
	SetDelegateVotes(key string, value *DelegateVotes) error
	GetAllDelegateVotes() ([]*DelegateVotes, error)

	SetBlockHistory(blockNumber uint32, blockHash common.Hash) error
	GetBlockHistory(blockNumber uint32) (*BlockHistory, error)
	AddTransactionHistory(txHash common.Hash, blockHash common.Hash) error
	GetTransactionHistory(txHash common.Hash) (common.Hash, error)
	SetTransactionExpiration(txHash common.Hash, value *TransactionExpiration) error
	GetTransactionExpiration(txHash common.Hash) (*TransactionExpiration, error)

	GetSlotAtTime(current uint64) uint64
	GetSlotTime(slotNum uint64) uint64

	ElectNextTermDelegates() []string

	ApplyPersistance(block *types.Block) error
}

//NewRole is creating new role
func NewRole(db *db.DBService) RoleInterface {
	r := &Role{Db: db}

	r.initRole()

	return r
}

//SetChainState is setting chain state
func (r *Role) SetChainState(value *ChainState) error {
	return SetChainStateRole(r.Db, value)
}

//GetChainState is geting chain state
func (r *Role) GetChainState() (*ChainState, error) {
	return GetChainStateRole(r.Db)
}

//SetCoreState is setting core state
func (r *Role) SetCoreState(value *CoreState) error {
	return SetCoreStateRole(r.Db, value)
}

//GetCoreState is getting core state
func (r *Role) GetCoreState() (*CoreState, error) {
	return GetCoreStateRole(r.Db)
}

//SetAccount is setting account
func (r *Role) SetAccount(name string, value *Account) error {
	return SetAccountRole(r.Db, name, value)
}

//GetAccount is getting account
func (r *Role) GetAccount(name string) (*Account, error) {
	return GetAccountRole(r.Db, name)
}

//IsAccountExist is checking account if it is exist
func (r *Role) IsAccountExist(name string) bool {
	_, err := GetAccountRole(r.Db, name)
	if err != nil {
		return false
	}
	return true
}

//SetBalance is setting balance
func (r *Role) SetBalance(accountName string, value *Balance) error {
	return SetBalanceRole(r.Db, accountName, value)
}

//GetBalance is getting balance
func (r *Role) GetBalance(accountName string) (*Balance, error) {
	return GetBalanceRole(r.Db, accountName)
}

//SetStakedBalance is setting staked balance
func (r *Role) SetStakedBalance(accountName string, value *StakedBalance) error {
	return SetStakedBalanceRole(r.Db, accountName, value)
}

//GetStakedBalance is getting staked balance
func (r *Role) GetStakedBalance(accountName string) (*StakedBalance, error) {
	return GetStakedBalanceRoleByName(r.Db, accountName)
}

//SetTransferCredit is setting transfer credit
func (r *Role) SetTransferCredit(name string, value *TransferCredit) error {
	return SetTransferCreditRole(r.Db, name, value)
}

//GetTransferCredit is getting transfer credit
func (r *Role) GetTransferCredit(name string, spender string) (*TransferCredit, error) {
	return GetTransferCreditRole(r.Db, name, spender)
}

//DeleteTransferCredit is deleting transfer credit
func (r *Role) DeleteTransferCredit(name string, spender string) error {
	return DeleteTransferCreditRole(r.Db, name, spender)
}

//SetDelegate is setting delegate
func (r *Role) SetDelegate(accountName string, value *Delegate) error {
	return SetDelegateRole(r.Db, accountName, value)
}

//GetDelegateByAccountName is getting delegate by account name.
func (r *Role) GetDelegateByAccountName(name string) (*Delegate, error) {
	return GetDelegateRoleByAccountName(r.Db, name)
}

//GetDelegateBySignKey is getting delegate by sign key
func (r *Role) GetDelegateBySignKey(key string) (*Delegate, error) {
	return GetDelegateRoleBySignKey(r.Db, key)
}

//GetDelegateParticipationRate is getting delegate participationRate
func (r *Role) GetDelegateParticipationRate() uint64 {
	rate, err := GetChainStateRole(r.Db)
	if err != nil {
		return 0
	}

	return 10000 * rate.RecentSlotFilled / 64
}

//SetScheduleDelegate is setting schedule delegate
func (r *Role) SetScheduleDelegate(value *ScheduleDelegate) error {
	return SetScheduleDelegateRole(r.Db, value)
}

//GetScheduleDelegate is getting schedule delegate
func (r *Role) GetScheduleDelegate() (*ScheduleDelegate, error) {
	return GetScheduleDelegateRole(r.Db)
}

//CreateDelegateVotes is creating delegate votes
func (r *Role) CreateDelegateVotes() error {
	return CreateDelegateVotesRole(r.Db)
}

//SetDelegateVotes is setting delegate votes
func (r *Role) SetDelegateVotes(key string, value *DelegateVotes) error {
	return SetDelegateVotesRole(r.Db, key, value)
}

//GetAllDelegateVotes is gettting all delegate votes
func (r *Role) GetAllDelegateVotes() ([]*DelegateVotes, error) {
	return GetAllDelegateVotesRole(r.Db)
}

//GetCandidateBySlot is getting candidate by slot
func (r *Role) GetCandidateBySlot(slotNum uint64) (string, error) {
	return GetCandidateBySlot(r.Db, slotNum)
}

//SetBlockHistory is setting block history
func (r *Role) SetBlockHistory(blockNumber uint32, blockHash common.Hash) error {
	return SetBlockHistoryRole(r.Db, blockNumber, blockHash)
}

//GetBlockHistory is getting block history
func (r *Role) GetBlockHistory(blockNumber uint32) (*BlockHistory, error) {
	return GetBlockHistoryRole(r.Db, blockNumber)
}

//AddTransactionHistory is to add transaction history
func (r *Role) AddTransactionHistory(txHash common.Hash, blockHash common.Hash) error {
	return AddTransactionHistoryRole(r.Db, txHash, blockHash)
}

//GetTransactionHistory is to get transaction history
func (r *Role) GetTransactionHistory(txHash common.Hash) (common.Hash, error) {
	return GetTransactionHistoryRole(r.Db, txHash)
}

//SetTransactionExpiration is to set transaction history
func (r *Role) SetTransactionExpiration(txHash common.Hash, value *TransactionExpiration) error {
	return SetTransactionExpirationRole(r.Db, txHash, value)
}

//GetTransactionExpiration is to get transaction expiration
func (r *Role) GetTransactionExpiration(txHash common.Hash) (*TransactionExpiration, error) {
	return GetTransactionExpirationRoleByHash(r.Db, txHash)
}

//ElectNextTermDelegates is to elect next term delegates
func (r *Role) ElectNextTermDelegates() []string {
	return ElectNextTermDelegatesRole(r.Db)
}

//ApplyPersistance is to apply persistence blocks to option db
func (r *Role) ApplyPersistance(block *types.Block) error {
	return ApplyPersistanceRole(r, r.Db, block)
}

//initRole is to init role
func (r *Role) initRole() {
	CreateChainStateRole(r.Db)
	CreateCoreStateRole(r.Db)

	CreateAccountRole(r.Db)
	CreateBalanceRole(r.Db)
	CreateStakedBalanceRole(r.Db)
	CreateTransferCreditRole(r.Db)

	CreateDelegateRole(r.Db)
	CreateDelegateVotesRole(r.Db)

	CreateBlockHistoryRole(r.Db)
	CreateTransactionHistoryObjectRole(r.Db)
	CreateTransactionExpirationRole(r.Db)
}
