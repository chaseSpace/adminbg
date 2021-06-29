package common

/*
buz.go: business common logic func
*/

func IsSuperAdmin(uid int32) (bool, error) {
	return uid == 1, nil
}
