package clientterm

import (
	"context"

	"github.com/Alphonnse/noteTakingApp/internal/converter"
	desc "github.com/Alphonnse/noteTakingApp/pkg/clientTerm"
)

func (i *Implemetation) CreateNote(ctx context.Context, req *desc.AddNoteRequest) (*desc.NoteIDReplay, error) {
	resp, err := i.clientTermService.CreateNote(ctx, converter.ToDescCreateNoteFromService(req)) 
	if err != nil {
		return nil, err
	}
	return converter.ToServiceCreateNoteFromDesc(resp), nil
}
