package converter

import (
	"github.com/Alphonnse/noteTakingApp/internal/model"
	repoModel "github.com/Alphonnse/noteTakingApp/internal/repository/clientTerm/model"
)

func ToRepoGetNoteFromService(getNote *model.GetNoteByID) (string, string) {
	return getNote.Username, getNote.NoteID
}

func ToRepoGetNotesFromService(getNotes *model.GetNotesByUsername) string {
	return getNotes.Username
}

func ToRepoCreateNoteFromService(addNoteRequest *model.AddNoteRequest) (string, string) {
	return addNoteRequest.Username, addNoteRequest.NoteText
}

func ToServiceNoteIDReplayFromRepo(noteID string, noteIDReplayRepo *repoModel.Notepad) *model.NoteIDReplay {
	return &model.NoteIDReplay{
		Username: noteIDReplayRepo.Username,
		NoteID:   noteID,
	}
}

func ToServiceNotesRespFromRepo(notes *repoModel.Notepad) *model.RequestedNotes {
	var ids []string
	var texts []string
	for id, noteText := range notes.Notes {
		ids = append(ids, id)
		texts = append(texts, noteText)

	}
	return &model.RequestedNotes{
		NoteID:   ids,
		NoteText: texts,
	}
}

func ToServiceNoteRespFromRepo(notes *repoModel.Notepad, noteID string) *model.RequestedNote {
	return &model.RequestedNote{
		NoteID:   noteID,
		NoteText: notes.Notes[noteID],
	}
}
