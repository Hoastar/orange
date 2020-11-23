/*
@Time : 2020/11/8 下午5:39
@Author : hoastar
@File : getState
@Software: GoLand
*/

package service

type ProcessState struct {
	Structure map[string][]map[string]interface{}
}

// 获取节点信息
func (p *ProcessState) GetNode(stateId string) (nodeValue map[string]interface{}, err error) {
	for _, node := range p.Structure["nodes"] {
		if node["id"] == stateId {
			nodeValue = node
		}
	}

	return
}