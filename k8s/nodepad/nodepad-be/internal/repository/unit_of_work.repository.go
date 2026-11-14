package repository

import (
	"context"
	"errors"
	"fmt"
	"nodepad-be/ent"
)

type UnitOfWorkRepository interface {
	Do(ctx context.Context, fn func(ctxTx context.Context) error) (err error)
}

type txKey struct{}

type unitOfWorkRepository struct {
	entClient *ent.Client
}

func NewUnitOfWorkRepository(entClient *ent.Client) UnitOfWorkRepository {
	return &unitOfWorkRepository{
		entClient: entClient,
	}
}

func (u *unitOfWorkRepository) Do(ctx context.Context, fn func(ctxTx context.Context) error) (err error) {
	tx, err := u.entClient.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			errRb := tx.Rollback()
			if errRb != nil {
				fmt.Println("errRb", errRb)
				err = errors.New("rollback err")
			}
		} else {
			errCm := tx.Commit()
			if errCm != nil {
				fmt.Println("errCm", errCm)
				err = errors.New("Commit err")
			}
		}
	}()

	ctxTx := context.WithValue(ctx, txKey{}, tx.Client())

	err = fn(ctxTx)

	return err
}

func GetClient(ctx context.Context, client *ent.Client) *ent.Client {
	clientAny := ctx.Value(txKey{})

	clientEntTx, ok := clientAny.(*ent.Client)
	if !ok {
		return client
	}

	return clientEntTx
}
