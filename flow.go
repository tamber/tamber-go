package tamber

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Block struct {
	Name      string `json:"name"`
	BlockType string `json:"block_type"`
	Id        uint32 `json:"id"`
	ProjectId uint32 `json:"project_id"`
}

type BlockResponse struct {
	TamberResponse `json:",inline"`
	Result         Block `json:"result"`
}

type CreateBlockParams struct {
	Block interface{} `json:"block"`
}

type BlockKey struct {
	Id   uint32 `json:"id"`
	Type string `json:"type"`
}

type RemoveBlockParams struct {
	Key    BlockKey `json:"key"`
	DryRun bool     `json:"dry_run"`
}

func (params *CreateBlockParams) AppendToBody(v *url.Values) {
	block, _ := json.Marshal(params.Block)
	v.Add("block", string(block))
}

func (params *RemoveBlockParams) AppendToBody(v *url.Values) {
	key, _ := json.Marshal(params.Key)
	v.Add("key", string(key))
	v.Add("dry_run", strconv.FormatBool(params.DryRun))
}
