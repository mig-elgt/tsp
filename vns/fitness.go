package vns

func (s *Solution) Fitness() float32 {
	if s.FitnessValue == 0 {
		r := &s.Routes[0]
		route := s.Customers[r.Start:(r.Start + r.Size)]
		s.FitnessValue = CalculateVehicleRouteCost(0, route, s.Cluster)
	}
	return s.FitnessValue
}

func CalculateVehicleRouteCost(idx int, route []int, cluster *Cluster) float32 {
	distance := cluster.CostMatrix[0][route[0]].Distance
	for i := 0; i < len(route)-1; i++ {
		a, b := route[i], route[i+1]
		distance += cluster.CostMatrix[a][b].Distance
	}
	distance += cluster.CostMatrix[route[len(route)-1]][0].Distance
	return 1 / distance

}
