package converter

import (
	"github.com/Alphonnse/noteTakingApp/internal/model"
	desc "github.com/Alphonnse/noteTakingApp/pkg/clientTerm"
)


func ToDescCreateNoteFromService(info *desc.AddNoteRequest) *model.AddNoteRequest {
	return &model.AddNoteRequest{
		Username: info.Username,
		NoteText: info.NoteText,
	}
}
func ToServiceCreateNoteFromDesc(info *model.NoteIDReplay) *desc.NoteIDReplay {
	return &desc.NoteIDReplay{
		Username: info.Username,
		NoteID: info.NoteID,
	}
}

func ToDescGetNoteFromService(info *desc.GetNoteByID) *model.GetNoteByID {
	return &model.GetNoteByID{
		Username: info.Username,
		NoteID: info.NoteID,
	}
}
func ToServiceGetNoteFromDesc(info *model.RequestedNote) *desc.RequestedNote {
	return &desc.RequestedNote{
		NoteID: info.NoteID,
		NoteText: info.NoteText,
	}
}

func ToDescGetNotesFromService(info *desc.GetNotesByUsername) *model.GetNotesByUsername {
	return &model.GetNotesByUsername{
		Username: info.Username,
	}
}
func ToServiceGetNotesFromDesc(info *model.RequestedNotes) *desc.RequestedNotes {
	return &desc.RequestedNotes{
		NoteID: info.NoteID,
		NoteText: info.NoteText,
	}
}
