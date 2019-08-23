function() {
	"listenAddress": {
		"type": "config",
		"value": "0.0.0.0:7201"
	},
	"logging": {
		"level": "info"
	},
	"metrics": {
		"scope": {
			"prefix": "coordinator"
		},
		"prometheus": {
			"handlerPath": "/metrics",
			"listenAddress": "0.0.0.0:7203"
		},
		"sanitization": "prometheus",
		"samplingRate": 1,
		"extended": "none"
	},
	"clusters": [
		{
			"namespaces": [
				{
					"namespace": "default",
					"type": "unaggregated",
					"retention": "48h"
				}
			],
			"client": {
				"config": {
					"service": {
						"env": "default_env",
						"zone": "embedded",
						"service": "m3db",
						"cacheDir": "/var/lib/m3kv",
						"etcdClusters": [
							{
								"zone": "embedded",
								"endpoints": [
									"127.0.0.1:2379"
								]
							}
						]
					}
				},
				"writeConsistencyLevel": "majority",
				"readConsistencyLevel": "unstrict_majority"
			}
		}
	],
	"tagOptions": {
		"idScheme": "quoted"
	}
}
