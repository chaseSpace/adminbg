package common

func IsSuperAdmin(uid int32) (bool, error) {
	// More custom logic can be added here
	return uid == 1, nil
}
