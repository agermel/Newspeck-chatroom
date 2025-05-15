package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"newspeak-chat/internal/svc"
	"newspeak-chat/internal/types"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/zeromicro/go-zero/core/logx"
)

type FilterMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterMessageLogic {
	return &FilterMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterMessageLogic) FilterMessage(req *types.FilterRequest) (resp *types.FilterResponse, err error) {
	// 大模型定义
	llm, err := openai.New()
	if err != nil {
		return nil, err
	}

	prompt := fmt.Sprintf(`
		You are a language censorship AI from George Orwell's *1984*. Your job is to review a sentence and return a strict JSON response.

		Rules:
		- Rewrite the sentence using Newspeak or equivalent Orwellian terms.
		- Determine if the message contains thoughtcrime (ideas against the Party).
		- Identify dangerous words or phrases ("triggers").
		- Set the danger level: "low", "medium", or "high".
		- The notes is Big Brother's warning.
		- Return a short note.

		Respond ONLY in this exact JSON format:

		{
		"original": "<original sentence>",
		"filtered": "<rewritten sentence>",
		"danger_level": "<low|medium|high>",
		"triggers": ["<list of flagged words>"],
		"note": "<short explanation>"
		}

		Sentence to review: "%s"
	`, req.Message)

	//  使用 LLM 处理消息
	output, err := llm.Call(l.ctx, prompt, llms.WithTemperature(0.2))
	if err != nil {
		return nil, err
	}

	var result types.FilterResponse
	err = json.Unmarshal([]byte(output), &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}
