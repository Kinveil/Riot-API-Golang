package apiclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/apiclient/ratelimiter"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/continent"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/patch"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/queue"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/region"
	"github.com/Kinveil-Engineering-Analysis/Riot-API-Golang/constants/summoner_spell"
)

// Matchlist is an array of strings that represent the match IDs.
type Matchlist []string

func (m Matchlist) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Matchlist) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

type GetMatchlistOptions struct {
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Queue     *queue.ID  `json:"queue"`
	Type      *string    `json:"type"`
	Start     *int       `json:"start"`
	Count     *int       `json:"count"`
}

func (c *client) GetMatchlist(continent continent.Continent, puuid string, opts *GetMatchlistOptions) (*Matchlist, error) {
	var params url.Values = make(map[string][]string)

	if opts != nil {
		if opts.StartTime != nil {
			params.Add("startTime", fmt.Sprintf("%d", opts.StartTime.UnixNano()/int64(time.Second)))
		}

		if opts.EndTime != nil {
			params.Add("endTime", fmt.Sprintf("%d", opts.EndTime.UnixNano()/int64(time.Second)))
		}

		if opts.Queue != nil {
			params.Add("queue", fmt.Sprintf("%d", *opts.Queue))
		}

		if opts.Type != nil {
			params.Add("type", *opts.Type)
		}

		if opts.Start != nil {
			params.Add("start", fmt.Sprintf("%d", *opts.Start))
		}

		if opts.Count != nil {
			params.Add("count", fmt.Sprintf("%d", *opts.Count))
		}
	}

	var res Matchlist
	_, err := c.dispatchAndUnmarshal(c.ctx, continent, "/lol/match/v5/matches/by-puuid", fmt.Sprintf("/%s/ids", puuid), params, ratelimiter.GetMatchlist, &res)
	return &res, err
}

type Match struct {
	Metadata MatchMetadata `json:"metadata"`
	Info     MatchInfo     `json:"info"`
}

type MatchMetadata struct {
	DataVersion  string   `json:"dataVersion"` // ex: 2
	MatchID      string   `json:"matchId"`     // ex: NA1_1234567890
	Participants []string `json:"participants"`
}

type MatchInfo struct {
	GameCreation       int                    `json:"gameCreation"`       // ex: 1483140696030
	GameDuration       int                    `json:"gameDuration"`       // ex: 1561
	GameEndTimestamp   int                    `json:"gameEndTimestamp"`   // ex: 1483142251591
	GameID             int                    `json:"gameId"`             // ex: 1234567890
	GameMode           string                 `json:"gameMode"`           // ex: CLASSIC
	GameName           string                 `json:"gameName"`           // ex: teambuilder-match-4742129795
	GameStartTimestamp int                    `json:"gameStartTimestamp"` // ex: 1483140696030
	GameType           string                 `json:"gameType"`           // ex: MATCHED_GAME
	GameVersion        patch.Patch            `json:"gameVersion"`        // ex: 6.24.204.6436
	MapID              int                    `json:"mapId"`              // ex: 11
	Participants       []MatchInfoParticipant `json:"participants"`
	PlatformID         region.Region          `json:"platformId"` // ex: NA1
	QueueID            queue.ID               `json:"queueId"`    // ex: 420
	Teams              []MatchInfoTeam        `json:"teams"`
	TournamentCode     string                 `json:"tournamentCode"`
}

