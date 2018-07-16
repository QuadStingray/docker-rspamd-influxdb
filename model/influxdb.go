package model

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

func SaveToInfluxDb(statistics RspamdStatistics, influxDbSettings InfluxDbSettings) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     influxDbSettings.dbUrl,
		Username: influxDbSettings.username,
		Password: influxDbSettings.password,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  influxDbSettings.db,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"host": "local"}
	fields := map[string]interface{}{
		"reject":          statistics.Actions.Reject,
		"soft_reject":     statistics.Actions.SoftReject,
		"rewrite_subject": statistics.Actions.RewriteSubject,
		"add_header":      statistics.Actions.AddHeader,
		"greylist":        statistics.Actions.Greylist,
		"no_action":       statistics.Actions.NoAction,
	}

	pt, err := client.NewPoint("rspamd_actions", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	fieldsStatistics := map[string]interface{}{
		"scanned":                 statistics.Scanned,
		"learned":                 statistics.Learned,
		"spam_count":              statistics.SpamCount,
		"ham_count":               statistics.HamCount,
		"connections":             statistics.Connections,
		"control_connections":     statistics.ControlConnections,
		"pools_allocated":         statistics.PoolsAllocated,
		"pools_freed":             statistics.PoolsFreed,
		"bytes_allocated":         statistics.BytesAllocated,
		"chunks_allocated":        statistics.ChunksAllocated,
		"shared_chunks_allocated": statistics.SharedChunksAllocated,
		"chunks_freed":            statistics.ChunksFreed,
		"chunks_oversized":        statistics.ChunksOversized,
		"fragmented":              statistics.Fragmented,
		"total_learns":            statistics.TotalLearns,
		"fuzzy":                   statistics.FuzzyHashes.RspamdCom,
	}

	ptStats, err := client.NewPoint("rspamd_stats", tags, fieldsStatistics, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(ptStats)

	err = c.Write(bp)
	if err != nil {
		log.Fatalf("could not write to influx db. check connection to %v and db %s with user %v with pwd %s", influxDbSettings.dbUrl, influxDbSettings.db, influxDbSettings.username, influxDbSettings.password)
	}
}
