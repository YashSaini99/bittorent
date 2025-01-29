package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Command line flags
	torrentPath := flag.String("torrent", "", "Path to .torrent file")
	magnetURL := flag.String("magnet", "", "Magnet URL")
	outputDir := flag.String("output", "downloads", "Output directory for downloaded files")
	flag.Parse()

	if *torrentPath == "" && *magnetURL == "" {
		log.Fatal("Either -torrent or -magnet flag must be provided")
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatal("Failed to create output directory:", err)
	}

	client := NewTorrentClient(*outputDir)
	defer client.Close()

	if *torrentPath != "" {
		if err := client.AddTorrentFile(*torrentPath); err != nil {
			log.Fatal("Failed to add torrent file:", err)
		}
	} else {
		if err := client.AddMagnetURL(*magnetURL); err != nil {
			log.Fatal("Failed to add magnet URL:", err)
		}
	}

	// Handle Ctrl+C gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		client.Close()
		os.Exit(0)
	}()

	// Start downloading and wait for completion
	client.Start()
}
