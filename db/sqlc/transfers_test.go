package db

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"bank-backend.ethan.klaric/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {

	accs, _ := testQueries.ListAccounts(context.Background(), ListAccountsParams{Limit: 5})
	l := len(accs) - 1
	randomAccIndex := rand.Intn(l)
	acc1 := accs[randomAccIndex]
	acc2 := accs[randomAccIndex+1]
	amt := float64(util.RandomMoney())

	arg := CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        amt,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, acc1.ID)
	require.Equal(t, transfer.ToAccountID, acc2.ID)
	require.Equal(t, amt, transfer.Amount)

	return transfer
}

func TestNewTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	amt := float64(util.RandomMoney())

	args := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: amt,
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, amt, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	err := testQueries.Deletetransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
