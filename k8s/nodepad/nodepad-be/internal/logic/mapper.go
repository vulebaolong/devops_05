package logic

import (
	"nodepad-be/ent"
	"nodepad-be/internal/common/helper"
	"nodepad-be/internal/types"
)

func NoteData(item *ent.Note) types.NoteData {
	var userID *string
	if item.UserID != nil {
		value := item.UserID.String()
		userID = &value
	}
	return types.NoteData{
		Id: item.ID.String(), UserId: userID, ShareKey: item.ShareKey,
		ShareUrl: "/notes/public/" + item.ShareKey,
		Title:    item.Title, Content: item.Content,
		CreatedAt: helper.TimeToRFC3339(item.CreatedAt),
		UpdatedAt: helper.TimeToRFC3339(item.UpdatedAt),
		DeletedAt: helper.OptionalTimeToRFC3339(item.DeletedAt),
	}
}

func AuthUser(item *ent.User) types.AuthUser {
	return types.AuthUser{
		Id: item.ID.String(), Email: item.Email,
		CreatedAt: helper.TimeToRFC3339(item.CreatedAt),
		UpdatedAt: helper.TimeToRFC3339(item.UpdatedAt),
		DeletedAt: helper.OptionalTimeToRFC3339(item.DeletedAt),
	}
}
