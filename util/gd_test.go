package util

import (
	"io"
	"testing"
)

func TestGetGDUserData(t *testing.T) {
	req := UsersParams{
		Secret: COMMON_KEY,
		Query:  "dexslender",
	}

	res, err := GDClient.Request(Users, req)
	if err != nil {
		t.Fatal(res)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	// 1:RobTop:2:16:13:20:17:175:6:0:9:483:10:35:11:2:14:0:15:0:16:71:3:2774:52:117:8:0:4:5#999:0:10
	user := DecodeData[PartialUser](string(data))
	t.Logf("%+v", user)
}

func TestGetUserInfo(t *testing.T) {
	req := UserInfoParams{ COMMON_KEY, "11437853" }
	res, err := GDClient.Request(UserInfo, req)
	if err != nil { t.Fatal(err) }
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil { t.Fatal(err) }

	// user := DecodeData[User](string(data))
	t.Logf(string(data))
}
