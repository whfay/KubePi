package cluster

import (
	"time"

	v1Cluster "github.com/KubeOperator/kubepi/internal/model/v1/cluster"
	V1ClusterRepo "github.com/KubeOperator/kubepi/internal/model/v1/clusterrepo"
)

const (
	clusterStatusInitializing = "Initializing"
	clusterStatusFailed       = "Failed"
	clusterStatusCompleted    = "Completed"
	clusterStatusSaved        = "Saved"
)

type Cluster struct {
	v1Cluster.Cluster
	KeyDataStr           string           `json:"keyDataStr"`
	CertDataStr          string           `json:"certDataStr"`
	CaDataStr            string           `json:"caDataStr"`
	ConfigFileContentStr string           `json:"configContentStr"`
	Accessable           bool             `json:"accessable"`
	MemberCount          int              `json:"memberCount"`
	ExtraClusterInfo     ExtraClusterInfo `json:"extraClusterInfo"`
}

type UpdateCluster struct {
	Labels []string `json:"labels"`
}

type ExtraClusterInfo struct {
	TotalNodeNum      int     `json:"totalNodeNum"`
	ReadyNodeNum      int     `json:"readyNodeNum"`
	CPUAllocatable    float64 `json:"cpuAllocatable"`
	CPURequested      float64 `json:"cpuRequested"`
	MemoryAllocatable float64 `json:"memoryAllocatable"`
	MemoryRequested   float64 `json:"memoryRequested"`
	Health            bool    `json:"health"`
	Message           string  `json:"message"`
}

type NamespaceRoles struct {
	Namespace string   `json:"namespace"`
	Roles     []string `json:"roles"`
}

type Member struct {
	Name           string           `json:"name"`
	ClusterRoles   []string         `json:"clusterRoles"`
	BindingName    string           `json:"bindingName"`
	CreateAt       time.Time        `json:"createAt"`
	NamespaceRoles []NamespaceRoles `json:"namespaceRoles"`
}

type Privilege struct {
	Url string `json:"url"`
}

type Repo struct {
	V1ClusterRepo.ClusterRepo
}

type CreateRepo struct {
	Repos   []string
	Cluster string
}
