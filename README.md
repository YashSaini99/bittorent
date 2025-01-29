# Go BitTorrent Client

A lightweight and efficient BitTorrent client implemented in Go, supporting both torrent files and magnet links.

## Features

- Download torrents using .torrent files or magnet URLs
- Progress bar with real-time download statistics
- Docker support for easy deployment
- Configurable download directory
- Graceful shutdown handling

## Installation

### Using Docker

```bash
# Pull from Docker Hub
docker pull yashsaini99/bittorrent

# Or build locally
docker build -t bittorrent .
```

### Building from Source

```bash
# Clone the repository
git clone https://github.com/yashsaini99/bittorrent.git

# Navigate to project directory
cd bittorrent

# Build the project
go build
```

## Usage

### Docker Usage

```bash
# Using magnet link
docker run -v $(pwd)/downloads:/app/downloads yashsaini99/bittorrent -output /app/downloads -magnet "your_magnet_url"

# Using torrent file
docker run -v $(pwd):/app/torrents -v $(pwd)/downloads:/app/downloads yashsaini99/bittorrent -output /app/downloads -torrent "/app/torrents/file.torrent"
```

### Command Line Usage

```bash
# Using magnet link
./bittorrent -magnet "your_magnet_url" -output "./downloads"

# Using torrent file
./bittorrent -torrent "path/to/file.torrent" -output "./downloads"
```

## Command Line Arguments

- `-magnet`: Magnet URL for the torrent
- `-torrent`: Path to .torrent file
- `-output`: Directory for downloaded files (default: "downloads")

## Dependencies

- [github.com/anacrolix/torrent](https://github.com/anacrolix/torrent)
- [github.com/cheggaaa/pb](https://github.com/cheggaaa/pb)

## Building the Docker Image

```bash
docker build -t bittorrent .
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

```
