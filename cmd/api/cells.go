package main

import (
	"net/http"

	"github.com/driveTest-Ericsson/backend/internal/store"
)

func (app *application) getCellsHandler(w http.ResponseWriter, r *http.Request) {
	fq := &store.PaginatedCellQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	cells, err := app.store.Cells.GetCells(ctx, fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, cells); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
