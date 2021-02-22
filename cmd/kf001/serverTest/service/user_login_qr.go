package service

/*
// 添加用户二维码扫描登录
func UserLoginQrSave(req mongoSql.UserLoginQr) (res mongoSql.UserLogin, err error) {
	if count, _ := mongo.Collection(&req).Where(bson.M{"_id": req.ID}).Count(); count == 0 {
		result, err := mongo.Collection(&req).InsertOne(&req)
		if err == nil {
			res.ID = result.InsertedID.(string)
		}
	} else {
		_, err = mongo.Collection(&req).Where(bson.M{"_id": req.ID}).UpdateOne(req)
	}
	return res, err
}

// 删除二维码
func UserLoginQrData(req mongoSql.UserLoginQr) (res mongoSql.UserLogin, err error) {
	if count, _ := mongo.Collection(&req).Where(bson.M{"_id": req.ID}).Count(); count == 0 {
		result, err := mongo.Collection(&req).InsertOne(&req)
		if err == nil {
			res.ID = result.InsertedID.(string)
		}
	} else {
		_, err = mongo.Collection(&req).Where(bson.M{"_id": req.ID}).UpdateOne(req)
	}
	return res, err
}

 */