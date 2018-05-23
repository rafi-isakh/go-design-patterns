// Copyright (c) 2018 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package proxy

import (
	"testing"
	"math/rand"
)

func Test_UserListProxy(t *testing.T) {
	someDatabase := UserList{}

	rand.Seed(2342342)
	for i := 0; i < 1000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase: someDatabase,
		StackCapacity: 2,
		StackCache: UserList{},
	}

	knowdIDs := [3]int32 {someDatabase[3].ID, someDatabase[4].ID, someDatabase[5].ID}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knowdIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knowdIDs[0] {
			t.Error("Returned username doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one succesful search in an empty cache, the size of it must be one")
		}

		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knowdIDs[0])

		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knowdIDs[0] {
			t.Error("Returned username doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an object that is stored on it")
		}

		if !proxy.DidLastSearchUsedCache {
			t.Error("The user should have been returned from the cache")
		}
	})

	t.Run("FindUser - Overflowing the stack", func(t *testing.T) {
		user1, err := proxy.FindUser(knowdIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		user2, _ := proxy.FindUser(knowdIDs[1])
		if proxy.DidLastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		user3, _ :=  proxy.FindUser(knowdIDs[2])
		if proxy.DidLastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		for i := 0; i < len(proxy.StackCache); i++ {
			if proxy.StackCache[i].ID == user1.ID {
				t.Error("User that should be gone was found")
			}
		}

		if len(proxy.StackCache) != 2 {
			t.Error("After inserting 3 users the cache should not grow more than two")
		}

		for _, v := range proxy.StackCache {
			if v != user2 && v != user3 {
				t.Error("A non expected user was found on the cache")
			}
		}
	})
}
