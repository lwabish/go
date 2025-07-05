package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/gin-gonic/gin"
	"github.com/lwabish/go/pkg/k8s"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	batchV1 "k8s.io/api/batch/v1"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"log"
	"net/http"
	"time"
)

var (
	pveApiUrl   = ""
	pveUser     = ""
	pvePassword = ""
	pveTimeout  = 300
)

const (
	responseOK   = "OK"
	jobNamespace = "default"
)

var homeCmd = &cobra.Command{
	Use: "home",
	Run: func(cmd *cobra.Command, args []string) {

		pveClient, err := proxmox.NewClient(pveApiUrl, nil, "", &tls.Config{InsecureSkipVerify: true}, "", pveTimeout)
		if err != nil {
			log.Fatalln(err)
		}
		pveClient.SetAPIToken(pveUser, pvePassword)

		k8s.InitClient()

		r := gin.Default()
		r.Group("api/v1")
		r.Any("/node", node(pveClient, k8s.GetClient()))
		r.Any("/vm", vm(pveClient))
		if err := r.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	},
}

func node(pc *proxmox.Client, kc *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			nodes, err := pc.GetNodeList(c)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, nodes)
		case http.MethodPost:
			param := struct {
				Name string `json:"name"`
				Op   string `json:"op"`
				Mac  string `json:"mac"`
			}{}
			if err := c.ShouldBindBodyWithJSON(&param); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if param.Op == "start" {
				macAddr := param.Mac
				j, err := kc.BatchV1().Jobs(jobNamespace).Create(c, &batchV1.Job{
					ObjectMeta: metav1.ObjectMeta{
						GenerateName: "home-api-send-wol-",
					},
					Spec: batchV1.JobSpec{
						TTLSecondsAfterFinished: lo.ToPtr(int32(3 * 3600)),
						Template: coreV1.PodTemplateSpec{
							Spec: coreV1.PodSpec{
								Containers: []coreV1.Container{
									{
										Name:  "home-api-send-wol",
										Image: "lwabish/go-wol",
										Args: []string{
											"wake", macAddr,
										},
									},
								},
								HostNetwork:   true,
								RestartPolicy: coreV1.RestartPolicyNever,
							},
						},
					},
				}, metav1.CreateOptions{})
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				if err = wait.PollUntilContextTimeout(c, 500*time.Millisecond, 30*time.Second, false,
					func(ctx context.Context) (done bool, err error) {
						job, err := kc.BatchV1().Jobs(jobNamespace).Get(ctx, j.GetName(), metav1.GetOptions{})
						if err != nil {
							log.Printf("Get job %s failed: %v", j.GetName(), err)
							return false, nil
						}
						if job.Status.Succeeded == 1 {
							return true, nil
						} else if job.Status.Failed == 1 {
							return true, fmt.Errorf("job failed")
						}
						return false, nil
					}); err != nil {
					// timeout is included
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				c.JSON(200, responseOK)
				return
			} else if param.Op == "shutdown" {
				s, err := pc.ShutdownNode(c, param.Name)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				c.JSON(200, lo.Ternary(s == "", responseOK, s))
				return
			}
		default:
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
		}
	}
}

func vm(pc *proxmox.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			vms, err := pc.GetVmList(c)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			c.JSON(200, vms)
		case http.MethodPost:
			param := struct {
				Name string `json:"name"`
				Op   string `json:"op"`
			}{}
			if err := c.ShouldBindBodyWithJSON(&param); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			vmRef, err := pc.GetVmRefByName(c, proxmox.GuestName(param.Name))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if param.Op == "start" {
				s, err := pc.StartVm(c, vmRef)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				c.JSON(200, s)
				return
			} else if param.Op == "shutdown" {
				s, err := pc.ShutdownVm(c, vmRef)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				c.JSON(200, s)
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid operation"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
		}
	}
}

func init() {
	homeCmd.Flags().StringVar(&pveApiUrl, "pve-api-url", pveApiUrl, "api url")
	homeCmd.Flags().StringVar(&pveUser, "pve-user", pveUser, "username")
	homeCmd.Flags().StringVar(&pvePassword, "pve-password", pvePassword, "password")
	homeCmd.Flags().IntVar(&pveTimeout, "pve-timeout", pveTimeout, "timeout in seconds")
}
