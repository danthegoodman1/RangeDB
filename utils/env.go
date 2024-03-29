package utils

import "os"

var (
	Env                    = os.Getenv("ENV")
	Env_TracingServiceName = os.Getenv("TRACING_SERVICE_NAME")
	Env_OLTPEndpoint       = os.Getenv("OLTP_ENDPOINT")

	Env_APIPort       = GetEnvOrDefault("API_PORT", "8190")
	Env_InternalPort  = GetEnvOrDefault("INTERNAL_PORT", "8191")
	Env_GossipPort    = MustEnvOrDefaultInt64("GOSSIP_PORT", 8192)
	Env_AdvertiseAddr = os.Getenv("ADVERTISE_ADDR") // API, e.g. localhost:8190
	// csv like localhost:8192,localhost:8292
	Env_GossipPeers       = os.Getenv("GOSSIP_PEERS")
	Env_GossipBroadcastMS = MustEnvOrDefaultInt64("GOSSIP_BROADCAST_MS", 1000)
	Env_GossipDebug       = os.Getenv("GOSSIP_DEBUG") == "1"

	Env_GCIntervalMs = MustEnvOrDefaultInt64("GC_INTERVAL_MS", 60_000*5) // 5 minute default
	Env_DBPath       = GetEnvOrDefault("DB_PATH", "/var/rangedb/data")
)
