package server

import (
	"crypto/tls"
	"github.com/samber/lo"
	"log"
	"net/http"

	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	pveApiUrl   = ""
	pveUser     = ""
	pvePassword = ""
	pveTimeout  = 300
)

var homeCmd = &cobra.Command{
	Use: "home",
	Run: func(cmd *cobra.Command, args []string) {

		pveClient, err := proxmox.NewClient(pveApiUrl, nil, "", &tls.Config{InsecureSkipVerify: true}, "", pveTimeout)
		if err != nil {
			log.Fatalln(err)
		}
		pveClient.SetAPIToken(pveUser, pvePassword)

		r := gin.Default()
		r.Group("api/v1")
		r.Any("/node", pc(pveClient))
		r.Any("/vm", vm(pveClient))
		if err := r.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	},
}

func pc(pc *proxmox.Client) gin.HandlerFunc {
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
			}{}
			if err := c.ShouldBindBodyWithJSON(&param); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if param.Op == "start" {
				//TODO: https://github.com/sabhiram/go-wol
			} else if param.Op == "shutdown" {
				s, err := pc.ShutdownNode(c, param.Name)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				c.JSON(200, lo.Ternary(s == "", "OK", s))
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