type MatchInfoParticipant struct {
	AllInPings                     int                             `json:"allInPings"`
	AssistMePings                  int                             `json:"assistMePings"`
	Assists                        int                             `json:"assists"`
	BaitPings                      int                             `json:"baitPings"`
	BaronKills                     int                             `json:"baronKills"`
	BasicPings                     int                             `json:"basicPings"`
	BountyLevel                    int                             `json:"bountyLevel"`
	Challenges                     *MatchInfoParticipantChallenges `json:"challenges"`
	ChampExperience                int                             `json:"champExperience"`
	ChampLevel                     int                             `json:"champLevel"`
	ChampionID                     int                             `json:"championId"`
	ChampionName                   string                          `json:"championName"`
	ChampionTransform              int                             `json:"championTransform"`
	CommandPings                   int                             `json:"commandPings"`
	ConsumablesPurchased           int                             `json:"consumablesPurchased"`
	DamageDealtToBuildings         int                             `json:"damageDealtToBuildings"`
	DamageDealtToObjectives        int                             `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int                             `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int                             `json:"damageSelfMitigated"`
	DangerPings                    int                             `json:"dangerPings"`
	Deaths                         int                             `json:"deaths"`
	DetectorWardsPlaced            int                             `json:"detectorWardsPlaced"`
	DoubleKills                    int                             `json:"doubleKills"`
	DragonKills                    int                             `json:"dragonKills"`
	EligibleForProgression         bool                            `json:"eligibleForProgression"`
	EnemyMissingPings              int                             `json:"enemyMissingPings"`
	EnemyVisionPings               int                             `json:"enemyVisionPings"`
	FirstBloodAssist               bool                            `json:"firstBloodAssist"`
	FirstBloodKill                 bool                            `json:"firstBloodKill"`
	FirstTowerAssist               bool                            `json:"firstTowerAssist"`
	FirstTowerKill                 bool                            `json:"firstTowerKill"`
	GameEndedInEarlySurrender      bool                            `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender           bool                            `json:"gameEndedInSurrender"`
	GetBackPings                   int                             `json:"getBackPings"`
	GoldEarned                     int                             `json:"goldEarned"`
	GoldSpent                      int                             `json:"goldSpent"`
	HoldPings                      int                             `json:"holdPings"`
	IndividualPosition             string                          `json:"individualPosition"`
	InhibitorKills                 int                             `json:"inhibitorKills"`
	InhibitorTakedowns             int                             `json:"inhibitorTakedowns"`
	InhibitorsLost                 int                             `json:"inhibitorsLost"`
	Item0                          int                             `json:"item0"`
	Item1                          int                             `json:"item1"`
	Item2                          int                             `json:"item2"`
	Item3                          int                             `json:"item3"`
	Item4                          int                             `json:"item4"`
	Item5                          int                             `json:"item5"`
	Item6                          int                             `json:"item6"`
	ItemsPurchased                 int                             `json:"itemsPurchased"`
	KillingSprees                  int                             `json:"killingSprees"`
	Kills                          int                             `json:"kills"`
	Lane                           string                          `json:"lane"`
	LargestCriticalStrike          int                             `json:"largestCriticalStrike"`
	LargestKillingSpree            int                             `json:"largestKillingSpree"`
	LargestMultiKill               int                             `json:"largestMultiKill"`
	LongestTimeSpentLiving         int                             `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int                             `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int                             `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int                             `json:"magicDamageTaken"`
	NeedVisionPings                int                             `json:"needVisionPings"`
	NeutralMinionsKilled           int                             `json:"neutralMinionsKilled"`
	NexusKills                     int                             `json:"nexusKills"`
	NexusLost                      int                             `json:"nexusLost"`
	NexusTakedowns                 int                             `json:"nexusTakedowns"`
	ObjectivesStolen               int                             `json:"objectivesStolen"`
	ObjectivesStolenAssists        int                             `json:"objectivesStolenAssists"`
	OnMyWayPings                   int                             `json:"onMyWayPings"`
	ParticipantID                  int                             `json:"participantId"`
	PentaKills                     int                             `json:"pentaKills"`
	Perks                          MatchInfoParticipantPerks       `json:"perks"`
	PhysicalDamageDealt            int                             `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int                             `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int                             `json:"physicalDamageTaken"`
	Placement                      int                             `json:"placement"`
	PlayerAugment1                 int                             `json:"playerAugment1"`
	PlayerAugment2                 int                             `json:"playerAugment2"`
	PlayerAugment3                 int                             `json:"playerAugment3"`
	PlayerAugment4                 int                             `json:"playerAugment4"`
	PlayerScore0                   int                             `json:"playerScore0"`
	PlayerScore1                   int                             `json:"playerScore1"`
	PlayerScore10                  int                             `json:"playerScore10"`
	PlayerScore11                  int                             `json:"playerScore11"`
	PlayerScore2                   int                             `json:"playerScore2"`
	PlayerScore3                   int                             `json:"playerScore3"`
	PlayerScore4                   int                             `json:"playerScore4"`
	PlayerScore5                   int                             `json:"playerScore5"`
	PlayerScore6                   int                             `json:"playerScore6"`
	PlayerScore7                   int                             `json:"playerScore7"`
	PlayerScore8                   int                             `json:"playerScore8"`
	PlayerScore9                   int                             `json:"playerScore9"`
	PlayerSubteamID                int                             `json:"playerSubteamId"`
	PlayedChampSelectPosition      int                             `json:"playedChampSelectPosition"`
	ProfileIcon                    int                             `json:"profileIcon"`
	PushPings                      int                             `json:"pushPings"`
	QuadraKills                    int                             `json:"quadraKills"`
	RiotIdGameName                 string                          `json:"riotIdGameName"`
	RiotIdTagline                  string                          `json:"riotIdTagline"`
	Role                           string                          `json:"role"`
	SightWardsBoughtInGame         int                             `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int                             `json:"spell1Casts"`
	Spell2Casts                    int                             `json:"spell2Casts"`
	Spell3Casts                    int                             `json:"spell3Casts"`
	Spell4Casts                    int                             `json:"spell4Casts"`
	SubteamPlacement               int                             `json:"subteamPlacement"`
	Summoner1Casts                 int                             `json:"summoner1Casts"`
	Summoner1ID                    summoner_spell.ID               `json:"summoner1Id"`
	Summoner2Casts                 int                             `json:"summoner2Casts"`
	Summoner2ID                    summoner_spell.ID               `json:"summoner2Id"`
	SummonerID                     string                          `json:"summonerId"`
	SummonerLevel                  int                             `json:"summonerLevel"`
	SummonerName                   string                          `json:"summonerName"`
	SummonerPuuid                  string                          `json:"puuid"`
	TeamEarlySurrendered           bool                            `json:"teamEarlySurrendered"`
	TeamID                         int                             `json:"teamId"`
	TeamPosition                   string                          `json:"teamPosition"`
	TimeCCingOthers                int                             `json:"timeCCingOthers"`
	TimePlayed                     int                             `json:"timePlayed"`
	TotalAllyJungleMinionsKilled   int                             `json:"totalAllyJungleMinionsKilled"`
	TotalDamageDealt               int                             `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int                             `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int                             `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int                             `json:"totalDamageTaken"`
	TotalEnemyJungleMinionsKilled  int                             `json:"totalEnemyJungleMinionsKilled"`
	TotalHeal                      int                             `json:"totalHeal"`
	TotalHealsOnTeammates          int                             `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int                             `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int                             `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int                             `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int                             `json:"totalUnitsHealed"`
	TripleKills                    int                             `json:"tripleKills"`
	TrueDamageDealt                int                             `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int                             `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int                             `json:"trueDamageTaken"`
	TurretKills                    int                             `json:"turretKills"`
	TurretTakedowns                int                             `json:"turretTakedowns"`
	TurretsLost                    int                             `json:"turretsLost"`
	UnrealKills                    int                             `json:"unrealKills"`
	VisionClearedPings             int                             `json:"visionClearedPings"`
	VisionScore                    int                             `json:"visionScore"`
	VisionWardsBoughtInGame        int                             `json:"visionWardsBoughtInGame"`
	WardsKilled                    int                             `json:"wardsKilled"`
	WardsPlaced                    int                             `json:"wardsPlaced"`
	Win                            bool                            `json:"win"`
}

