package tsdb

import (
	"fmt"
	"os"
	"testing"
	"time"

	l "github.com/go-kit/kit/log"
	"github.com/prometheus/tsdb/fileutil"
)

var (
	w = &SegmentWAL{
		dirFile:       getDF(),
		logger:        l.NewNopLogger(),
		flushInterval: 20 * time.Second,
		donec:         make(chan struct{}),
		stopc:         make(chan struct{}),
		actorc:        make(chan func() error, 1),
		segmentSize:   walSegmentSizeBytes,
		crc32:         newCRC32(),
	}
)

func getDF() *os.File {
	df, err := fileutil.OpenDir("/Users/emarcian/rnd/prometheus/wal")
	if err != nil {
		fmt.Println("bla")
	}
	return df
}

func TestSegementSync(t *testing.T) {
	w.Sync()
}

func TestSegementLogDeletes(t *testing.T) {
	s1 := []Stone{Stone{ref: 1, intervals: []Interval{Interval{Maxt: 2, Mint: 1}}}}
	err := w.LogDeletes(s1)
	if err != nil {
		fmt.Println("LogDeletes Failed")
	}
}
