package config

// MachineReplicas are a number of machine following the same specification.
type MachineReplicas struct {
	Spec  Machine
	Count int
}

// Cluster is a set of Machines.
type Cluster struct {
	// Name is the cluster name. Defaults to "cluster".
	Name string

	// PrivateKey is the path to the private SSH key used to login into the cluster
	// machines.
	PrivateKey string

	// Templates describe the machines we want created for this cluster.
	Templates []MachineReplicas
}