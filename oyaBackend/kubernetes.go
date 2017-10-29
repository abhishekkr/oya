package oyaBackend

import "log"

type KubernetesAPI struct {
	SchedularURL         string
	SchedularBearerToken string
}

func (kube *KubernetesAPI) CreateJob() (string, error) {
	log.Println("create job")
	return "WIP", nil
}
