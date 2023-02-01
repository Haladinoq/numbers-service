package numbers

import ( 
	"github.com/golobby/container/v3"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/api"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data"
)

//RunGeneralConfig general config
func RunGeneralConfig(container *container.Container) {
	data.RunConfig(container)
	core.RunConfig(container)
	api.RunConfig(container)
}
