package main // package clause
import (
	"log"
	"time"
	"quadstingray/rspamd-influxdb/model"
)

func main() {

	var settings model.Settings = model.Parser()

	for true {
		log.Println("rspamd information import started")
		stats := model.ParseRspamdInfos(settings.RspamdSettings)
		model.SaveToInfluxDb(stats, settings.InfluxDbSettings)
		log.Printf("sleep for %v seconds", settings.Interval)
		time.Sleep(time.Duration(settings.Interval) * time.Second)
	}

}


