package main

import (
	"flag"
	"fmt"
	"github.com/nrekretep/cloudpaint/adapter/cloudfoundry"
	"github.com/nrekretep/cloudpaint/services"
)

func main() {

	usernameFlag := flag.String("u", "", "Username used for Cloudfoundry login.")
	passwordFlag := flag.String("p", "", "Password used for Cloudfoundry login.")
	flag.Parse()

	if *usernameFlag == "" || *passwordFlag == "" {
		fmt.Println("Please use -h flag for correct usage.")
		return
	}

	cc := cloudfoundry.NewCloudController("https://api.run.pivotal.io")

	err := cc.Login(*usernameFlag, *passwordFlag)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cc.GetStacks()
	err = cc.GetBuildpacks()
	err = cc.GetQuotaDefinitions()
	err = cc.GetOrganizations()
	err = cc.GetSpaces()
	err = cc.GetApps()

	createDiagramService := services.NewCreateDiagramService(cc)

	fmt.Println(createDiagramService.RenderTemplate())

}
