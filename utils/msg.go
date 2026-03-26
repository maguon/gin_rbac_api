package utils

func GetUserFavorPosMsg(name string) map[string]interface{} {
	objInterface := make(map[string]interface{})
	objInterface["title"] = "职位关注"
	objInterface["content"] = "用户" + name + "关注了您的职位"
	return objInterface
}

func GetUserFavorBizMsg(name string) map[string]interface{} {
	objInterface := make(map[string]interface{})
	objInterface["title"] = "企业关注"
	objInterface["content"] = "用户" + name + "关注了您的公司"
	return objInterface
}

func GetUserSendPosMsg(userName string) map[string]interface{} {
	objInterface := make(map[string]interface{})
	objInterface["title"] = "收到简历"
	objInterface["content"] = "用户" + userName + "向您投递了简历"
	return objInterface
}

func GetBizViewUserMsg(bizName string) map[string]interface{} {
	objInterface := make(map[string]interface{})
	objInterface["title"] = "查看简历"
	objInterface["content"] = bizName + "查看了您投递了简历"
	return objInterface
}

func GetBizFavorUserMsg(bizName string) map[string]interface{} {
	objInterface := make(map[string]interface{})
	objInterface["title"] = "关注简历"
	objInterface["content"] = bizName + "关注了您的简历"
	return objInterface
}

func GetBizPosViewUserMsg(bizName string) map[string]interface{} {
	objInterface := make(map[string]interface{})
	objInterface["title"] = "邀请面试"
	objInterface["content"] = bizName + "邀请您参与面试"
	return objInterface
}
