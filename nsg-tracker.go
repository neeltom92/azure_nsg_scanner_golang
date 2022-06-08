
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var actoken string
var sg_fnl string
var sg_fnl_user string
var slack_val string
var user_condition int = 0

var input string
var duplicate_op []interface{}

var duplicate_user []interface{}

func slack_push(value_slk string) {

	actoken1 := "{\"text\":"
	actoken2 := "\""
	actoken3 := value_slk
	actoken4 := "\"}"

	slack_val := actoken1 + actoken2 + actoken3 + actoken4

	url := "https://hooks.slack.com/services/T4asdasdaBRsdsdRB3sdsFB/BJWsdsdGRJVdsd1P/xxxxxxxxxxxxxxxxxxxxxxxxxxx"

	payload := strings.NewReader(slack_val)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
  req.Header.Add("postman-token", "xxxxxxxxxxxxxxxxxxxxxxxxxxx")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	_ = body

}

func final_activity_logs(final_logs_act string) {

	input = final_logs_act

	m := map[string]interface{}{}

	err := json.Unmarshal([]byte(input), &m)

	if err != nil {
		panic(err)
	}
	parseMap(m)

	dup_keys := make(map[interface{}]interface{})
	var list_dup_keys []interface{}

	for _, entry := range duplicate_op {

		if _, value_sg := dup_keys[entry]; !value_sg {
			dup_keys[entry] = true
			list_dup_keys = append(list_dup_keys, entry)

		}

	}

	_ = list_dup_keys
	_ = dup_keys

	for key_sg, entry_sg := range dup_keys {
		_ = entry_sg

		sg_fnl = key_sg.(string)

		mfnla := map[string]interface{}{}

		err := json.Unmarshal([]byte(sg_fnl), &mfnla)

		if err != nil {
			panic(err)
		}
		parseMapfnla(mfnla)

	}

	if user_condition == 1 {


		dup_keys_user := make(map[interface{}]interface{})
		var list_dup_keys_user []interface{}

		for _, entry_user := range duplicate_user {

			if _, value_sg_user := dup_keys_user[entry_user]; !value_sg_user {
				dup_keys_user[entry_user] = true
				list_dup_keys_user = append(list_dup_keys_user, entry_user)

			}

		}

		for key_sg_user, entry_sg_user := range dup_keys_user {
			_ = entry_sg_user

			sg_fnl_user = key_sg_user.(string)

			user_final := strings.Split(sg_fnl_user, "@")

			userfpl := "USER :: @" + user_final[0]

			slack_push(userfpl)

		}

	}

}

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {

		switch concreteVal := val.(type) {

		case map[string]interface{}:

			parseMap(val.(map[string]interface{}))

		case []interface{}:

			parseArray(val.([]interface{}))

		default:

			if key == "localizedValue" {

				if concreteVal == "Create or Update Security Rule" {

					mnsg := map[string]interface{}{}

					err := json.Unmarshal([]byte(input), &mnsg)

					if err != nil {
						panic(err)
					}
					parseMapnsg(mnsg)

				}

			}

		}
	}
}

func parseMapnsg(aMapnsg map[string]interface{}) {
	for keynsg, valnsg := range aMapnsg {

		switch concreteValnsg := valnsg.(type) {

		case map[string]interface{}:

			parseMapnsg(valnsg.(map[string]interface{}))

		case []interface{}:

			parseArraynsg(valnsg.([]interface{}))

		default:

			if keynsg == "localizedValue" {

				if concreteValnsg == "Accepted" {

					mac := map[string]interface{}{}

					err := json.Unmarshal([]byte(input), &mac)

					if err != nil {
						panic(err)
					}
					parseMapac(mac)

				}

			}

		}
	}
}

func parseArraynsg(anArraynsg []interface{}) {
	for insg, valnsg := range anArraynsg {
		switch concreteValnsg := valnsg.(type) {
		case map[string]interface{}:
			parseMapnsg(valnsg.(map[string]interface{}))
		case []interface{}:
			_ = insg
			parseArraynsg(valnsg.([]interface{}))
		default:
			_ = concreteValnsg

		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			_ = i
			parseArray(val.([]interface{}))
		default:
			_ = concreteVal

		}
	}
}

func parseMapac(aMapac map[string]interface{}) {
	for keyac, valac := range aMapac {

		switch concreteValac := valac.(type) {

		case map[string]interface{}:

			parseMapac(valac.(map[string]interface{}))

		case []interface{}:

			parseArrayac(valac.([]interface{}))

		default:

			if keyac == "responseBody" {

				_ = concreteValac
				duplicate_op = append(duplicate_op, concreteValac)

				muser := map[string]interface{}{}

				err := json.Unmarshal([]byte(input), &muser)

				if err != nil {
					panic(err)
				}
				parseMapuser(muser)

			}

		}
	}
}

func parseArrayac(anArrayac []interface{}) {
	for iac, valac := range anArrayac {
		switch concreteValac := valac.(type) {
		case map[string]interface{}:
			parseMapac(valac.(map[string]interface{}))
		case []interface{}:
			_ = iac
			parseArrayac(valac.([]interface{}))
		default:
			_ = concreteValac

		}
	}
}

