package mockgen_host

import (
	"log"
	"os"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_cli/controllers"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_cli/flags"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_host/configurations"
)

func Run() {
	parsed := flags.Parse(os.Args[0], os.Args[1:])

	container := configurations.ConfigureDI(parsed)
	if err := container.Invoke(func(controller *controllers.GenerateMocksController) error {
		return controller.Execute()
	}); err != nil {
		log.Fatal(err)
	}
}
