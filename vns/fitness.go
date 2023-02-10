package vns

func (s *Solution) Fitness() float32 {
	if s.FitnessValue == 0 {
		r := &s.Routes[0]
		route := s.Customers[r.Start:(r.Start + r.Size)]
		s.FitnessValue = CalculateVehicleRouteCost(0, route, s.Cluster)
	}
	return s.FitnessValue
}

const timeWindowPenaltyCost = 20000
const waitingTimePenaltyCost = 1000

func CalculateVehicleRouteCost(idx int, route []int, cluster *Cluster) float32 {
	var waitingTime, timeWindowPenalty, waitingTimePenalty, capacityTotal, cost float32
	if len(route) > 0 {
		stops := cluster.Stops
		customer := route[0] - 1
		arrivalTime := cluster.Vehicles[idx].ShiftStartNumber + (cluster.CostMatrix[0][customer].Time / 60.0)
		if arrivalTime < stops[customer].StartTimeWindow {
			waitingTime = stops[customer].StartTimeWindow - arrivalTime
		} else if arrivalTime > stops[customer].EndTimeWindow {
			timeWindowPenalty += (arrivalTime - stops[customer].EndTimeWindow) * timeWindowPenaltyCost
		}
		departureTime := arrivalTime + waitingTime + float32(stops[customer].DurationServiceTime)/60.0
		distanceTotal := cluster.CostMatrix[0][route[0]].Distance
		capacityTotal += stops[route[0]-1].Demand
		for i := 0; i < len(route)-1; i++ {
			var waitingTime float32 = 0.0
			from := route[i] - 1
			to := route[i+1] - 1
			arrivalTime = departureTime + (cluster.CostMatrix[from][to].Time / 60.0)
			if arrivalTime < stops[to].StartTimeWindow {
				waitingTime = stops[to].StartTimeWindow - arrivalTime
				waitingTimePenalty += (waitingTime * waitingTimePenaltyCost)
			} else if arrivalTime > stops[to].EndTimeWindow {
				timeWindowPenalty += (arrivalTime - stops[to].EndTimeWindow) * timeWindowPenaltyCost
			}
			departureTime = arrivalTime + waitingTime + float32(stops[to].DurationServiceTime)/60.0
			distanceTotal += cluster.CostMatrix[route[i]][route[i+1]].Distance
			capacityTotal += stops[to].Demand
		}
		distanceTotal += cluster.CostMatrix[route[len(route)-1]][0].Distance
		cost += distanceTotal
	}
	return cost

}
