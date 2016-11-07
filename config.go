package main

import (
	//"encoding/json"
	"fmt"
	"os"
)

var (
	cmtsUsername string
	cmtsPassword string
)

/*
type Configuration struct {
	Users  []string
	Groups []string
}
*/

// becose the github repo is public, we must protect ours screts
// get cmts username and password from system enviroment
func init() {
	cmtsUsername = os.Getenv("RDMC_CMTS_USER")
	cmtsPassword = os.Getenv("RDMC_CMTS_PASS")
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
	fmt.Println("Environment RDMC_CMTS_USER:" + cmtsUsername)
	fmt.Println("Environment RDMC_CMTS_PASS:" + cmtsPassword)

}
