package models


// *************
// Position Player Model
// *************

type PositionPlayer struct {
	// Player id (auto incremented by database)
	ID		int     `json:"id"`
	Name		string  `json:"name"`
	Team		string  `json:"team"`
	G		int     `json:"games"`
	PA		int     `json:"plateAppearances"`
	HR		int     `json:"homeRuns"`
	R		int     `json:"runs"`
	RBI		int     `json:"runsBattedIn"`
	SB		int     `json:"stolenBases"`
	WRCPlus		int	`json:"weightedRunsCreatedPlus"`
	BbRate		float64 `json:"walkRate"`
	KRate		float64 `json:"strikeoutRate"`
	ISO		float64 `json:"isolatedPower"`
	BABIP		float64 `json:"battingAvgBallsInPlay"`
	AVG		float64 `json:"battingAvg"`
	OBP		float64 `json:"onBasePct"`
	SLG		float64 `json:"sluggingPct"`
	WOBA		float64 `json:"weightedOnBaseAvg"`
	XWOBA		float64 `json:"expWeightedOnBaseAvg"`
	BsR		float64 `json:"baseRunning"`
	WAR		float64 `json:"winsAboveReplacement"`

}

type CreatePositionPlayerRequest struct {
	Name		string  `json:"name"`
	Team		string  `json:"team"`
	G		int     `json:"games"`
	PA		int     `json:"plateAppearances"`
	HR		int     `json:"homeRuns"`
	R		int     `json:"runs"`
	RBI		int     `json:"runsBattedIn"`
	SB		int     `json:"stolenBases"`
	WRCPlus		int	`json:"weightedRunsCreatedPlus"`
	BbRate		float64 `json:"walkRate"`
	KRate		float64 `json:"strikeoutRate"`
	ISO		float64 `json:"isolatedPower"`
	BABIP		float64 `json:"battingAvgBallsInPlay"`
	AVG		float64 `json:"battingAvg"`
	OBP		float64 `json:"onBasePct"`
	SLG		float64 `json:"sluggingPct"`
	WOBA		float64 `json:"weightedOnBaseAvg"`
	XWOBA		float64 `json:"expWeightedOnBaseAvg"`
	BsR		float64 `json:"baseRunning"`
	WAR		float64 `json:"winsAboveReplacement"`
}

type UpdatePositionPlayerRequest struct {
	Name		string  `json:"name"`
	Team		string  `json:"team"`
	G		int     `json:"games"`
	PA		int     `json:"plateAppearances"`
	HR		int     `json:"homeRuns"`
	R		int     `json:"runs"`
	RBI		int     `json:"runsBattedIn"`
	SB		int     `json:"stolenBases"`
	WRCPlus		int	`json:"weightedRunsCreatedPlus"`
	BbRate		float64 `json:"walkRate"`
	KRate		float64 `json:"strikeoutRate"`
	ISO		float64 `json:"isolatedPower"`
	BABIP		float64 `json:"battingAvgBallsInPlay"`
	AVG		float64 `json:"battingAvg"`
	OBP		float64 `json:"onBasePct"`
	SLG		float64 `json:"sluggingPct"`
	WOBA		float64 `json:"weightedOnBaseAvg"`
	XWOBA		float64 `json:"expWeightedOnBaseAvg"`
	BsR		float64 `json:"baseRunning"`
	WAR		float64 `json:"winsAboveReplacement"`
}

type UpdatedPositionPlayer struct {
	UpdatedMap map[string]int
}

type DeletedPositionPlayer struct {
	DeletedMap map[string]int
}

func NewPositionPlayer(name, team string, g, pa, hr, r, rbi, sb, wrcPlus int, bbRate, kRate, iso, babip, avg, obp, slg, woba, xWoba, bsr, war float64) *PositionPlayer {
	return &PositionPlayer{
		Name:		name,
		Team:		team,
		G:		g,
		PA:		pa,
		HR:		hr,
		R:		r,
		RBI:		rbi,
		SB:		sb,
		WRCPlus:	wrcPlus,
		BbRate:		bbRate,
		KRate:		kRate,
		ISO:		iso,
		BABIP:		babip,
		AVG:		avg,
		OBP:		obp,
		SLG:		slg,
		WOBA:		woba,
		XWOBA:		xWoba,
		BsR:		bsr,
		WAR:		war,
	}
}

// *************
// Pitcher Model
// *************

