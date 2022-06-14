package currency

import "github.com/scylladb/gocqlx/v2/table"

const CreateUsersBalanceTable = `CREATE TABLE IF NOT EXISTS thetan.users_balance (
		  	user_id text,
			currency_type int,
			balance bigint,
			pending_transfer UUID,
			pending_amount bigint,
			updated_at timestamp,
		  PRIMARY KEY (user_id, currency_type))`

const CreateTransfersTable = `CREATE TABLE IF NOT EXISTS thetan.transfers (
		  	transfer_id UUID,
			currency_type tinyint,
			source_id text,
			dest_id text,
			amount bigint,
			state tinyint, -- '2:new', '1:locked', '0:complete'
			client_id UUID, -- the client performing the transfer
			created_at timestamp,
		  PRIMARY KEY (transfer_id))`

var balanceMetadata = table.Metadata{
	Name: "thetan.users_balance",
	Columns: []string{"user_id",
		"currency_type",
		"balance",
		"pending_transfer",
		"pending_amount",
		"updated_at"},
	PartKey: []string{"user_id"},
	SortKey: []string{"currency_type"},
}

var transferMetadata = table.Metadata{
	Name: "thetan.transfers",
	Columns: []string{"transfer_id",
		"currency_type",
		"source_id",
		"dest_id",
		"amount",
		"state",
		"client_id",
		"created_at"},
	PartKey: []string{"transfer_id"},
	//SortKey: []string{"currency_type"},
}

// Client id has to be updated separately to let it expire
const InsertTransfer = `
INSERT INTO transfers
  (transfer_id, src_bic, src_ban, dst_bic, dst_ban, amount, state)
  VALUES (?, ?, ?, ?, ?, ?, 'new')
  IF NOT EXISTS
`

// Because of a Cassandra/Scylla bug we can't supply NULL as a parameter marker
// Always check the row exists to not accidentally add a transfer
const SetTransferClient = `
UPDATE transfers USING TTL 30
  SET client_id = ?
  WHERE transfer_id = ?
  IF amount != NULL AND client_id = NULL
`

const SET_TRANSFER_STATE = `
UPDATE transfers
  SET state = ?
  WHERE transfer_id = ?
  IF amount != NULL AND client_id = ?
`

// Always check the row exists to not accidentally add a transfer
const ClearTransferClient = `
UPDATE transfers
  SET client_id = NULL
  WHERE transfer_id = ?
  IF amount != NULL AND client_id = ?
`

const DeleteTransfer = `
DELETE FROM transfers
  WHERE transfer_id = ?
  IF client_id = ?
`

const FetchTransfer = `
SELECT source_id, dest_id, amount, state, currency_type
  FROM transfers
  WHERE transfer_id = ?
`

const FetchTransferClient = `
SELECT client_id
  FROM transfers
  WHERE transfer_id = ?
`

// FetchDeadTransfers Cassandra/Scylla don't handle IF client_id = NUll queries
// correctly. But NULLs are implicitly converted to mintimeuuids
// during comparison. Use one bug to workaround another.
// WHERE client_id < minTimeuuid('1979-08-12 21:35+0000')
const FetchDeadTransfers = `
SELECT transfer_id
  FROM transfers
  ALLOW FILTERING
`

// LockAccount Condition balance column:
// 1) To avoid accidentally inserting a new account here
// 2) To get it back (Scylla only)
const LockAccount = `
UPDATE users_balance
  SET pending_transfer = ?, pending_amount = ?
  WHERE user_id = ? AND currency_type = ?
  IF balance != NULL AND pending_amount != NULL AND pending_transfer = NULL
`

// UnlockAccount Always check the row exists in IF to not accidentally add a transfer
const UnlockAccount = `
UPDATE users_balance
  SET pending_transfer = NULL, pending_amount = 0
  WHERE user_id = ? AND currency_type = ?
  IF balance != NULL AND pending_transfer = ?
`

const FetchBalance = `
SELECT balance, pending_amount
  FROM users_balance
  WHERE user_id = ? AND currency_type = ?
`

// Always check the row exists in IF to not accidentally add a transfer
const UpdateBalance = `
UPDATE users_balance
  SET pending_amount = 0, balance = ?
  WHERE user_id = ? AND currency_type = ?
  IF balance != NULL AND pending_transfer = ?
`

const CheckBalance = `
SELECT SUM(balance) FROM users_balance
`

const PersistTotal = `
UPDATE lightest.check SET amount = ?  WHERE name = 'total'
`

const FetchTotal = `
SELECT amount FROM lightest.check WHERE name = 'total'
`
