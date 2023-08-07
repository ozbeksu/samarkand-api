package main

import (
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/types"
)

func createAdmin(aID, cID int) *ent.User {
	pass, _ := types.HashPassword("secret")
	return db.User.Create().
		SetUsername("admin").
		SetEmail("admin@samarkand-api.com").
		SetPassword(pass).
		SetActive(true).
		AddAttachmentIDs(aID, cID).
		SaveX(ctx)
}

func createUser(aID, cID int) *ent.User {
	return db.User.Create().
		SetUsername(faker.Username()).
		SetEmail(faker.Email()).
		SetPassword("secret").
		SetActive(faker.Bool()).
		AddAttachmentIDs(aID, cID).
		SaveX(ctx)
}

func createProfile(userID, avatarID, coverID int) *ent.Profile {
	return db.Profile.Create().
		SetFirstName(faker.FirstName()).
		SetLastName(faker.LastName()).
		SetUserID(userID).
		SetAvatarID(avatarID).
		SetCoverID(coverID).
		SaveX(ctx)
}

func makeUsers(n int) []*ent.User {
	var u *ent.User
	var users []*ent.User

	for i := 1; i < n; i++ {
		a := createImage(i, "avatars", "jpg", 200, 200).SaveX(ctx)
		c := createImage(i, "covers", "jpg", 800, 392).SaveX(ctx)

		if i == 1 {
			u = createAdmin(a.ID, c.ID)
		} else {
			u = createUser(a.ID, c.ID)
		}
		createProfile(u.ID, a.ID, c.ID)

		users = append(users, u)
	}

	return users
}

func addFollowers(users []*ent.User, n int) {
	for _, user := range users {
		IDs := []int{getRandIntInRange(1, n), getRandIntInRange(1, n), getRandIntInRange(1, n)}
		user.Update().AddFollowingIDs(IDs...).SaveX(ctx)
	}
}
