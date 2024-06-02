package cmd

import (
	"awesomeProject2/util"
	"log"
)

func ParseDestiny() {
	destiny := &util.DestinyMetrics{}
	if err := util.GetOutputJson(destiny); err != nil {
		log.Fatal(err)
	}
	//p, err := json.Marshal(destiny.PodStartupLatency)
	//sp, err := json.Marshal(destiny.StatelessPodStartupLatency)
	//p = append(p, sp...)

}

//func sendResult() {
//	urlPath := os.Getenv("URL_PATH")
//	// http://10.70.12.130:31989/openapi/v1/{cluster_deploy_id}/testing/{testing_id}
//	req, err := http.NewRequest("POST", urlPath+fmt.Sprintf("/openapi/v1/%s/testing/%s?status=%d&type_id=%s", v.ClusterDeployID, v.TestID, 1, v.TypeID), fileContent)
//	fmt.Println("返回测试结果的地址为:" + urlPath + fmt.Sprintf("/openapi/v1/%s/testing/%s?status=%d&type_id=%s", v.ClusterDeployID, v.TestID, 1, v.TypeID))
//
//	if err != nil {
//		log.Print(err)
//	}
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//	req.Header.Set("Authorization", "Internal testing")
//
//	httpClient := &http.Client{}
//	resp, err := httpClient.Do(req)
//	if err != nil {
//		log.Print(err)
//	}
//	defer resp.Body.Close()
//
//	// 处理响应
//	if resp.StatusCode == http.StatusNotFound {
//		log.Print("Request failed with status code: " + resp.Status)
//	}
//}
