package models

type SysUser struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	UserAccount string `json:"user_account"`
	DepID       string `json:"dep_id"`
}

func GetUserListByDepartmentID(deptID string) ([]*SysUser, error) {
	var users []*SysUser
	rows, err := db.Raw("select sys_user.id,sys_user.username,sys_user.user_account,sys_user_depart.dep_id from sys_user left join sys_user_depart on sys_user_depart.user_id=sys_user.id where sys_user_depart.dep_id=?", deptID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, username, userAccount, depId string
		rows.Scan(&id, &username, &userAccount, &depId)
		user := SysUser{
			ID:          id,
			Username:    username,
			UserAccount: userAccount,
			DepID:       depId,
		}
		users = append(users, &user)
	}
	return users, nil
}

func GetUserByMobile(mobile string) (*SysUser, error) {
	var user SysUser
	if err := db.Table("user").Where("mobile=?", mobile).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
