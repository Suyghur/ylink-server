//@File     csinfo.go
//@Time     2022/05/12
//@Author   #Suyghur,

package model

type CsInfo struct {
	CsId         string `json:"cs_id"`
	CsNickname   string `json:"cs_nickname"`
	CsAvatarUrl  string `json:"cs_avatar_url"`
	CsSignature  string `json:"cs_signature"`
	OnlineStatus int64  `json:"online_status"`
}
