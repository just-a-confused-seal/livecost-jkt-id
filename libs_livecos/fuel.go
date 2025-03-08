package libs_livecos

import "net/http"
import "io/ioutil"
//import "encoding/json"
import "fmt"

func Fuel_pertamina(){
	fmt.Println("[FETCHING MYPERTAMINA LOCATION - LIVECOS-ID]")

	var baseURL = "https://mypertamina.id/api/spbu/get?city=Jakarta%20Barat%20(Kota)"
	var client = &http.Client{}
//	var response_obj map[string] any

	request, _ := http.NewRequest("GET", baseURL, nil)

	response, _ := client.Do(request)

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}
