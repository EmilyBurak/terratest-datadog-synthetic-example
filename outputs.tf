# Output variables -- monitor ID needed for API call for terratest

output "synthetic-monitor-id" {
  value = datadog_synthetics_test.example.monitor_id
}