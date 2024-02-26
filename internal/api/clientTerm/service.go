package clientterm

import (
	"github.com/Alphonnse/noteTakingApp/internal/service"
	desc "github.com/Alphonnse/noteTakingApp/pkg/clientTerm"
)


type Implemetation struct {
	desc.UnimplementedClientTermServer
	clientTermService service.ClientTermService
}

func NewImplementation(clientTermService service.ClientTermService) *Implemetation {
	return &Implemetation{
		clientTermService: clientTermService,
	}
}
