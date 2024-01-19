package k8s

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
)

func CleanNodeLabels(keywords []string, erase bool) {

	nodeList, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, n := range nodeList.Items {
		fmt.Printf("%s found labels matched filters:\n", n.Name)

		// init empty metadata with empty label
		var patchData = map[string]interface{}{}
		patchData["metadata"] = map[string]interface{}{"labels": map[string]interface{}{}}

		for k := range n.Labels {
			if match(k, keywords) {
				fmt.Printf("\t%s\n", k)
				// kubernetes strategy patch requires payload as follows to delete object
				patchData["metadata"].(map[string]interface{})["labels"].(map[string]interface{})[k] = nil
			}
		}
		if erase {
			payload, _ := json.Marshal(patchData)
			//fmt.Printf("\t%s\n", payload)
			_, err := client.CoreV1().Nodes().Patch(context.TODO(), n.Name, types.StrategicMergePatchType, payload, metav1.PatchOptions{})
			if err != nil {
				panic(err)
			}
		}
	}
}

func match(labelKey string, keywords []string) bool {
	for _, k := range keywords {
		if !strings.Contains(labelKey, k) {
			return false
		}
	}
	return true
}
