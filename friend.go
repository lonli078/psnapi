package psnapi


const (
	users_url = "https://us-prof.np.community.playstation.net/userProfile/v1/users/"
)

func (a *Api) Get_my_info() (map[string]interface{}, error){
	url2 := "me/profile2?"
	m := make(map[string]string)
	m["fields"] = "npId,onlineId,avatarUrls,plus,aboutMe,languagesUsed,trophySummary(@default,progress,earnedTrophies),isOfficiallyVerified,personalDetail(@default,profilePictureUrls),personalDetailSharing,personalDetailSharingRequestMessageFlag,primaryOnlineStatus,presences(@titleInfo,hasBroadcastData),friendRelation,requestMessageFlag,blocking,mutualFriendsCount,following,followerCount,friendsCount,followingUsersCount"
	m["avatarSizes"] = "m"
	m["profilePictureSizes"] = "m"
	m["languagesUsedLanguageSet"] = "set3"
	m["psVitaTitleIcon"] = "circled"
	m["titleIconSize"] = "s"
	data := make(map[string]interface{})
	err := a.make_get_request(users_url+url2+a.make_params(m), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Api) Get_friends(limit, offset string) (map[string]interface{}, error){
	url2 := "me/friends/profiles2?"
	m := make(map[string]string)
	m["fields"] = "onlineId,avatarUrls,following,friendRelation,isOfficiallyVerified,personalDetail(@default,profilePictureUrls),personalDetailSharing,plus,presences(@titleInfo,hasBroadcastData,lastOnlineDate),primaryOnlineStatus,trophySummary(@default)"
	m["sort"] = "name-onlineId"
	//m["userFilter"] = "online"
	m["avatarSizes"] = "m"
	m["profilePictureSizes"] = "m"
	m["offset"] = offset
	m["limit"] = limit
	data := make(map[string]interface{})
	err := a.make_get_request(users_url+url2+a.make_params(m), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Api) Get_info(psn_id string) (map[string]interface{}, error){
	url2 := "/profile2?"
	m := make(map[string]string)
	m["fields"] = "npId,onlineId,avatarUrls,plus,aboutMe,languagesUsed,trophySummary(@default,progress,earnedTrophies),isOfficiallyVerified,personalDetail(@default,profilePictureUrls),personalDetailSharing,personalDetailSharingRequestMessageFlag,primaryOnlineStatus,presences(@titleInfo,hasBroadcastData),friendRelation,requestMessageFlag,blocking,mutualFriendsCount,following,followerCount,friendsCount,followingUsersCount"
	m["avatarSizes"] = "m"
	m["profilePictureSizes"] = "m"
	m["languagesUsedLanguageSet"] = "set3"
	m["psVitaTitleIcon"] = "circled"
	m["titleIconSize"] = "s"
	m["avatarSizes"] = "m"
	data := make(map[string]interface{})
	err := a.make_get_request(users_url+psn_id+url2+a.make_params(m), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Api) Get_friends_of_friend(psn_id, limit, offsset string) (map[string]interface{}, error){
	url2 := "/friends/profiles2?"
	m := make(map[string]string)
	m["fields"] = "onlineId,avatarUrls,plus,trophySummary(@default),isOfficiallyVerified,personalDetail(@default,profilePictureUrls),primaryOnlineStatus,presences(@titleInfo,hasBroadcastData)"
	m["sort"] = "name-onlineId"
	m["avatarSizes"] = "m"
	m["profilePictureSizes"] = "m"
	m["extendPersonalDetailTarget"] = "true"
	m["offset"] = offsset
	m["limit"] = limit
	data := make(map[string]interface{})
	err := a.make_get_request(users_url+psn_id+url2+a.make_params(m), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
