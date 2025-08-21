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

	// if err := Validate.Struct(cells); err != nil {
	// 	app.badRequestResponse(w, r, err)
	// 	return
	// }

	if err := app.jsonResponse(w, http.StatusOK, cells); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
