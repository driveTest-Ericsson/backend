package db

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"time"

	"github.com/driveTest-Ericsson/backend/internal/store"
)

// define array of datas here, for randomness

func Seed(store store.Storage, db *sql.DB) {

	ctx := context.Background()

	cells := generateCells(200)

	for _, cell := range cells {
		if err := store.Cells.Create(ctx, cell); err != nil {
			log.Println("Error creating cell:", err)
			return
		}
	}

	log.Println("Seeding Completed")
}

func generateCells(num int) []*store.Cell {
	// Some sample values to pick from
	cellTechs := []string{"GSM", "UMTS", "LTE", "NR"}
	plmns := []string{"00101", "00202", "310260", "40445"} // MCC+MNC
	bands := []string{"Band 1", "Band 3", "Band 5", "Band 7", "Band 20"}

	rand.Seed(time.Now().UnixNano())
	cells := make([]*store.Cell, num)

	for i := range num {

		cells[i] = &store.Cell{
			UserID:        1,
			Lat:           35.698703 + rand.Float32()*0.5,
			Long:          51.337664 + rand.Float32()*0.5,
			CellTech:      cellTechs[rand.Intn(len(cellTechs))],
			CellIdentity:  rand.Intn(255), // bigint in DB, but small random for now
			PLMN:          plmns[rand.Intn(len(plmns))],
			Lac:           rand.Intn(65534-1) + 1, // 1–65534
			Rac:           rand.Intn(256),         // 0–255
			Tac:           rand.Intn(65534-1) + 1, // 1–65534
			FrequencyBand: bands[rand.Intn(len(bands))],
			ARFCN:         rand.Intn(2000),                    // arbitrary channel number
			FrequencyMHZ:  float32(700 + rand.Float32()*2000), // 700–2700 MHz
			// 2G metrics
			RXLev:  rand.Intn(64),                    // 0–63
			RXQual: rand.Intn(8),                     // 0–7
			ECN0:   float32(-20 + rand.Float32()*10), // -20 to -10 dB
			CI:     float32(rand.Intn(30)),           // 0–30 dB
			// 3G
			RSCP: rand.Intn(81) - 120, // -120 to -40
			// 4G
			RSRP:        rand.Intn(97) - 140,              // -140 to -44
			RSRQ:        float32(-20 + rand.Float32()*17), // -20 to -3
			SINR:        float32(rand.Intn(31)),           // 0–30 dB
			GeneratedAt: time.Now().Format("2006-01-02 15:04:05"),
		}
	}

	return cells
}
