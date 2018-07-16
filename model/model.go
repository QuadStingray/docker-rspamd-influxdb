package model

type RspamdStatistics struct {
	ReadOnly bool  `json:"read_only"`
	Scanned  int32 `json:"scanned"`
	Learned  int32 `json:"learned"`
	Actions struct {
		Reject         int32 `json:"reject"`
		SoftReject     int32 `json:"soft reject"`
		RewriteSubject int32 `json:"rewrite subject"`
		AddHeader      int32 `json:"add header"`
		Greylist       int32 `json:"greylist"`
		NoAction       int32 `json:"no action"`
	} `json:"actions"`
	SpamCount             int32 `json:"spam_count"`
	HamCount              int32 `json:"ham_count"`
	Connections           int32 `json:"connections"`
	ControlConnections    int32 `json:"control_connections"`
	PoolsAllocated        int32 `json:"pools_allocated"`
	PoolsFreed            int32 `json:"pools_freed"`
	BytesAllocated        int32 `json:"bytes_allocated"`
	ChunksAllocated       int32 `json:"chunks_allocated"`
	SharedChunksAllocated int32 `json:"shared_chunks_allocated"`
	ChunksFreed           int32 `json:"chunks_freed"`
	ChunksOversized       int32 `json:"chunks_oversized"`
	Fragmented            int32 `json:"fragmented"`
	TotalLearns           int32 `json:"total_learns"`
	Statfiles []struct {
		Revision  int32  `json:"revision"`
		Used      int32  `json:"used"`
		Total     int32  `json:"total"`
		Size      int32  `json:"size"`
		Symbol    string `json:"symbol"`
		Type      string `json:"type"`
		Languages int32  `json:"languages"`
		Users     int32  `json:"users"`
	} `json:"statfiles"`
	FuzzyHashes struct {
		RspamdCom int `json:"rspamd.com"`
	} `json:"fuzzy_hashes"`
}

type InfluxDbSettings struct {
	dbUrl    string
	username string
	password string
	db       string
}

type RspamdSettings struct {
	rspamdUrl      string
	rspamdPassword string
}

type Settings struct {
	Interval         int
	RspamdSettings   RspamdSettings
	InfluxDbSettings InfluxDbSettings
}
