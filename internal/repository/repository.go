package repository

import (
	"context"

	"github.com/Alphonnse/noteTakingApp/internal/model"
)

type ClientTermRepository interface {
	CreateNote(context.Context, *model.AddNoteRequest) (*model.NoteIDReplay, error)
	GetNotesUsername(context.Context, *model.GetNotesByUsername) (*model.RequestedNotes, error)
	GetNoteID(context.Context, *model.GetNoteByID) (*model.RequestedNote, error)
}
