package utils

import (
	"encoding/json"
	"google.golang.org/protobuf/types/known/timestamppb"
	"user/domain/model"
	. "user/proto"
)

// SwapTo 通过json tag 进行结构体赋值
// 此函数主要用于将客户端传来的数据解析到结构体中
// 还用于将数据解析传给客户端
func SwapTo(request, category interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	return json.Unmarshal(dataByte, category)
}

// UserForResponse 类型转化，此函数用于将数据库查出的数据赋值给Response传给客户端
func UserForResponse(response *UserInfoResponse, userModel *model.User) *UserInfoResponse {
	response.UserId = userModel.ID
	response.Username = userModel.UserName
	response.FirstName = userModel.FirstName
	response.LastName = userModel.LastName
	response.Email = userModel.Email
	response.IsActive = userModel.IsActive
	response.Permission = userModel.Permission
	response.CreateDate = timestamppb.New(userModel.CreateDate)
	response.UpdateDate = timestamppb.New(userModel.UpdateDate)
	return response
}
