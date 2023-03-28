CREATE TABLE position_players {
	player_id serial NOT NULL,
	player_name varchar(50) NOT NULL,
	team varchar(3),
	position varchar(10) NOT NULL,
	games integer CHECK(games >= 0),
	pa integer CHECK(pa >= 0),
	hr integer CHECK(hr >= 0),
	runs integer CHECK(runs >= 0),
	rbi integer CHECK(rbi >= 0),
	sb integer CHECK(sb >= 0),
	bb_rate float(4) CHECK(bb_rate >= 0),
	k_rate float(4) CHECK(k_rate >= 0),
	iso float(3) CHECK(iso >= 0),
	babip float(3) CHECK(babip >= 0),
	average float(3) CHECK(average >= 0),
	obp float(3) CHECK(obp >= 0),
	slg float(3) CHECK(slg >= 0),
	woba float(3) CHECK(woba >= 0),
	wrc_plus integer CHECK(wrc_plus >= 0),
	war float(3),
	PRIMARY KEY (player_id),
}

CREATE TABLE teams {
	team_id serial NOT NULL,
	team varchar(3) NOT NULL,
	games integer CHECK(games >= 0),
	pa integer CHECK(pa >= 0),
	hr integer CHECK(hr >= 0),
	runs integer CHECK(runs >= 0),
	rbi integer CHECK(rbi >= 0),
	sb integer CHECK(sb >= 0),
	bb_rate float(4) CHECK(bb_rate >= 0),
	k_rate float(4) CHECK(k_rate >= 0),
	iso float(3) CHECK(iso >= 0),
	babip float(3) CHECK(babip >= 0),
	average float(3) CHECK(average >= 0),
	obp float(3) CHECK(obp >= 0),
	slg float(3) CHECK(slg >= 0),
	woba float(3) CHECK(woba >= 0),
	wrc_plus integer CHECK(wrc_plus >= 0),
	war float(3),
	PRIMARY KEY (team_id),
}