type MatchInfoParticipantChallenges struct {
	Assist12StreakCount                       int     `json:"12AssistStreakCount"`
	AbilityUses                               int     `json:"abilityUses"`
	AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
	AlliedJungleMonsterKills                  int     `json:"alliedJungleMonsterKills"`
	BaronBuffGoldAdvantageOverThreshold       int     `json:"baronBuffGoldAdvantageOverThreshold"`
	BaronTakedowns                            int     `json:"baronTakedowns"`
	BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
	BountyGold                                int     `json:"bountyGold"`
	BuffsStolen                               int     `json:"buffsStolen"`
	CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
	ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
	ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
	DamagePerMinute                           float64 `json:"damagePerMinute"`
	DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
	DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
	DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
	DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
	DoubleAces                                int     `json:"doubleAces"`
	DragonTakedowns                           int     `json:"dragonTakedowns"`
	EarliestBaron                             float64 `json:"earliestBaron"`
	EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
	EarlyLaningPhaseGoldExpAdvantage          float64 `json:"earlyLaningPhaseGoldExpAdvantage"`
	EffectiveHealAndShielding                 float64 `json:"effectiveHealAndShielding"`
	ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
	ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
	EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
	EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
	EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
	EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
	EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
	FastestLegendary                          float64 `json:"fastestLegendary"`
	EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
	FirstTurretKilled                         int     `json:"firstTurretKilled"`
	FirstTurretKilledTime                     float64 `json:"firstTurretKilledTime"`
	FlawlessAces                              int     `json:"flawlessAces"`
	FullTeamTakedown                          int     `json:"fullTeamTakedown"`
	GameLength                                float64 `json:"gameLength"`
	GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
	GoldPerMinute                             float64 `json:"goldPerMinute"`
	HadOpenNexus                              int     `json:"hadOpenNexus"`
	ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
	InitialBuffCount                          int     `json:"initialBuffCount"`
	InitialCrabCount                          int     `json:"initialCrabCount"`
	JungleCsBefore10Minutes                   float64 `json:"jungleCsBefore10Minutes"`
	JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
	KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
	Kda                                       float64 `json:"kda"`
	KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
	KillParticipation                         float64 `json:"killParticipation"`
	KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
	KillingSprees                             int     `json:"killingSprees"`
	KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
	KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
	KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
	KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
	KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
	KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
	LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
	LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
	LaningPhaseGoldExpAdvantage               float64 `json:"laningPhaseGoldExpAdvantage"`
	LegendaryCount                            int     `json:"legendaryCount"`
	LostAnInhibitor                           int     `json:"lostAnInhibitor"`
	MaxCsAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent"`
	MaxKillDeficit                            int     `json:"maxKillDeficit"`
	MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
	MejaisFullStackInTime                     int     `json:"mejaisFullStackInTime"`
	MoreEnemyJungleThanOpponent               float64 `json:"moreEnemyJungleThanOpponent"`
	MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
	MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
	Multikills                                int     `json:"multikills"`
	MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
	MythicItemUsed                            int     `json:"mythicItemUsed"`
	OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
	OutnumberedKills                          int     `json:"outnumberedKills"`
	OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
	PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
	PerfectGame                               int     `json:"perfectGame"`
	PickKillWithAlly                          int     `json:"pickKillWithAlly"`
	PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
	PoroExplosions                            int     `json:"poroExplosions"`
	QuickCleanse                              int     `json:"quickCleanse"`
	QuickFirstTurret                          int     `json:"quickFirstTurret"`
	QuickSoloKills                            int     `json:"quickSoloKills"`
	RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
	SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
	ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
	ShortestTimeToAceFromFirstTakedown        float64 `json:"shortestTimeToAceFromFirstTakedown"`
	SkillshotsDodged                          int     `json:"skillshotsDodged"`
	SkillshotsHit                             int     `json:"skillshotsHit"`
	SnowballsHit                              int     `json:"snowballsHit"`
	SoloBaronKills                            int     `json:"soloBaronKills"`
	SoloKills                                 int     `json:"soloKills"`
	SoloTurretsLategame                       int     `json:"soloTurretsLategame"`
	StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
	SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
	SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
	TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
	Takedowns                                 int     `json:"takedowns"`
	TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
	TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
	TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
	TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
	TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
	TeamBaronKills                            int     `json:"teamBaronKills"`
	TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
	TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
	TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
	ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
	TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
	TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
	TurretTakedowns                           int     `json:"turretTakedowns"`
	TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
	TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
	TwoWardsOneSweeperCount                   int     `json:"twoWardsOneSweeperCount"`
	UnseenRecalls                             int     `json:"unseenRecalls"`
	VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
	VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
	WardTakedowns                             int     `json:"wardTakedowns"`
	WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
	WardsGuarded                              int     `json:"wardsGuarded"`
}

