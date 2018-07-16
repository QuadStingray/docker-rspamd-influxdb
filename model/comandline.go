package model

import (
	"flag"
	"log"
)

func Parser() Settings {
	var interval int
	var rspamdUrl string
	var rspamdPassword string
	var influxHost string
	var influxDB string
	var influxPwd string
	var influxUser string

	flag.IntVar(&interval, "interval", 3600, "seconds between statistics import")
	flag.StringVar(&rspamdUrl, "rspamdUrl", "https://RSPAMD-URL", "url of your rspamd webinterface installation")
	flag.StringVar(&rspamdPassword, "rspamdPassword", "PASSWORD", "password for rspamd authentication")

	flag.StringVar(&influxHost, "influxHost", "http://influxdb:8086", "host of your influxdb instance")
	flag.StringVar(&influxDB, "influxDB", "rspamd", "influxdb database")
	flag.StringVar(&influxUser, "influxUser", "DEFAULT", "influxdb username")
	flag.StringVar(&influxPwd, "influxPwd", "DEFAULT", "influxdb password")

	flag.Parse()

	log.Println("**************************************************************")
	log.Println("******** Parser started with following commands **************")
	log.Printf("**  interval %v", interval)
	log.Println("**  rspamdUrl " + rspamdUrl)

	if (rspamdPassword != "DEFAULT") {
		log.Println("**  rspamdPassword DEFAULT")
	} else {
		log.Println("**  rspamdPassword is not Default")
	}

	log.Println("**  influxHost " + influxHost)
	log.Println("**  influxDB " + influxDB)
	log.Println("**  influxUser " + influxUser)

	log.Println("**  influxPwd " + influxPwd)
	if (influxPwd != "DEFAULT") {
		log.Println("**  influxPwd DEFAULT")
	} else {
		log.Println("**  influxPwd is not Default")
	}

	log.Println("**************************************************************")
	log.Println("**************************************************************")
	return Settings{interval, RspamdSettings{rspamdUrl, rspamdPassword}, InfluxDbSettings{influxHost, influxUser, influxPwd, influxDB}}
}


