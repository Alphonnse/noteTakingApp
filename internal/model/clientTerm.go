package model

type AddNoteRequest struct {
	Username string
	NoteText string
}

type NoteIDReplay struct {
	Username string
	NoteID   string
}

type GetNotesByUsername struct {
	Username string
}

type GetNoteByID struct {
	Username string
	NoteID   string
}

type RequestedNote struct {
	NoteID  string 
	NoteText string
}

type RequestedNotes struct {
	NoteID   []string
	NoteText []string
}
