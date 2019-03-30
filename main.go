package main

import (
	"flag"
	"fmt"
	"github.com/nrekretep/cloudpaint/services"
)

func main() {

	usernameFlag := flag.String("u", "", "Username used for Cloudfoundry login.")
	passwordFlag := flag.String("p", "", "Password used for Cloudfoundry login.")
	diagramFlag := flag.String("d", "", "Name of the diagram to paint.")
	apiFlag := flag.String("a", "", "URL of the Cloud Controller API.")
	appGUIDFlag := flag.String("appguid", "", "Single app guid.")
	flag.Parse()

	if *usernameFlag == "" || *passwordFlag == "" || *diagramFlag == "" || *apiFlag == "" {
		fmt.Println("Please use -h flag for correct usage.")
		return
	}

	config := services.Config{Usename: *usernameFlag, Password: *passwordFlag, ApiUrl: *apiFlag}

	if *diagramFlag == "single-app" {
		diagramService, _ := services.NewSingleAppDiagramService(&config)
		rawDiagram, _ := diagramService.GetRawDiagram(*appGUIDFlag)
		fmt.Println(rawDiagram)
	}

}
