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
