# Optimizer

Optimizer Services provides an API to start a route opitmization. It implements a Table gRPC client to get a distance matrix and a VNS gRPC clients to performs the optimization.

## Features

* Endpoint to start a route optimization for TSP.

## API Usage & Example

URL Base: http://localhost:8080 for local tests or use your own domain.

### TSP Endpoint

Endpoint: /api/v1/tsp

Example
POST http://localhost:8080/api/v1/tsp

```json
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
```

Response Status 200 (OK)

```json
{
	"route": [
	  {
	     "name": "C",
		 "location": {
		    "lat": 52.523239,
			"lng": 13.428554
		  }
	  },
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
	  }
	],
	"total_distance": 6635.8
}
```

Response Status 500 (Internal Server Error)

```json
{
   "error": {
      "status": 500,
      "error": "INTERNAL",
      "description": "Something went wrong..."
	}
}
```
