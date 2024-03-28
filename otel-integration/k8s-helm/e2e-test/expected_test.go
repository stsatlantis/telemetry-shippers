package e2e

var expectedSchemaURL = map[string]bool{
	"https://opentelemetry.io/schemas/1.6.1": false,
	"https://opentelemetry.io/schemas/1.9.0": false,
}

const expectedScopeVersion = "0.96.0"

var expectedScopeNames = map[string]bool{
	"otelcol/hostmetricsreceiver/network":    false,
	"otelcol/hostmetricsreceiver/cpu":        false,
	"otelcol/hostmetricsreceiver/filesystem": false,
	"otelcol/hostmetricsreceiver/memory":     false,
	"otelcol/hostmetricsreceiver/load":       false,
	"otelcol/hostmetricsreceiver/disk":       false,
	"otelcol/kubeletstatsreceiver":           false,
	"otelcol/prometheusreceiver":             false,
}

var expectedResourceAttributesKubeletstatreceiver = map[string]string{
	"k8s.pod.uid":              "",
	"k8s.pod.name":             "",
	"k8s.namespace.name":       "",
	"k8s.container.name":       "",
	"k8s.daemonset.name":       "",
	"k8s.deployment.name":      "",
	"cx.otel_integration.name": "coralogix-integration-helm",
	"k8s.cluster.name":         "otel-integration-agent-e2e",
	"k8s.node.name":            "otel-integration-agent-e2e-control-plane",
	"host.name":                "otel-integration-agent-e2e-control-plane",
	"host.id":                  "",
	"os.type":                  "linux",
}

var expectedResourceAttributesHostmetricsreceiver = map[string]string{
	"cx.otel_integration.name": "coralogix-integration-helm",
	"k8s.cluster.name":         "otel-integration-agent-e2e",
	"k8s.node.name":            "otel-integration-agent-e2e-control-plane",
	"host.name":                "otel-integration-agent-e2e-control-plane",
	"host.id":                  "",
	"os.type":                  "linux",
}

var expectedResourceAttributesPrometheusreceiver = map[string]string{
	"service.name":             "opentelemetry-collector",
	"net.host.name":            "",
	"service.instance.id":      "",
	"net.host.port":            "",
	"http.scheme":              "http",
	"service_version":          expectedScopeVersion,
	"cx.otel_integration.name": "coralogix-integration-helm",
	"k8s.cluster.name":         "otel-integration-agent-e2e",
	"k8s.node.name":            "otel-integration-agent-e2e-control-plane",
	"host.name":                "otel-integration-agent-e2e-control-plane",
	"host.id":                  "",
	"os.type":                  "linux",
}

var expectedMetrics map[string]bool = map[string]bool{
	"system.network.connections":                     false,
	"system.network.dropped":                         false,
	"system.network.errors":                          false,
	"system.network.io":                              false,
	"system.network.packets":                         false,
	"system.cpu.time":                                false,
	"system.filesystem.inodes.usage":                 false,
	"system.filesystem.usage":                        false,
	"system.memory.usage":                            false,
	"system.memory.utilization":                      false,
	"system.cpu.load_average.15m":                    false,
	"system.cpu.load_average.1m":                     false,
	"system.cpu.load_average.5m":                     false,
	"system.cpu.utilization":                         false,
	"system.disk.io":                                 false,
	"system.disk.io_time":                            false,
	"system.disk.merged":                             false,
	"system.disk.operation_time":                     false,
	"system.disk.operations":                         false,
	"system.disk.pending_operations":                 false,
	"system.disk.weighted_io_time":                   false,
	"k8s.node.cpu.time":                              false,
	"k8s.node.cpu.utilization":                       false,
	"k8s.node.filesystem.available":                  false,
	"k8s.node.filesystem.capacity":                   false,
	"k8s.node.filesystem.usage":                      false,
	"k8s.node.memory.available":                      false,
	"k8s.node.memory.major_page_faults":              false,
	"k8s.node.memory.page_faults":                    false,
	"k8s.node.memory.rss":                            false,
	"k8s.node.memory.usage":                          false,
	"k8s.node.memory.working_set":                    false,
	"k8s.node.network.errors":                        false,
	"k8s.node.network.io":                            false,
	"k8s.pod.cpu.time":                               false,
	"k8s.pod.cpu.utilization":                        false,
	"k8s.pod.filesystem.available":                   false,
	"k8s.pod.filesystem.capacity":                    false,
	"k8s.pod.filesystem.usage":                       false,
	"k8s.pod.memory.available":                       false,
	"k8s.pod.memory.major_page_faults":               false,
	"k8s.pod.memory.page_faults":                     false,
	"k8s.pod.memory.rss":                             false,
	"k8s.pod.memory.usage":                           false,
	"k8s.pod.memory.working_set":                     false,
	"k8s.pod.network.errors":                         false,
	"k8s.pod.network.io":                             false,
	"container.cpu.time":                             false,
	"container.cpu.utilization":                      false,
	"container.filesystem.available":                 false,
	"container.filesystem.capacity":                  false,
	"container.filesystem.usage":                     false,
	"container.memory.available":                     false,
	"container.memory.major_page_faults":             false,
	"container.memory.page_faults":                   false,
	"container.memory.rss":                           false,
	"container.memory.usage":                         false,
	"container.memory.working_set":                   false,
	"otelcol_process_memory_rss":                     false,
	"otelcol_processor_refused_metric_points":        false,
	"otelcol_receiver_accepted_metric_points":        false,
	"scrape_duration_seconds":                        false,
	"otelcol_exporter_queue_capacity":                false,
	"otelcol_otelsvc_k8s_ip_lookup_miss":             false,
	"otelcol_process_runtime_total_alloc_bytes":      false,
	"otelcol_otelsvc_k8s_pod_added":                  false,
	"otelcol_process_runtime_total_sys_memory_bytes": false,
	"otelcol_process_uptime":                         false,
	"otelcol_processor_accepted_metric_points":       false,
	"otelcol_processor_batch_metadata_cardinality":   false,
	"otelcol_receiver_refused_log_records":           false,
	"otelcol_receiver_refused_metric_points":         false,
	"otelcol_processor_dropped_metric_points":        false,
	"scrape_samples_post_metric_relabeling":          false,
	"otelcol_exporter_queue_size":                    false,
	"otelcol_exporter_send_failed_metric_points":     false,
	"otelcol_exporter_sent_metric_points":            false,
	"otelcol_exporter_sent_log_records":              false,
	"otelcol_otelsvc_k8s_pod_table_size":             false,
	"otelcol_otelsvc_k8s_pod_updated":                false,
	"otelcol_scraper_errored_metric_points":          false,
	"up":                                             false,
	"otelcol_otelsvc_k8s_namespace_added":            false,
	"otelcol_process_cpu_seconds":                    false,
	"otelcol_process_runtime_heap_alloc_bytes":       false,
	"otelcol_scraper_scraped_metric_points":          false,
	"scrape_samples_scraped":                         false,
	"scrape_series_added":                            false,
	"otelcol_receiver_accepted_log_records":          false,
	"otelcol_processor_batch_timeout_trigger_send":   false,
	"otelcol_exporter_send_failed_log_records":       false,
	"otelcol_rpc_client_duration":                    false,
	"otelcol_processor_batch_batch_send_size":        false,
	"otelcol_rpc_client_responses_per_rpc":           false,
	"otelcol_rpc_client_response_size":               false,
	"otelcol_rpc_client_request_size":                false,
	"otelcol_rpc_client_requests_per_rpc":            false,
}