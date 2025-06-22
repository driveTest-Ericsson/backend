package store

import (
	"context"
	"database/sql"
	"errors"
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

func (s *CellStore) Create(ctx context.Context, cell *Cell) error {
	query := `
	INSERT INTO cell (lat, long, cell_identity, plmn, lac, rac, tac, frequency_band, arfcn, frequency_mhz, rxlev, rxqual, ec_n0, c_i, rscp, rsrp, rsrq, sinr, generated_at, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
	RETURNING id,created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		cell.Lat,
		cell.Long,
		cell.CellIdentity,
		cell.PLMN,
		cell.Lac,
		cell.Rac,
		cell.Tac,
		cell.FrequencyBand,
		cell.ARFCN,
		cell.FrequencyMHZ,
		cell.RXLev,
		cell.RXQual,
		cell.ECN0,
		cell.CI,
		cell.RSCP,
		cell.RSRP,
		cell.RSRQ,
		cell.SINR,
		cell.GeneratedAt,
	).Scan(
		&cell.ID,
		&cell.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *CellStore) GetByID(ctx context.Context, id int64) (*Cell, error) {
	query := `
	SELECT *
	FROM cells
	WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var cell Cell
	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&cell.ID,
		&cell.Lat,
		&cell.Long,
		&cell.CellIdentity,
		&cell.PLMN,
		&cell.Lac,
		&cell.Rac,
		&cell.Tac,
		&cell.FrequencyBand,
		&cell.ARFCN,
		&cell.FrequencyMHZ,
		&cell.RXLev,
		&cell.RXQual,
		&cell.ECN0,
		&cell.CI,
		&cell.RSCP,
		&cell.RSRP,
		&cell.RSRQ,
		&cell.SINR,
		&cell.GeneratedAt,
		&cell.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &cell, nil

}

func (s *CellStore) GetCells(ctx context.Context, cq *PaginatedCellQuery) (*[]Cell, error) {
	query := `
	SELECT *
	FROM cells
	ORDER BY created_at ` + cq.Sort + `
	LIMIT $1 OFFSET $2;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, cq.Limit, cq.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cells []Cell

	for rows.Next() {
		var cell Cell
		err := rows.Scan(
			&cell.ID,
			&cell.Lat,
			&cell.Long,
			&cell.CellIdentity,
			&cell.PLMN,
			&cell.Lac,
			&cell.Rac,
			&cell.Tac,
			&cell.FrequencyBand,
			&cell.ARFCN,
			&cell.FrequencyMHZ,
			&cell.RXLev,
			&cell.RXQual,
			&cell.ECN0,
			&cell.CI,
			&cell.RSCP,
			&cell.RSRP,
			&cell.RSRQ,
			&cell.SINR,
			&cell.GeneratedAt,
			&cell.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		cells = append(cells, cell)
	}
	return &cells, nil

}

func (s *CellStore) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM cells WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
