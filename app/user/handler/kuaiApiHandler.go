package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CubeBaseResponse struct {
	Errno  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
}

type KuaiResultResponse struct {
	CubeBaseResponse
	Data struct {
		TopicList []Topic `json:"topicList"`
	} `json:"data,omitempty"`
}

type Topic struct {
	TopicTitle string `json:"topic"`
	Ranking    string `json:"ranking"`
	Heat       string `json:"heat"`
	HotTag     string `json:"hot_tag"`
	PCurl      string `json:"pcurl"`
}

func KuaiApitest(ctx *gin.Context) {
	k := new(KuaiResultResponse)
	for i := 0; i < 15; i++ {
		k.Data.TopicList = append(k.Data.TopicList, Topic{
			TopicTitle: fmt.Sprintf("%d", i),
			Ranking:    "",
			Heat:       "",
			HotTag:     "",
			PCurl:      fmt.Sprintf("http://www.baibai.com/hell0?h=%d", i),
		})
	}
	ctx.AbortWithStatusJSON(0, k)
}