func activity_logs(access_token string) {

	end_time := time.Now().UTC()

	count := 100000

	start_time := end_time.Add(time.Duration(-count) * time.Minute)
	end_time_end := end_time.String()
	start_time_start := start_time.String()

	end_time_detail := strings.Split(end_time_end, " ")

	dateend := end_time_detail[0]
	timenow := end_time_detail[1]

	timeend_end := strings.Split(timenow, ":")

	timeendhr := timeend_end[0]
	timeendmins := timeend_end[1]

	start_time_detail := strings.Split(start_time_start, " ")

	datestart := start_time_detail[0]
	timestart := start_time_detail[1]

	timestart_start := strings.Split(timestart, ":")

	timestarthr := timestart_start[0]
	timestartmins := timestart_start[1]
	var bearer = "Bearer " + access_token

	url := "https://management.azure.com/subscriptions/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/providers/microsoft.insights/eventtypes/management/values?api-version=2015-04-01&%24filter=eventTimestamp%20ge%20'" + datestart + "T" + timestarthr + "%3A" + timestartmins + "%3A19.0646424Z'%20and%20eventTimestamp%20le%20'" + dateend + "T" + timeendhr + "%3A" + timeendmins + "%3A19.0646424Z'"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", bearer)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "xxxxxxxxxxxxxxxxxxxxxxxx")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fnl_logs := string(body)

	final_activity_logs(fnl_logs)
}

func get_authtoken() {

	url := "https://login.microsoftonline.com/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/oauth2/token"

	payload := strings.NewReader("------WebKitFormBoundary7MA4YWasasxkTrZu0gW\r\nContent-Disposition: form-data; name=\"grant_type\"\r\n\r\nclient_credentials\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"client_id\"\r\n\r\nf147431b-afc3-4f1a-875d-361426bc0258\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"client_secret\"\r\n\r\nbM2@tGjCbC2SbY.ur.iYbD4PH/tIrit0\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"resource\"\r\n\r\nhttps://management.azure.com/\r\n------WebKitFormBoundary7MA4asasasasYWxkTrZu0gW--")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxasaskTrZu0gW")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(body)

	auth_token_response := string(body)
	sec := map[string]interface{}{}
	if err := json.Unmarshal([]byte(auth_token_response), &sec); err != nil {
		panic(err)
	}
	value, ok := sec["access_token"]
	if ok {
		actoken, _ := value.(string)
		activity_logs(actoken)
	} else {
		fmt.Println("access token not not found")
	}
}

func parseMapfnla(aMap map[string]interface{}) {
	for key, val := range aMap {

		switch concreteVal := val.(type) {

		case map[string]interface{}:

			parseMapfnla(val.(map[string]interface{}))

		case []interface{}:

			parseArrayfnla(val.([]interface{}))

		default:

			if key == "sourceAddressPrefix" {

				if concreteVal == "Internet" || concreteVal == "*"   {

					user_condition = 1

					mfnla1 := map[string]interface{}{}

					err := json.Unmarshal([]byte(sg_fnl), &mfnla1)

					if err != nil {
						panic(err)
					}
					parseMapfnla1(mfnla1)

				}
			}

		}
	}
}

func parseMapfnla1(aMap map[string]interface{}) {
	for key, val := range aMap {

		switch concreteVal := val.(type) {

		case map[string]interface{}:

			parseMapfnla1(val.(map[string]interface{}))

		case []interface{}:

			parseArrayfnla1(val.([]interface{}))

		default:

			if key == "id" {

				sg_fnl1 := concreteVal.(string)

				sg_final := strings.Split(sg_fnl1, "/")

				//fmt.Println(sg_final[1])

				//fmt.Println("NSG :", sg_final[8])

				sgfnl := "NSG :: " + sg_final[8]

				slack_push(sgfnl)

				//fmt.Println("Rule :", sg_final[10])

				sgrulefnl := "RULE :: " + sg_final[10]

				slack_push(sgrulefnl)

				desc_nsg := "`For the above NSG group, the rule mentioned has been opened to the world by User in STAGE, pls check and restrict if its not required`"
				slack_push(desc_nsg)

				//fmt.Println("under the above NSG group, for the rule mentioned, it's been opened to the world by User")

			}

		}
	}
}

func parseArrayfnla(anArraynsg []interface{}) {
	for insg, valnsg := range anArraynsg {
		switch concreteValnsg := valnsg.(type) {
		case map[string]interface{}:
			parseMapfnla(valnsg.(map[string]interface{}))
		case []interface{}:
			_ = insg
			parseArrayfnla(valnsg.([]interface{}))
		default:
			_ = concreteValnsg

		}
	}
}

func parseArrayfnla1(anArraynsg1 []interface{}) {
	for insg, valnsg := range anArraynsg1 {
		switch concreteValnsg := valnsg.(type) {
		case map[string]interface{}:
			parseMapfnla1(valnsg.(map[string]interface{}))
		case []interface{}:
			_ = insg
			parseArrayfnla1(valnsg.([]interface{}))
		default:
			_ = concreteValnsg

		}
	}
}

func parseMapuser(aMapuser map[string]interface{}) {
	for keyuser, valuser := range aMapuser {

		switch concreteValuser := valuser.(type) {

		case map[string]interface{}:

			parseMapuser(valuser.(map[string]interface{}))

		case []interface{}:

			parseArrayuser(valuser.([]interface{}))

		default:

			if keyuser == "caller" {

				duplicate_user = append(duplicate_user, concreteValuser)

			}

		}
	}
}

func parseArrayuser(anArrayuser []interface{}) {
	for iuser, valuser := range anArrayuser {
		switch concreteValuser := valuser.(type) {
		case map[string]interface{}:
			parseMapuser(valuser.(map[string]interface{}))
		case []interface{}:
			_ = iuser
			parseArrayuser(valuser.([]interface{}))
		default:
			_ = concreteValuser

		}
	}
}

func main() {
	fmt.Println("")
	get_authtoken()
}