type MatchInfoParticipantPerks struct {
	StatPerks MatchInfoParticipantPerksStatPerks `json:"statPerks"`
	Styles    []MatchInfoParticipantPerksStyles  `json:"styles"`
}

type MatchInfoParticipantPerksStatPerks struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type MatchInfoParticipantPerksStyles struct {
	Description string                                     `json:"description"`
	Selections  []MatchInfoParticipantPerksStylesSelection `json:"selections"`
	Style       int                                        `json:"style"`
}

type MatchInfoParticipantPerksStylesSelection struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}

type MatchInfoTeam struct {
	Bans       []MatchInfoTeamBan      `json:"bans"`
	Objectives MatchInfoTeamObjectives `json:"objectives"`
	TeamID     int                     `json:"teamId"`
	Win        bool                    `json:"win"`
}

type MatchInfoTeamBan struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
}

type MatchInfoTeamObjectives struct {
	Baron      MatchInfoTeamObjectiveType `json:"baron"`
	Champion   MatchInfoTeamObjectiveType `json:"champion"`
	Dragon     MatchInfoTeamObjectiveType `json:"dragon"`
	Horde      MatchInfoTeamObjectiveType `json:"horde"`
	Inhibitor  MatchInfoTeamObjectiveType `json:"inhibitor"`
	RiftHerald MatchInfoTeamObjectiveType `json:"riftHerald"`
	Tower      MatchInfoTeamObjectiveType `json:"tower"`
}

type MatchInfoTeamObjectiveType struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

