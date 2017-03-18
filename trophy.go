package psnapi

const (
	trophy_url = "https://us-tpy.np.community.playstation.net/trophy/v1/trophyTitles/"
)

func (a *Api) GetTrophies(limit, offset, npLanguage, platform, comparedUser string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["fields"] = "@default"
	m["npLanguage"] = npLanguage
	m["offset"] = offset
	m["limit"] = limit
	if comparedUser != "" {
		m["comparedUser"] = comparedUser
	}
	if platform == "" {
		m["platform"] = "PS3,PSVITA,PS4"
	}
	data := make(map[string]interface{})
	err := a.make_get_request(trophy_url+"?"+a.make_params(m), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Api) GetGameTrophies(npCommunicationId, group, npLanguage, comparedUser string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["fields"] = "@default,trophyRare,trophyEarnedRate"
	m["npLanguage"] = npLanguage
	if comparedUser != "" {
		m["comparedUser"] = comparedUser
	}
	url2 := "/trophyGroups/"+group+"/trophies?"
	data := make(map[string]interface{})
	err := a.make_get_request(trophy_url+npCommunicationId+url2+a.make_params(m), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Api) GetGameGroupList(npCommunicationId, npLanguage string)  (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := a.make_get_request(trophy_url+npCommunicationId+"/trophyGroups/?npLanguage="+npLanguage, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Api) GetTrophyInfo(npCommunicationId, group, trophyId, npLanguage string)(map[string]interface{}, error) {
	m := make(map[string]string)
	m["npLanguage"] = npLanguage
	m["fields"] = "@default,trophyRare,trophyEarnedRate"
	data := make(map[string]interface{})
	err := a.make_get_request(trophy_url+npCommunicationId+"/trophyGroups/"+group+"/trophies/"+trophyId+"/?"+a.make_params(m), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

