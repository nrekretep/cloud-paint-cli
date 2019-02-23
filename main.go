package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nrekretep/cloudpaint/adapter/cloudfoundry"
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
	appInfo, err := cc.GetAppInfo("4d49fc09-3b8c-408c-903d-4ca0e4abd43a")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.Marshal(appInfo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("--------\n", string(data))
	}

}
