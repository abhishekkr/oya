package oyaController

import (
	"log"

	oyaBackend "github.com/abhishekkr/oya/oyaBackend"

	"github.com/gin-gonic/gin"
)

type Kubernetes struct {
	KubernetesAPI oyaBackend.KubernetesAPI
}

func (kube *Kubernetes) Create(ctx *gin.Context) {
	if ctx.Param("type") == "job" {
		kube.KubernetesAPI.CreateJob()
	} else {
		log.Println("unhandled create required")
	}
	log.Println("WIP")
	ctx.String(200, "WIP")
}
