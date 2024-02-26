package clientterm

import (
	"context"
	"fmt"

	"github.com/Alphonnse/noteTakingApp/internal/model"
)

func (s *service) GetNoteID(ctx context.Context, getNoteByID *model.GetNoteByID) (*model.RequestedNote, error) {
	requestedNote, err := s.clientTermRepository.GetNoteID(ctx, getNoteByID)
	if err != nil {
		return nil, fmt.Errorf("Error while creating note: %s", err.Error())
	}
	return requestedNote, nil
}
