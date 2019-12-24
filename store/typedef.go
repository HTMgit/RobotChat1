package store

type KeyWordData struct {
	KeyID     string `db:"keyid"`
	Backworld string `db:"back_world"`
}

type ExchangeWordData struct {
	KeyID   string `db:"keyid"`
	Exworld string `db:"exchange_world"`
}

type UserInfoData struct {
	Uid         string `db:"uid"`
	RobotStatus int    `db:"robot_status"`
	GreetStatus int    `db:"greet_status"`
}
