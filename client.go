package main

import (
	"log"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/cheggaaa/pb/v3"
)

type TorrentClient struct {
	client    *torrent.Client
	outputDir string
	torrents  []*torrent.Torrent
}

func NewTorrentClient(outputDir string) *TorrentClient {
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = outputDir

	client, err := torrent.NewClient(cfg)
	if err != nil {
		log.Fatal("Failed to create torrent client:", err)
	}

	return &TorrentClient{
		client:    client,
		outputDir: outputDir,
		torrents:  make([]*torrent.Torrent, 0),
	}
}

func (c *TorrentClient) AddTorrentFile(path string) error {
	t, err := c.client.AddTorrentFromFile(path)
	if err != nil {
		return err
	}
	c.torrents = append(c.torrents, t)
	return nil
}

func (c *TorrentClient) AddMagnetURL(url string) error {
	t, err := c.client.AddMagnet(url)
	if err != nil {
		return err
	}
	c.torrents = append(c.torrents, t)
	return nil
}

func (c *TorrentClient) Start() {
	for _, t := range c.torrents {
		<-t.GotInfo()
		t.DownloadAll()

		bar := pb.Full.Start64(t.Length())
		bar.Set(pb.Bytes, true)
		bar.SetMaxWidth(100)

		done := make(chan struct{})
		go func() {
			defer close(done)
			for {
				select {
				case <-time.After(time.Second):
					stats := t.Stats()
					bytesRead := stats.BytesRead.Int64()
					bar.SetCurrent(bytesRead)
					if bytesRead >= t.Length() {
						return
					}
				}
			}
		}()

		<-t.Closed()
		bar.Finish()
		<-done
	}
}

func (c *TorrentClient) Close() {
	c.client.Close()
}
