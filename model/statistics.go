package model

import (
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
)

func ParseRspamdInfos(rspamdSettings RspamdSettings) RspamdStatistics {
	rspamdUrl := rspamdSettings.rspamdUrl + "/stat?password=" + rspamdSettings.rspamdPassword
	res, err := http.Get(rspamdUrl)
	if err != nil {
		log.Fatal(err)
	}

	if (res.StatusCode == 403) {
		log.Fatalf("authorization with to rspamd failed. please check password!")
		os.Exit(1)
	} else if (res.StatusCode == 404) {
		log.Fatalf("http server returned status code 404. please check url!")
		os.Exit(1)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	rspamdStatistics := RspamdStatistics{}
	jsonErr := json.Unmarshal(body, &rspamdStatistics)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return rspamdStatistics
}

