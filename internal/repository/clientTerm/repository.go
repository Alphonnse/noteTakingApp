package clientterm

import (
	"context"
	"sync"

	clientTermModel "github.com/Alphonnse/noteTakingApp/internal/model"
	"github.com/Alphonnse/noteTakingApp/internal/repository/clientTerm/converter"
	repoModel "github.com/Alphonnse/noteTakingApp/internal/repository/clientTerm/model"
	"github.com/rs/xid"
)

type repository struct {
	data map[string]*repoModel.Notepad
	mx   sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]*repoModel.Notepad),
	}
}

// this function creates the notes
func (r *repository) CreateNote(_ context.Context, noteRequest *clientTermModel.AddNoteRequest) (*clientTermModel.NoteIDReplay, error) {
	usernameRepo, noteText := converter.ToRepoCreateNoteFromService(noteRequest)
	userNotes := r.data[usernameRepo]
	if userNotes == nil {
		userNotes = &repoModel.Notepad{
			Username: usernameRepo,
			Notes:    make(map[string]string),
		}
	}
	noteID := xid.New().String()
	userNotes.Notes[noteID] = noteText
	r.data[usernameRepo] = userNotes

	return converter.ToServiceNoteIDReplayFromRepo(noteID, r.data[usernameRepo]), nil
}

// this function returns the notes by username
func (r *repository) GetNotesUsername(_ context.Context, getNoteReq *clientTermModel.GetNotesByUsername) (*clientTermModel.RequestedNotes, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	usernameRepo := converter.ToRepoGetNotesFromService(getNoteReq)

	note, ok := r.data[usernameRepo]
	if !ok {
		return nil, clientTermModel.ErrorUserNotFound
	}
	return converter.ToServiceNotesRespFromRepo(note), nil
}

// this function returns the note by ID
func (r *repository) GetNoteID(_ context.Context, getNoteReq *clientTermModel.GetNoteByID) (*clientTermModel.RequestedNote, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	usernameRepo, neededNoteID := converter.ToRepoGetNoteFromService(getNoteReq)

	note, ok := r.data[usernameRepo]
	if !ok {
		return nil, clientTermModel.ErrorUserNotFound
	}
	return converter.ToServiceNoteRespFromRepo(note, neededNoteID), nil
}