func (m MatchInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *MatchInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func (c *client) GetMatch(continent continent.Continent, matchID string) (*Match, error) {
	var res Match
	_, err := c.dispatchAndUnmarshal(c.ctx, continent, "/lol/match/v5/matches", fmt.Sprintf("/%s", matchID), nil, ratelimiter.GetMatch, &res)
	return &res, err
}

type MatchTimeline struct {
	Metadata MatchTimelineMetadata `json:"metadata"`
	Info     MatchTimelineInfo     `json:"info"`
}

type MatchTimelineMetadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type MatchTimelineInfo struct {
	FrameInterval int                        `json:"frameInterval"`
	Frames        []MatchTimelineFrame       `json:"frames"`
	GameID        int                        `json:"gameId"`
	Participants  []MatchTimelineParticipant `json:"participants"`
}

type MatchTimelineParticipant struct {
	ParticipantID int    `json:"participantId"`
	Puuid         string `json:"puuid"`
}

type MatchTimelineFrame struct {
	Timestamp         int                             `json:"timestamp"`
	ParticipantFrames []MatchTimelineParticipantFrame `json:"participantFrames"`
	Events            []interface{}                   `json:"events"`
}

type MatchTimelineFrameEventType string

const (
	AscendedEvent           MatchTimelineFrameEventType = "ASCENDED_EVENT"
	BuildingKill            MatchTimelineFrameEventType = "BUILDING_KILL"
	CapturePoint            MatchTimelineFrameEventType = "CAPTURE_POINT"
	ChampionKill            MatchTimelineFrameEventType = "CHAMPION_KILL"
	ChampionSpecialKill     MatchTimelineFrameEventType = "CHAMPION_SPECIAL_KILL"
	ChampionTransform       MatchTimelineFrameEventType = "CHAMPION_TRANSFORM"
	DragonSoulGiven         MatchTimelineFrameEventType = "DRAGON_SOUL_GIVEN"
	EliteMonsterKill        MatchTimelineFrameEventType = "ELITE_MONSTER_KILL"
	GameEnd                 MatchTimelineFrameEventType = "GAME_END"
	ItemDestroyed           MatchTimelineFrameEventType = "ITEM_DESTROYED"
	ItemPurchased           MatchTimelineFrameEventType = "ITEM_PURCHASED"
	ItemSold                MatchTimelineFrameEventType = "ITEM_SOLD"
	ItemUndo                MatchTimelineFrameEventType = "ITEM_UNDO"
	LevelUp                 MatchTimelineFrameEventType = "LEVEL_UP"
	ObjectiveBountyFinish   MatchTimelineFrameEventType = "OBJECTIVE_BOUNTY_FINISH"
	ObjectiveBountyPreStart MatchTimelineFrameEventType = "OBJECTIVE_BOUNTY_PRESTART"
	PauseEnd                MatchTimelineFrameEventType = "PAUSE_END"
	PoroKingSummon          MatchTimelineFrameEventType = "PORO_KING_SUMMON"
	SkillLevelUp            MatchTimelineFrameEventType = "SKILL_LEVEL_UP"
	TurretPlateDestroyed    MatchTimelineFrameEventType = "TURRET_PLATE_DESTROYED"
	WardKill                MatchTimelineFrameEventType = "WARD_KILL"
	WardPlaced              MatchTimelineFrameEventType = "WARD_PLACED"
)

type MatchTimelineEvent_AscendedEvent struct {
	Timestamp int                         `json:"timestamp"`
	Type      MatchTimelineFrameEventType `json:"type"`
	// TODO
}

type MatchTimelineEvent_BuildingKill struct {
	AssistingParticipantIDs []int                           `json:"assistingParticipantIds"`
	Bounty                  int                             `json:"bounty"`
	BuildingType            MatchTimelineEvent_BuildingType `json:"buildingType"`
	KillerID                int                             `json:"killerId"`
	LaneType                MatchTimelineEvent_LaneType     `json:"laneType"`
	Position                MatchTimelinePosition           `json:"position"`
	TeamID                  int                             `json:"teamId"`
	Timestamp               int                             `json:"timestamp"`
	TowerType               *MatchTimelineEvent_TowerType   `json:"towerType"`
	Type                    MatchTimelineFrameEventType     `json:"type"`
}

type MatchTimelineEvent_CapturePoint struct {
	Timestamp int                         `json:"timestamp"`
	Type      MatchTimelineFrameEventType `json:"type"`
	// TODO
}

type MatchTimelineEvent_ChampionKill struct {
	AssistingParticipantIDs []int                       `json:"assistingParticipantIds"`
	Bounty                  int                         `json:"bounty"`
	KillStreakLength        int                         `json:"killStreakLength"`
	KillerID                int                         `json:"killerId"`
	Position                MatchTimelinePosition       `json:"position"`
	ShutdownBounty          int                         `json:"shutdownBounty"`
	Timestamp               int                         `json:"timestamp"`
	Type                    MatchTimelineFrameEventType `json:"type"`
	VictimDamageDealt       []MatchTimelineDamage       `json:"victimDamageDealt"`
	VictimDamageReceived    []MatchTimelineDamage       `json:"victimDamageReceived"`
	VictimID                int                         `json:"victimId"`
}

type MatchTimelineEvent_ChampionSpecialKill struct {
	KillerType      MatchTimelineEvent_KillerType `json:"killerType"`
	KillerID        int                           `json:"killerId"`
	MultiKillLength *int                          `json:"multiKillLength"`
	Position        MatchTimelinePosition         `json:"position"`
	Timestamp       int                           `json:"timestamp"`
	Type            MatchTimelineFrameEventType   `json:"type"`
}

type MatchTimelineEvent_ChampionTransform struct {
	ParticipantID int                                      `json:"participantId"`
	Timestamp     int                                      `json:"timestamp"`
	TransformType MatchTimelineEvent_ChampionTransformType `json:"transformType"`
	Type          MatchTimelineFrameEventType              `json:"type"`
}

type MatchTimelineEvent_ChampionTransformType string

const (
	MatchTimelineEvent_ChampionTransformType_Assassin MatchTimelineEvent_ChampionTransformType = "Assassin"
	MatchTimelineEvent_ChampionTransformType_Slayer   MatchTimelineEvent_ChampionTransformType = "Slayer"
)

type MatchTimelineEvent_DragonSoulGiven struct {
	Name      MatchTimelineEvent_DragonSoul `json:"name"`
	TeamID    int                           `json:"teamId"`
	Timestamp int                           `json:"timestamp"`
	Type      MatchTimelineFrameEventType   `json:"type"`
}

type MatchTimelineEvent_EliteMonsterKill struct {
	AssistingParticipantIDs []int                             `json:"assistingParticipantIds"`
	Bounty                  int                               `json:"bounty"`
	KillerID                int                               `json:"killerId"`
	KillerTeamID            int                               `json:"killerTeamId"`
	MonsterSubType          MatchTimelineEvent_MonsterSubType `json:"monsterSubType"`
	MonsterType             MatchTimelineEvent_MonsterType    `json:"monsterType"`
	Position                MatchTimelinePosition             `json:"position"`
	Timestamp               int                               `json:"timestamp"`
	Type                    MatchTimelineFrameEventType       `json:"type"`
}

type MatchTimelineEvent_GameEnd struct {
	GameID        int                         `json:"gameId"`
	RealTimestamp int                         `json:"realTimestamp"`
	Timestamp     int                         `json:"timestamp"`
	Type          MatchTimelineFrameEventType `json:"type"`
	WinningTeam   int                         `json:"winningTeam"`
}

type MatchTimelineEvent_ItemDestroyed struct {
	ItemID        int                         `json:"itemId"`
	ParticipantID int                         `json:"participantId"`
	Timestamp     int                         `json:"timestamp"`
	Type          MatchTimelineFrameEventType `json:"type"`
}

type MatchTimelineEvent_ItemPurchased struct {
	ItemID        int                         `json:"itemId"`
	ParticipantID int                         `json:"participantId"`
	Timestamp     int                         `json:"timestamp"`
	Type          MatchTimelineFrameEventType `json:"type"`
}

type MatchTimelineEvent_ItemSold struct {
	ItemID        int                         `json:"itemId"`
	ParticipantID int                         `json:"participantId"`
	Timestamp     int                         `json:"timestamp"`
	Type          MatchTimelineFrameEventType `json:"type"`
}

type MatchTimelineEvent_ItemUndo struct {
	AfterID       int                         `json:"afterId"`
	BeforeID      int                         `json:"beforeId"`
	GoldGain      int                         `json:"goldGain"`
	ParticipantID int                         `json:"participantId"`
	Timestamp     int                         `json:"timestamp"`
	Type          MatchTimelineFrameEventType `json:"type"`
}

type MatchTimelineEvent_LevelUp struct {
	Level         int                         `json:"level"`
	ParticipantID int                         `json:"participantId"`
	Timestamp     int                         `json:"timestamp"`
	Type          MatchTimelineFrameEventType `json:"type"`
}

type MatchTimelineEvent_ObjectiveBountyFinish struct {
	TeamID    int                         `json:"teamId"`
	Timestamp int                         `json:"timestamp"`
	Type      MatchTimelineFrameEventType `json:"type"`
}

type MatchTimelineEvent_ObjectiveBountyPreStart struct {
	ActualStartTime int                         `json:"actualStartTime"`
	TeamID          int                         `json:"teamId"`
	Timestamp       int                         `json:"timestamp"`
	Type            MatchTimelineFrameEventType `json:"type"`
}

type MatchTimelineEvent_PauseEnd struct {
	RealTimestamp int                         `json:"realTimestamp"`
	Timestamp     int                         `json:"timestamp"`
	Type          MatchTimelineFrameEventType `json:"type"`
}

type MatchTimelineEvent_PoroKingSummon struct {
	Timestamp int `json:"timestamp"`
	Type      MatchTimelineFrameEventType
	// TODO
}

type MatchTimelineEvent_SkillLevelUp struct {
	LevelUpType   MatchTimelineEvent_LevelUpType `json:"levelUpType"`
	ParticipantID int                            `json:"participantId"`
	SkillSlot     int                            `json:"skillSlot"`
	Timestamp     int                            `json:"timestamp"`
	Type          MatchTimelineFrameEventType    `json:"type"`
}

type MatchTimelineEvent_TurretPlateDestroyed struct {
	KillerID  int                         `json:"killerId"`
	LaneType  MatchTimelineEvent_LaneType `json:"laneType"`
	Position  MatchTimelinePosition       `json:"position"`
	TeamID    int                         `json:"teamId"`
	Timestamp int                         `json:"timestamp"`
	Type      MatchTimelineFrameEventType `json:"type"`
}

type MatchTimelineEvent_WardKill struct {
	KillerID  int                         `json:"killerId"`
	Timestamp int                         `json:"timestamp"`
	Type      MatchTimelineFrameEventType `json:"type"`
	WardType  MatchTimelineEvent_WardType `json:"wardType"`
}

type MatchTimelineEvent_WardPlaced struct {
	CreatorID int                         `json:"creatorId"`
	Timestamp int                         `json:"timestamp"`
	Type      MatchTimelineFrameEventType `json:"type"`
	WardType  MatchTimelineEvent_WardType `json:"wardType"`
}

type MatchTimelineEvent_BuildingType string

const (
	MatchTimelineEvent_BuildingType_InhibitorBuilding MatchTimelineEvent_BuildingType = "INHIBITOR_BUILDING"
	MatchTimelineEvent_BuildingType_TowerBuilding     MatchTimelineEvent_BuildingType = "TOWER_BUILDING"
)

type MatchTimelineEvent_DragonSoul string

const (
	MatchTimelineEvent_DragonSoul_Chemtech MatchTimelineEvent_DragonSoul = "Chemtech"
	MatchTimelineEvent_DragonSoul_Cloud    MatchTimelineEvent_DragonSoul = "Cloud"
	MatchTimelineEvent_DragonSoul_Hextech  MatchTimelineEvent_DragonSoul = "Hextech"
	MatchTimelineEvent_DragonSoul_Infernal MatchTimelineEvent_DragonSoul = "Infernal"
	MatchTimelineEvent_DragonSoul_Mountain MatchTimelineEvent_DragonSoul = "Mountain"
	MatchTimelineEvent_DragonSoul_Ocean    MatchTimelineEvent_DragonSoul = "Ocean"
)

type MatchTimelineEvent_KillerType string

const (
	MatchTimelineEvent_KillerType_Ace             MatchTimelineEvent_KillerType = "ACE"
	MatchTimelineEvent_KillerType_KillMulti       MatchTimelineEvent_KillerType = "KILL_MULTI"
	MatchTimelineEvent_KillerType_KillFirst_Blood MatchTimelineEvent_KillerType = "KILL_FIRST_BLOOD"
)

type MatchTimelineEvent_LaneType string

const (
	MatchTimelineEvent_LaneType_BotLane MatchTimelineEvent_LaneType = "BOT_LANE"
	MatchTimelineEvent_LaneType_MidLane MatchTimelineEvent_LaneType = "MID_LANE"
	MatchTimelineEvent_LaneType_TopLane MatchTimelineEvent_LaneType = "TOP_LANE"
)

type MatchTimelineEvent_LevelUpType string

const (
	MatchTimelineEvent_LevelUpType_Evolve MatchTimelineEvent_LevelUpType = "EVOLVE"
	MatchTimelineEvent_LevelUpType_Normal MatchTimelineEvent_LevelUpType = "NORMAL"
)

type MatchTimelineEvent_MonsterSubType string

const (
	MatchTimelineEvent_MonsterSubType_AirDragon      MatchTimelineEvent_MonsterSubType = "AIR_DRAGON"
	MatchTimelineEvent_MonsterSubType_ChemtechDragon MatchTimelineEvent_MonsterSubType = "CHEMTECH_DRAGON"
	MatchTimelineEvent_MonsterSubType_EarthDragon    MatchTimelineEvent_MonsterSubType = "EARTH_DRAGON"
	MatchTimelineEvent_MonsterSubType_ElderDragon    MatchTimelineEvent_MonsterSubType = "ELDER_DRAGON"
	MatchTimelineEvent_MonsterSubType_FireDragon     MatchTimelineEvent_MonsterSubType = "FIRE_DRAGON"
	MatchTimelineEvent_MonsterSubType_WaterDragon    MatchTimelineEvent_MonsterSubType = "WATER_DRAGON"
	MatchTimelineEvent_MonsterSubType_HextechDragon  MatchTimelineEvent_MonsterSubType = "HEXTECH_DRAGON"
)

type MatchTimelineEvent_MonsterType string

const (
	MatchTimelineEvent_MonsterType_Baron      MatchTimelineEvent_MonsterType = "BARON_NASHOR"
	MatchTimelineEvent_MonsterType_Dragon     MatchTimelineEvent_MonsterType = "DRAGON"
	MatchTimelineEvent_MonsterType_RiftHerald MatchTimelineEvent_MonsterType = "RIFTHERALD"
)

type MatchTimelineEvent_TowerType string

const (
	MatchTimelineEvent_TowerType_BaseTurret  MatchTimelineEvent_TowerType = "BASE_TURRET"
	MatchTimelineEvent_TowerType_InnerTurret MatchTimelineEvent_TowerType = "INNER_TURRET"
	MatchTimelineEvent_TowerType_NexusTurret MatchTimelineEvent_TowerType = "NEXUS_TURRET"
	MatchTimelineEvent_TowerType_OuterTurret MatchTimelineEvent_TowerType = "OUTER_TURRET"
)

type MatchTimelineEvent_WardType string

const (
	MatchTimelineEvent_WardType_ControlWard MatchTimelineEvent_WardType = "CONTROL_WARD"
	MatchTimelineEvent_WardType_SightWard   MatchTimelineEvent_WardType = "SIGHT_WARD"
	MatchTimelineEvent_WardType_TeemoMush   MatchTimelineEvent_WardType = "TEEMO_MUSHROOM"
	MatchTimelineEvent_WardType_YellowTrink MatchTimelineEvent_WardType = "YELLOW_TRINKET"
)

// Need to unmarshal MatchTimelineFrame because ParticipantFrames comes as an object with keys as participant IDs.
func (m *MatchTimelineFrame) UnmarshalJSON(data []byte) error {
	// Define a temporary struct with the same fields as MatchTimelineFrame
	temp := &struct {
		Timestamp         int                                      `json:"timestamp"`
		ParticipantFrames map[string]MatchTimelineParticipantFrame `json:"participantFrames"`
		Events            []json.RawMessage                        `json:"events"`
	}{}

	// Unmarshal data into the temporary struct
	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}

	// Manually assign the values to the MatchTimelineFrame
	m.Timestamp = temp.Timestamp

	for _, rawMsg := range temp.Events {
		// Unmarshal the type field to determine the event type
		var typeHolder struct {
			Type MatchTimelineFrameEventType `json:"type"`
		}
		if err := json.Unmarshal(rawMsg, &typeHolder); err != nil {
			return err
		}

		var event interface{}
		switch typeHolder.Type {
		case AscendedEvent:
			event = new(MatchTimelineEvent_AscendedEvent)
		case BuildingKill:
			event = new(MatchTimelineEvent_BuildingKill)
		case CapturePoint:
			event = new(MatchTimelineEvent_CapturePoint)
		case ChampionKill:
			event = new(MatchTimelineEvent_ChampionKill)
		case ChampionSpecialKill:
			event = new(MatchTimelineEvent_ChampionSpecialKill)
		case ChampionTransform:
			event = new(MatchTimelineEvent_ChampionTransform)
		case DragonSoulGiven:
			event = new(MatchTimelineEvent_DragonSoulGiven)
		case EliteMonsterKill:
			event = new(MatchTimelineEvent_EliteMonsterKill)
		case GameEnd:
			event = new(MatchTimelineEvent_GameEnd)
		case ItemDestroyed:
			event = new(MatchTimelineEvent_ItemDestroyed)
		case ItemPurchased:
			event = new(MatchTimelineEvent_ItemPurchased)
		case ItemSold:
			event = new(MatchTimelineEvent_ItemSold)
		case ItemUndo:
			event = new(MatchTimelineEvent_ItemUndo)
		case LevelUp:
			event = new(MatchTimelineEvent_LevelUp)
		case ObjectiveBountyFinish:
			event = new(MatchTimelineEvent_ObjectiveBountyFinish)
		case ObjectiveBountyPreStart:
			event = new(MatchTimelineEvent_ObjectiveBountyPreStart)
		case PauseEnd:
			event = new(MatchTimelineEvent_PauseEnd)
		case PoroKingSummon:
			event = new(MatchTimelineEvent_PoroKingSummon)
		case SkillLevelUp:
			event = new(MatchTimelineEvent_SkillLevelUp)
		case TurretPlateDestroyed:
			event = new(MatchTimelineEvent_TurretPlateDestroyed)
		case WardKill:
			event = new(MatchTimelineEvent_WardKill)
		case WardPlaced:
			event = new(MatchTimelineEvent_WardPlaced)
		default:
			continue // Skip unknown event types
		}

		// Unmarshal the event
		if err := json.Unmarshal(rawMsg, event); err != nil {
			return err
		}

		m.Events = append(m.Events, event)
	}

	// Convert map to slice
	m.ParticipantFrames = make([]MatchTimelineParticipantFrame, 0, len(temp.ParticipantFrames))
	for _, pf := range temp.ParticipantFrames {
		m.ParticipantFrames = append(m.ParticipantFrames, pf)
	}

	return nil
}

