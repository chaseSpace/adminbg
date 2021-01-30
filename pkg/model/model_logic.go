package model

func IsReservedMenuId(mfID int32) bool {
	return mfID == MenuRootId || mfID == MenuDefaultId
}
