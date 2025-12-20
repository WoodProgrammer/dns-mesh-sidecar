# DNS Proxy

A simple DNS proxy server written in Go that forwards DNS queries to an upstream DNS server.

## Features

- Listens on UDP port 53 (configurable)
- Forwards DNS queries to upstream DNS server (default: 1.1.1.1)
- Concurrent request handling using goroutines
- Configurable upstream DNS server
- Verbose logging option
- 5-second timeout for upstream responses

## Installation

```bash
go build -o dns-proxy
```

## Usage

### Basic usage (requires sudo for port 53):

```bash
sudo ./dns-proxy
```

### Run on a different port (no sudo required):

```bash
./dns-proxy -listen :5353
```

### Use a different upstream DNS server:

```bash
sudo ./dns-proxy -upstream 8.8.8.8:53
```

### Enable verbose logging:

```bash
sudo ./dns-proxy -verbose
```

### Combine options:

```bash
./dns-proxy -listen :5353 -upstream 8.8.8.8:53 -verbose
```

## Command-line Flags

- `-listen`: Address to listen on (default: `:53`)
- `-upstream`: Upstream DNS server address (default: `1.1.1.1:53`)
- `-verbose`: Enable verbose logging (default: `false`)

## Testing

You can test the DNS proxy using `dig` or `nslookup`:

### Using dig:

```bash
# If running on port 53
dig @localhost example.com

# If running on port 5353
dig @localhost -p 5353 example.com
```

### Using nslookup:

```bash
# If running on port 53
nslookup example.com localhost

# If running on port 5353
nslookup -port=5353 example.com localhost
```

## Notes

- Running on port 53 requires root/administrator privileges
- The proxy handles UDP DNS queries only
- Maximum DNS message size is 512 bytes (standard UDP DNS limit)
- Each query is handled in a separate goroutine for concurrent processing

## Example Output

```
DNS Proxy v0.0.1
Listening on: :53
Upstream DNS: 1.1.1.1:53
Starting DNS proxy...
DNS proxy listening on :53
```

With verbose mode enabled:

```
DNS Proxy v0.0.1
Listening on: :53
Upstream DNS: 1.1.1.1:53
Starting DNS proxy...
DNS proxy listening on :53
2025/12/10 15:30:45 Received 45 bytes from 127.0.0.1:54321
2025/12/10 15:30:45 Processing query from 127.0.0.1:54321
2025/12/10 15:30:45 Forwarded query to 1.1.1.1:53
2025/12/10 15:30:45 Received 61 bytes from upstream
2025/12/10 15:30:45 Sent response to 127.0.0.1:54321
```
