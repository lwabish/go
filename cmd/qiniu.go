/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/qiniu/go-sdk/v7/cdn"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// qiniuCmd represents the qiniu command
var qiniuCmd = &cobra.Command{
	Use:   "qiniu",
	Short: "Manipulate qiniu",
	Run: func(cmd *cobra.Command, args []string) {
		checkCert()
	},
}

var (
	qiniuCdnCertName = "auto"
	keyPath          = "key.pem"
	certPath         = "cert.pem"
	accessKey        = ""
	secretKey        = ""
)

func checkCert() {
	mac := credentials.NewCredentials(accessKey, secretKey)
	cm := cdn.NewCdnManager(mac)

	// fixme: 校验status code
	res, err := cm.GetCertList("", 9999)
	if err != nil {
		log.Fatalln(err)
	}
	if len(res.Certs) == 0 {
		log.Fatalln(fmt.Errorf("no certs found"))
	}
	c, found := lo.Find(res.Certs, func(item struct {
		CertID     string   `json:"certid"`
		Name       string   `json:"name"`
		CommonName string   `json:"common_name"`
		DNSNames   []string `json:"dnsnames"`
		NotBefore  int      `json:"not_before"`
		NotAfter   int      `json:"not_after"`
		CreateTime int      `json:"create_time"`
	}) bool {
		return item.Name == qiniuCdnCertName
	})
	if !found {
		log.Fatalln(fmt.Errorf("no certs found"))
	}
	dateExpire := time.Unix(int64(c.NotAfter), 0)
	delta := dateExpire.Sub(time.Now())
	log.Printf("certs will expire in %d days\n", delta/(time.Hour*24))
	if delta > time.Hour*24*30 {
		log.Printf("certs delta > 30 days, skip")
		return
	}

	uploadCert(cm)

	// todo: 更新域名的证书(qiniu库还没实现)，然后删除老的证书
}

func uploadCert(cm *cdn.CdnManager) {
	certBs, err := os.ReadFile(certPath)
	if err != nil {
		log.Fatalln(err)
	}
	keyBs, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := cm.UploadCert(qiniuCdnCertName, "wubw.fun", string(keyBs), string(certBs))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("cert id: %s\n", res.CertID)
}

func init() {
	rootCmd.AddCommand(qiniuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qiniuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	qiniuCmd.Flags().StringVar(&qiniuCdnCertName, "cert-name", qiniuCdnCertName, "cert name")
	qiniuCmd.Flags().StringVar(&certPath, "cert-path", certPath, "cert path")
	qiniuCmd.Flags().StringVar(&keyPath, "key-path", keyPath, "key path")
	qiniuCmd.Flags().StringVar(&accessKey, "access-key", accessKey, "access key")
	qiniuCmd.Flags().StringVar(&secretKey, "secret-key", secretKey, "secret key")
}
