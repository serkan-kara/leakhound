package finding

import "time"

type ScanResult struct {
	Findings     []Finding
	FilesScanned int
	Duration     time.Duration
}
