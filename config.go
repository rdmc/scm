package main

import
//"encoding/json"

"os"

var config struct {
	cmtsUsername string
	cmtsPassword string
	cmtsAddr     string
}

/*
type Configuration struct {
	Users  []string
	Groups []string
}
*/

// becose the github repo is public, we must protect ours screts
// get cmts username and password from system enviroment
func init() {
	config.cmtsUsername = os.Getenv("SCM_CMTS_USER")
	config.cmtsPassword = os.Getenv("SCM_CMTS_PASS")
	config.cmtsAddr = os.Getenv("SCM_CMTS_ADDR")
	/*
		        file, _ := os.Open("conf.json")
			decoder := json.NewDecoder(file)
			configuration := Configuration{}
			err := decoder.Decode(&configuration)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(configuration.Users) // output: [UserA, UserB]
	*/
	//fmt.Println("Environment RDMC_CMTS_USER:" + cmtsUsername)
	//fmt.Println("Environment RDMC_CMTS_PASS:" + cmtsPassword)

}
