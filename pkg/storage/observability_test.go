package storage

import (
    "testing"
    "time"
    "github.com/sirrobot01/decypharr/internal/config"
)

func TestBuildDownloadObservability(t *testing.T) {
    now := time.Now()
    e := &Entry{Protocol: config.ProtocolTorrent, AddedOn: now, UpdatedAt: now, IsDownloading: true, DownloadUncached: true, ActiveProvider: "torbox", Providers: map[string]*ProviderEntry{"torbox": {Provider: "torbox", ID: "abc", Status: "downloading"}}}
    o := BuildDownloadObservability(e)
    if o.PipelineStage != "downloading" || o.CacheMode != "uncached" || o.ProviderID != "abc" || o.ProviderState != "downloading" { t.Fatalf("unexpected observability: %+v", o) }
}