type Pitcher struct {
	// Player id (auto incremented by database)
	ID		int     `json:"id"`
	Name		string  `json:"name"`
	Team		string  `json:"team"`
	W		int	`json:"wins"`
	L		int	`json:"losses"`
	SV		int	`json:"saves"`
	G		int     `json:"games"`
	GS		int	`json:"gamesSaved"`
	IP		float64	`json:"inningsPitched"`
	K9		float64	`json:"strikeoutsPerNine"`
	BB9		float64	`json:"walksPerNine"`
	HR9		float64	`json:"homeRunsPerNine"`
	BABIP		float64	`json:"battingAvgBallsInPlay"`
	LOB		float64	`json:"leftOnBase"`
	GB		float64	`json:"groundballRate"`
	HRFB		float64	`json:"homeRunToFlyBallRatio"`
	VFA		float64	`json:"fourseamFastballVelocity"`
	ERA		float64	`json:"earnedRunAvg"`
	XERA		float64	`json:"expectedEarnedRunAvg"`
	FIP		float64	`json:"fielderIndependentPitching"`
	XFIP		float64	`json:"expectedFielderIndependentPitching"`
	WAR		float64 `json:"winsAboveReplacement"`
}

type CreatePitcherRequest struct {
	Name		string  `json:"name"`
	Team		string  `json:"team"`
	W		int	`json:"wins"`
	L		int	`json:"losses"`
	SV		int	`json:"saves"`
	G		int     `json:"games"`
	GS		int	`json:"gamesSaved"`
	IP		float64	`json:"inningsPitched"`
	K9		float64	`json:"strikeoutsPerNine"`
	BB9		float64	`json:"walksPerNine"`
	HR9		float64	`json:"homeRunsPerNine"`
	BABIP		float64	`json:"battingAvgBallsInPlay"`
	LOB		float64	`json:"leftOnBase"`
	GB		float64	`json:"groundballRate"`
	HRFB		float64	`json:"homeRunToFlyBallRatio"`
	VFA		float64	`json:"fourseamFastballVelocity"`
	ERA		float64	`json:"earnedRunAvg"`
	XERA		float64	`json:"expectedEarnedRunAvg"`
	FIP		float64	`json:"fielderIndependentPitching"`
	XFIP		float64	`json:"expectedFielderIndependentPitching"`
	WAR		float64 `json:"winsAboveReplacement"`
}

type UpdatePitcherRequest struct {
	Name		string  `json:"name"`
	Team		string  `json:"team"`
	W		int	`json:"wins"`
	L		int	`json:"losses"`
	SV		int	`json:"saves"`
	G		int     `json:"games"`
	GS		int	`json:"gamesSaved"`
	IP		float64	`json:"inningsPitched"`
	K9		float64	`json:"strikeoutsPerNine"`
	BB9		float64	`json:"walksPerNine"`
	HR9		float64	`json:"homeRunsPerNine"`
	BABIP		float64	`json:"battingAvgBallsInPlay"`
	LOB		float64	`json:"leftOnBase"`
	GB		float64	`json:"groundballRate"`
	HRFB		float64	`json:"homeRunToFlyBallRatio"`
	VFA		float64	`json:"fourseamFastballVelocity"`
	ERA		float64	`json:"earnedRunAvg"`
	XERA		float64	`json:"expectedEarnedRunAvg"`
	FIP		float64	`json:"fielderIndependentPitching"`
	XFIP		float64	`json:"expectedFielderIndependentPitching"`
	WAR		float64 `json:"winsAboveReplacement"`
}

type UpdatedPitcher struct {
	UpdatedMap map[string]int
}

type DeletedPitcher struct {
	DeletedMap map[string]int
}

func NewPitcher(name, team string, w, l, sv, g, gs int, ip, k9, bb9, hr9, babip, lob, gb, hrfb, vfa, era, xera, fip, xfip, war float64) *Pitcher {
	return &Pitcher{
		Name:		name,
		Team:		team,
		W:		w,
		L:		l,
		SV:		sv,
		G:		g,
		GS:		gs,
		IP:		ip,
		K9:		k9,
		BB9:		bb9,
		HR9:		hr9,
		BABIP:		babip,
		LOB:		lob,
		GB:		gb,
		HRFB:		hrfb,
		VFA:		vfa,
		ERA:		era,
		XERA:		xera,
		FIP:		fip,
		XFIP:		xfip,
		WAR:		war,
	}
}
