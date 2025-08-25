package gd

type (
	UsersParams struct {
		Secret string `query:"secret"`
		Query  string `query:"str"`
	}
	UserInfoParams struct {
		Secret      string `query:"secret"`
		TargetAccID string `query:"targetAccountID"`
	}
	DailyParams struct {
		Secret string `query:"secret"`
		// Get Weekly?
		//	0 - false
		//	1 - true
		Weekly int `query:"weekly"`
	}
	DownloadGJLevel22Params struct{}

)
