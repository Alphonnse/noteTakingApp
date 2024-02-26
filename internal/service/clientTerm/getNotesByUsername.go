package clientterm

import (
	"context"
	"fmt"

	"github.com/Alphonnse/noteTakingApp/internal/model"
)

func (s *service) GetNotesUsername(ctx context.Context, getNotesByUsername *model.GetNotesByUsername) (*model.RequestedNotes, error) {
	requestedNotes, err := s.clientTermRepository.GetNotesUsername(ctx, getNotesByUsername)
	if err != nil {
		return nil, fmt.Errorf("Error while geting notes from DB: %s", err.Error()) 
	}
	return requestedNotes, nil
}
