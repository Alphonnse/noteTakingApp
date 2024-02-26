package clientterm

import (
	"context"
	"fmt"

	"github.com/Alphonnse/noteTakingApp/internal/model"
)

func (s *service) CreateNote(ctx context.Context, addNoteReq *model.AddNoteRequest) (*model.NoteIDReplay, error) {
	replay, err := s.clientTermRepository.CreateNote(ctx, addNoteReq)
	if err != nil {
		return nil, fmt.Errorf("Error while geting note from DB: %s", err.Error()) 
	}
	return replay, nil
}
