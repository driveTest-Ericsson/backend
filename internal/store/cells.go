package store

import (
	"context"
	"database/sql"
)

type CellStore struct {
	db *sql.DB
}

type Cell struct {
	ID            int64   `json:"id"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	CellTech      string  `json:"cell_tech"`
	CellIdentity  int8    `json:"cell_identity"`
	PLMN          string  `json:"plmn"`
	Lac           int8    `json:"lac"`
	Rac           int8    `json:"rac"`
	Tac           int8    `json:"tac"`
	FrequencyBand string  `json:"frequency_band"`
	ARFCN         int8    `json:"arfcn"`
	FrequencyMHZ  float32 `json:"frequency_mhx"`
	RXLev         int8    `json:"rxlev"`
	RXQual        int8    `json:"rxqual"`
	ECN0          float32 `json:"ecn0"`
	CI            float32 `json:"c_i"`
	RSCP          int8    `json:"rscp"`
	RSRP          int8    `json:"rsrp"`
	RSRQ          float32 `json:"rsrq"`
	SINR          float32 `json:"sinr"`
	GeneratedAt   string  `json:"generated_at"`
	CreatedAt     string  `json:"created_at"`
}

func (s *CellStore) Create(context.Context, *sql.Tx, *Cell) error {
	return nil
}
func (s *CellStore) GetByID(context.Context, int64) (*Cell, error) {
	return nil, nil
}

func (s *CellStore) Delete(context.Context, int64) error {
	return nil
}