type MatchTimelineParticipantFrame struct {
	ChampionStats            MatchTimelineChampionStats `json:"championStats"`
	CurrentGold              int                        `json:"currentGold"`
	DamageStats              MatchTimelineDamageStats   `json:"damageStats"`
	GoldPerSecond            int                        `json:"goldPerSecond"`
	JungleMinionsKilled      int                        `json:"jungleMinionsKilled"`
	Level                    int                        `json:"level"`
	MinionsKilled            int                        `json:"minionsKilled"`
	ParticipantID            int                        `json:"participantId"`
	Position                 MatchTimelinePosition      `json:"position"`
	TimeEnemySpentControlled int                        `json:"timeEnemySpentControlled"`
	TotalGold                int                        `json:"totalGold"`
	XP                       int                        `json:"xp"`
}

type MatchTimelineChampionStats struct {
	AbilityHaste         int `json:"abilityHaste"`
	AbilityPower         int `json:"abilityPower"`
	Armor                int `json:"armor"`
	ArmorPen             int `json:"armorPen"`
	ArmorPenPercent      int `json:"armorPenPercent"`
	AttackDamage         int `json:"attackDamage"`
	AttackSpeed          int `json:"attackSpeed"`
	BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
	BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
	CCReduction          int `json:"ccReduction"`
	CooldownReduction    int `json:"cooldownReduction"`
	Health               int `json:"health"`
	HealthMax            int `json:"healthMax"`
	HealthRegen          int `json:"healthRegen"`
	Lifesteal            int `json:"lifesteal"`
	MagicPen             int `json:"magicPen"`
	MagicPenPercent      int `json:"magicPenPercent"`
	MagicResist          int `json:"magicResist"`
	MovementSpeed        int `json:"movementSpeed"`
	Omnivamp             int `json:"omnivamp"`
	PhysicalVamp         int `json:"physicalVamp"`
	Power                int `json:"power"`
	PowerMax             int `json:"powerMax"`
	PowerRegen           int `json:"powerRegen"`
	SpellVamp            int `json:"spellVamp"`
}

