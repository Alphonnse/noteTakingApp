package clientterm

import (
	"context"

	"github.com/Alphonnse/noteTakingApp/internal/converter"
	desc "github.com/Alphonnse/noteTakingApp/pkg/clientTerm"
)

func (i *Implemetation) GetNotesUsername(ctx context.Context, req *desc.GetNotesByUsername) (*desc.RequestedNotes, error) {
	resp, err := i.clientTermService.GetNotesUsername(ctx, converter.ToDescGetNotesFromService(req))
	if err != nil {
		return nil, err
	}
	return converter.ToServiceGetNotesFromDesc(resp), nil
}
