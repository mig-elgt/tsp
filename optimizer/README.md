# Optimizer

Optimizer Services provides an API to start a route opitmization. It implements a Table gRPC client to get a distance matrix and a VNS gRPC clients to performs the optimization.

## Features

* Endpoint to start a route optimization for TSP.

## API Usage & Example

URL Base: http://localhost:8080 for local tests or use your own domain.

### TSP Endpoint

```
Endpoint: /api/v1/tsp

Example

Request:
POST http://localhost:8080/api/v1/tsp

Body
{
	"stops": [
		{
			"name": "A",
			"location": {
				"lat": 52.517033,
				"lng": 13.388798
			}
		},
		{
			"name": "B",
			"location": {
				"lat": 52.529432,
				"lng": 13.39763
			}
		},
		{
			"name": "C",
			"location": {
				"lat": 52.523239,
				"lng": 13.428554
			}
		}
	]
}

Responses:

Status: 200 - Ok
Body Respose
{
	"route": [
		{
			"name": "C",
			"stop_id": 3,
			"location": {
				"lat": 52.523239,
				"lng": 13.428554
			}
		},
		{
			"name": "A",
			"stop_id": 1,
			"location": {
				"lat": 52.517033,
				"lng": 13.388798
			}
		},
		{
			"name": "B",
			"stop_id": 2,
			"location": {
				"lat": 52.529432,
				"lng": 13.39763
			}
		}
	],
	"total_distance": 6635.8
}

Status: 500 - Internal Server Error
Body Respose
{
   "error": {
      "status": 500,
      "error": "INTERNAL",
      "description": "Something went wrong...",
   }
}
```