type MatchTimelineDamageStats struct {
	MagicDamageDone               int `json:"magicDamageDone"`
	MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
	MagicDamageTaken              int `json:"magicDamageTaken"`
	PhysicalDamageDone            int `json:"physicalDamageDone"`
	PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
	PhysicalDamageTaken           int `json:"physicalDamageTaken"`
	TotalDamageDone               int `json:"totalDamageDone"`
	TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
	TotalDamageTaken              int `json:"totalDamageTaken"`
	TrueDamageDone                int `json:"trueDamageDone"`
	TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
	TrueDamageTaken               int `json:"trueDamageTaken"`
}

type MatchTimelinePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MatchTimelineDamage struct {
	Basic          bool                    `json:"basic"`
	MagicDamage    int                     `json:"magicDamage"`
	Name           string                  `json:"name"`
	ParticipantID  int                     `json:"participantId"`
	PhysicalDamage int                     `json:"physicalDamage"`
	SpellName      string                  `json:"spellName"`
	SpellSlot      int                     `json:"spellSlot"`
	TrueDamage     int                     `json:"trueDamage"`
	Type           MatchTimelineDamageType `json:"type"`
}

type MatchTimelineDamageType string

const (
	MatchTimelineDamageType_Minion  MatchTimelineDamageType = "MINION"
	MatchTimelineDamageType_Monster MatchTimelineDamageType = "MONSTER"
	MatchTimelineDamageType_Tower   MatchTimelineDamageType = "TOWER"
	MatchTimelineDamageType_Other   MatchTimelineDamageType = "OTHER"
)

func (c *client) GetMatchTimeline(continent continent.Continent, matchID string) (*MatchTimeline, error) {
	var res MatchTimeline
	_, err := c.dispatchAndUnmarshal(c.ctx, continent, "/lol/match/v5/matches", fmt.Sprintf("/%s/timeline", matchID), nil, ratelimiter.GetMatchTimeline, &res)
	return &res, err
}
