package clientterm

import (
	"context"

	"github.com/Alphonnse/noteTakingApp/internal/converter"
	desc "github.com/Alphonnse/noteTakingApp/pkg/clientTerm"
)

func (i *Implemetation) GetNoteID(ctx context.Context, req *desc.GetNoteByID) (*desc.RequestedNote, error) {
	resp, err := i.clientTermService.GetNoteID(ctx, converter.ToDescGetNoteFromService(req))
	if err != nil {
		return nil, err
	}
	return converter.ToServiceGetNoteFromDesc(resp), nil
}
