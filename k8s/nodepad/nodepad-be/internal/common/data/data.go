package data

import (
	"context"
	"fmt"
	"log"
	"nodepad-be/ent"
	"time"

	_ "nodepad-be/ent/runtime"

	_ "github.com/lib/pq"
)

type Data struct {
	Db *ent.Client
}

// NewData .
func NewData(dsn string) (*Data, func(), error) {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("open postgres: %w", err)
	}

	cleanup := func() {
		_ = client.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.QueryContext(ctx, "SELECT 1 + 1")
	if err != nil {
		log.Fatalf("❌ [ENT] failed connection to PostgreSQL: %v", err)
	}

	fmt.Println("✅ [ENT] Connection To PostgreSQL Successfully")

	return &Data{Db: client}, cleanup, nil
}
