package image

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"io/ioutil"
	"os"
	"strings"
)

func NewClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return cli
}

func ScanImageLabels(f []string) {
	if len(f) != 3 {
		_, err := fmt.Fprintln(os.Stderr, "filter长度不能小于3")
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}
	labelKey := f[0]
	labelValue := f[1]
	tagFilter := f[2]
	cli := NewClient()
	// 按标签匹配镜像
	list, err := cli.ImageList(context.TODO(), types.ImageListOptions{
		Filters: filters.NewArgs(filters.Arg("label", labelKey+"="+labelValue)),
	})
	if err != nil {
		return
	}
	var svcRevision = "本次交付的微服务分支及git revision如下\n"
	for _, v := range list {
		//fmt.Println(v)
		// 按tag过滤镜像
		if !strings.Contains(v.RepoTags[0], tagFilter) {
			continue
		}
		line := v.Labels["title"] + "\t" + v.Labels["source"] + "\t" + v.Labels["revision"] + "\n"
		fmt.Println(line)
		svcRevision += line
	}
	saveToFile(svcRevision, "release.txt")
}

func saveToFile(d string, n string) {
	err := ioutil.WriteFile(n, []byte(d), 0644)
	if err != nil {
		panic(err)
	}
}
