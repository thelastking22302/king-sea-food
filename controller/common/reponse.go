package common

type Reponse struct {
	Data         interface{} `json:"data"`
	Paggings     interface{} `json:"paggings"`
	AccessToken  interface{} `json:"access_token"`
	RefreshToken interface{} `json:"refresh_token"`
}

func ReponseDataToken(data, accesstoken, refreshToken interface{}) *Reponse {
	return &Reponse{Data: data, Paggings: nil, AccessToken: accesstoken, RefreshToken: refreshToken}
}
func MutiResponse(data, paggings, accessToken, refreshToken interface{}) *Reponse {
	return &Reponse{Data: data, Paggings: paggings, AccessToken: accessToken, RefreshToken: refreshToken}
}
func ReponseData(data interface{}) *Reponse {
	return &Reponse{Data: data, Paggings: nil, AccessToken: nil, RefreshToken: nil}
}
