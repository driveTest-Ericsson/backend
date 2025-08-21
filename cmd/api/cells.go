package main

import (
	"net/http"

	"github.com/driveTest-Ericsson/backend/internal/store"
)

// GetUser godoc
//
//	@Summary		Fetches cells
//	@Description	Fetches all cells
//	@Tags			cell
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	store.Cell
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/cells/ [get]
func (app *application) getCellsHandler(w http.ResponseWriter, r *http.Request) {
	cq := &store.PaginatedCellQuery{
		Limit:   20,
		Offset:  0,
		OrderBy: "created_at",
	}

	cq, err := cq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(cq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	cells, err := app.store.Cells.GetCells(ctx, cq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, cells); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type PostCellPayload struct {
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	CellTech      string  `json:"cell_tech"`
	CellIdentity  int     `json:"cell_identity"`
	PLMN          string  `json:"plmn"`
	Lac           int     `json:"lac"`
	Rac           int     `json:"rac"`
	Tac           int     `json:"tac"`
	FrequencyBand string  `json:"frequency_band"`
	ARFCN         int     `json:"arfcn"`
	FrequencyMHZ  float32 `json:"frequency_mhz"`
	RXLev         int     `json:"rxlev"`
	RXQual        int     `json:"rxqual"`
	ECN0          float32 `json:"ecn0"`
	CI            float32 `json:"c_i"`
	RSCP          int     `json:"rscp"`
	RSRP          int     `json:"rsrp"`
	RSRQ          float32 `json:"rsrq"`
	SINR          float32 `json:"sinr"`
	GeneratedAt   string  `json:"generated_at"`
}

// registerUserHandler godoc
//
//	@Summary		Create Cell Data
//	@Description	Adds the new data to Cell datas
//	@Tags			cell
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		PostCellPayload	true	"cell params"
//	@Success		201		{object}	store.Cell		"Cell Created"
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/cells/ [post]
func (app *application) postCellHandler(w http.ResponseWriter, r *http.Request) {
	var payload PostCellPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	cell := &store.Cell{
		Lat:           payload.Lat,
		Long:          payload.Long,
		CellTech:      payload.CellTech,
		CellIdentity:  payload.CellIdentity,
		PLMN:          payload.PLMN,
		Lac:           payload.Lac,
		Rac:           payload.Rac,
		Tac:           payload.Tac,
		FrequencyBand: payload.FrequencyBand,
		ARFCN:         payload.ARFCN,
		FrequencyMHZ:  payload.FrequencyMHZ,
		// 2G metrics
		RXLev:  payload.RXLev,
		RXQual: payload.RXQual,
		ECN0:   payload.ECN0,
		CI:     float32(payload.CellIdentity),
		// 3G
		RSCP: payload.RSCP,
		// 4G
		RSRP:        payload.RSRP,
		RSRQ:        payload.RSRQ,
		SINR:        payload.SINR,
		GeneratedAt: payload.GeneratedAt,
	}

	ctx := r.Context()

	cell, err := app.store.Cells.Create(ctx, cell)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, cell); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
