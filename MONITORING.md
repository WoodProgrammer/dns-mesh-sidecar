# Monitoring DNS Activity in Sidecar Containers

This document describes the monitoring capabilities of the DNS mesh sidecar, which exposes metrics in Prometheus format for observability and troubleshooting.

## Metrics Endpoint

The sidecar exposes metrics at the following endpoint:

- **Default Address**: `:9090/metrics`
- **Format**: Prometheus exposition format

You can access the metrics by sending an HTTP GET request to `http://<sidecar-host>:9090/metrics`.

## Available Metrics

The DNS mesh sidecar tracks various metrics to help you monitor DNS query performance, errors, and policy enforcement:

### DNS Query Metrics

- `dns_queries_total` - Total number of DNS queries processed
- `dns_query_duration_seconds` - Histogram of DNS query durations
- `dns_upstream_queries_total` - Total number of queries forwarded to upstream DNS servers

### Error Metrics

The sidecar tracks errors by type, allowing you to identify specific failure modes:

```go
ErrorTypeParse           = "parse"           // DNS query parsing errors
ErrorTypeUpstreamDial    = "upstream_dial"    // Failed to connect to upstream DNS server
ErrorTypeUpstreamWrite   = "upstream_write"   // Failed to write query to upstream server
ErrorTypeUpstreamRead    = "upstream_read"    // Failed to read response from upstream server
ErrorTypeUpstreamTimeout = "upstream_timeout" // Upstream DNS server timeout
ErrorTypeClientWrite     = "client_write"     // Failed to write response to client
ErrorTypePolicyFetch     = "policy_fetch"     // Failed to fetch DNS policy
```

- `dns_errors_total{type="<error_type>"}` - Counter of errors by type

### Policy Metrics

- `dns_policy_updates_total` - Total number of policy updates received
- `dns_policy_fetch_duration_seconds` - Histogram of policy fetch durations

## Grafana Dashboard

A pre-configured Grafana dashboard is available for visualizing the sidecar metrics.

### Importing the Dashboard

1. Navigate to your Grafana instance
2. Go to **Dashboards** → **Import**
3. Upload the dashboard JSON file: `dashboards/sidecar-dashboard.json`
4. Select your Prometheus data source
5. Click **Import**

The dashboard includes panels for:
- Query rate and latency
- Error rates by type
- Upstream server performance
- Policy fetch status

## Example Prometheus Configuration

To scrape metrics from the sidecar, add the following to your `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: 'dns-mesh-sidecar'
    static_configs:
      - targets: ['localhost:9090']
    metrics_path: '/metrics'
    scrape_interval: 15s
```

## Alerting Recommendations

Consider setting up alerts for:

- High error rates (particularly `upstream_timeout` and `policy_fetch` errors)
- Elevated DNS query latency
- Policy fetch failures
- Upstream connectivity issues

Example Prometheus alert rule:

```yaml
groups:
  - name: dns_sidecar_alerts
    rules:
      - alert: HighDNSErrorRate
        expr: rate(dns_errors_total[5m]) > 0.1
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High DNS error rate detected"
          description: "DNS sidecar error rate is {{ $value }} errors/sec"
```
