package helpers



func MatchUserTypeToUid(c *gin.Context, userId string) (err error){
	userType := c.GetString("user_type")
	uid:= c.GetString("uid")
	err nil
	
}