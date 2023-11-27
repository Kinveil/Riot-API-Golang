package apiclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/continent"
	"github.com/junioryono/Riot-API-Golang/src/constants/patch"
	"github.com/junioryono/Riot-API-Golang/src/constants/queue"
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
	StartTime *time.Time   `json:"startTime"`
	EndTime   *time.Time   `json:"endTime"`
	Queue     *queue.Queue `json:"queue"`
	Type      *string      `json:"type"`
	Start     *int         `json:"start"`
	Count     *int         `json:"count"`
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
	_, err := c.dispatchAndUnmarshal(continent, "/lol/match/v5/matches/by-puuid", fmt.Sprintf("/%s/ids", puuid), params, ratelimiter.GetMatchlist, &res)
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
	PlatformID         string                 `json:"platformId"` // ex: NA1
	QueueId            int                    `json:"queueId"`    // ex: 420
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
	ChampionID                     int64                           `json:"championId"`
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
	ParticipantId                  int                             `json:"participantId"`
	PentaKills                     int                             `json:"pentaKills"`
	Perks                          MatchInfoParticipantPerks       `json:"perks"`
	PhysicalDamageDealt            int                             `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int                             `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int                             `json:"physicalDamageTaken"`
	PlayedChampSelectPosition      int                             `json:"playedChampSelectPosition"`
	ProfileIcon                    int                             `json:"profileIcon"`
	PushPings                      int                             `json:"pushPings"`
	QuadraKills                    int                             `json:"quadraKills"`
	RiotIdName                     string                          `json:"riotIdName"`
	RiotIdTagline                  string                          `json:"riotIdTagline"`
	Role                           string                          `json:"role"`
	SightWardsBoughtInGame         int                             `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int                             `json:"spell1Casts"`
	Spell2Casts                    int                             `json:"spell2Casts"`
	Spell3Casts                    int                             `json:"spell3Casts"`
	Spell4Casts                    int                             `json:"spell4Casts"`
	Summoner1Casts                 int                             `json:"summoner1Casts"`
	Summoner1Id                    int                             `json:"summoner1Id"`
	Summoner2Casts                 int                             `json:"summoner2Casts"`
	Summoner2Id                    int                             `json:"summoner2Id"`
	SummonerID                     string                          `json:"summonerId"`
	SummonerLevel                  int                             `json:"summonerLevel"`
	SummonerName                   string                          `json:"summonerName"`
	SummonerPuuid                  string                          `json:"puuid"`
	TeamEarlySurrendered           bool                            `json:"teamEarlySurrendered"`
	TeamID                         int                             `json:"teamId"`
	TeamPosition                   string                          `json:"teamPosition"`
	TimeCCingOthers                int                             `json:"timeCCingOthers"`
	TimePlayed                     int                             `json:"timePlayed"`
	TotalDamageDealt               int                             `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int                             `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int                             `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int                             `json:"totalDamageTaken"`
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
	Assist12StreakCount                      int     `json:"12AssistStreakCount"`
	AbilityUses                              int     `json:"abilityUses"`
	AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
	AlliedJungleMonsterKills                 float64 `json:"alliedJungleMonsterKills"`
	BaronTakedowns                           int     `json:"baronTakedowns"`
	BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
	BountyGold                               int     `json:"bountyGold"`
	BuffsStolen                              int     `json:"buffsStolen"`
	CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
	ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
	DamagePerMinute                          float64 `json:"damagePerMinute"`
	DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
	DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
	DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
	DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
	DoubleAces                               int     `json:"doubleAces"`
	DragonTakedowns                          int     `json:"dragonTakedowns"`
	EarlyLaningPhaseGoldExpAdvantage         float64 `json:"earlyLaningPhaseGoldExpAdvantage"`
	EffectiveHealAndShielding                float64 `json:"effectiveHealAndShielding"`
	ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
	ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
	EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
	EnemyJungleMonsterKills                  float64 `json:"enemyJungleMonsterKills"`
	EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
	EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
	EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
	EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
	FirstTurretKilledTime                    float64 `json:"firstTurretKilledTime"`
	FlawlessAces                             int     `json:"flawlessAces"`
	FullTeamTakedown                         int     `json:"fullTeamTakedown"`
	GameLength                               float64 `json:"gameLength"`
	GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
	GoldPerMinute                            float64 `json:"goldPerMinute"`
	HadOpenNexus                             int     `json:"hadOpenNexus"`
	ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
	InitialBuffCount                         int     `json:"initialBuffCount"`
	InitialCrabCount                         int     `json:"initialCrabCount"`
	JungleCsBefore10Minutes                  float64 `json:"jungleCsBefore10Minutes"`
	JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
	KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
	Kda                                      float64 `json:"kda"`
	KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
	KillParticipation                        float64 `json:"killParticipation"`
	KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
	KillingSprees                            int     `json:"killingSprees"`
	KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
	KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
	KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
	KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
	KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
	KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
	LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
	LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
	LaningPhaseGoldExpAdvantage              float64 `json:"laningPhaseGoldExpAdvantage"`
	LegendaryCount                           int     `json:"legendaryCount"`
	LostAnInhibitor                          int     `json:"lostAnInhibitor"`
	MaxCsAdvantageOnLaneOpponent             float64 `json:"maxCsAdvantageOnLaneOpponent"`
	MaxKillDeficit                           int     `json:"maxKillDeficit"`
	MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
	MoreEnemyJungleThanOpponent              float64 `json:"moreEnemyJungleThanOpponent"`
	MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
	MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
	Multikills                               int     `json:"multikills"`
	MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
	MythicItemUsed                           int     `json:"mythicItemUsed"`
	OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
	OutnumberedKills                         int     `json:"outnumberedKills"`
	OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
	PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
	PerfectGame                              int     `json:"perfectGame"`
	PickKillWithAlly                         int     `json:"pickKillWithAlly"`
	PoroExplosions                           int     `json:"poroExplosions"`
	QuickCleanse                             int     `json:"quickCleanse"`
	QuickFirstTurret                         int     `json:"quickFirstTurret"`
	QuickSoloKills                           int     `json:"quickSoloKills"`
	RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
	SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
	ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
	SkillshotsDodged                         int     `json:"skillshotsDodged"`
	SkillshotsHit                            int     `json:"skillshotsHit"`
	SnowballsHit                             int     `json:"snowballsHit"`
	SoloBaronKills                           int     `json:"soloBaronKills"`
	SoloKills                                int     `json:"soloKills"`
	SoloTurretsLategame                      int     `json:"soloTurretsLategame"`
	StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
	SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
	SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
	TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
	Takedowns                                int     `json:"takedowns"`
	TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
	TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
	TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
	TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
	TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
	TeamBaronKills                           int     `json:"teamBaronKills"`
	TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
	TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
	TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
	ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
	TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
	TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
	TurretTakedowns                          int     `json:"turretTakedowns"`
	TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
	TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
	UnseenRecalls                            int     `json:"unseenRecalls"`
	VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
	VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
	WardTakedowns                            int     `json:"wardTakedowns"`
	WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
	WardsGuarded                             int     `json:"wardsGuarded"`
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
	PickTurn   int   `json:"pickTurn"`
	ChampionID int64 `json:"championId"`
}

type MatchInfoTeamObjectives struct {
	Baron      MatchInfoTeamObjectiveType `json:"baron"`
	Champion   MatchInfoTeamObjectiveType `json:"champion"`
	Dragon     MatchInfoTeamObjectiveType `json:"dragon"`
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
	_, err := c.dispatchAndUnmarshal(continent, "/lol/match/v5/matches", fmt.Sprintf("/%s", matchID), nil, ratelimiter.GetMatch, &res)
	return &res, err
}

type MatchTimeline struct {
	MatchTimelineMetadata MatchTimelineMetadata `json:"metadata"`
	Info                  MatchTimelineInfo     `json:"info"`
}

type MatchTimelineMetadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type MatchTimelineInfo struct {
	FrameInterval int64                      `json:"frameInterval"`
	Frames        []MatchTimelineFrame       `json:"frames"`
	GameID        int64                      `json:"gameId"`
	Participants  []MatchTimelineParticipant `json:"participants"`
}

type MatchTimelineParticipant struct {
	ParticipantId int    `json:"participantId"`
	Puuid         string `json:"puuid"`
}

type MatchTimelineFrame struct {
	Timestamp         int64                          `json:"timestamp"`
	ParticipantFrames MatchTimelineParticipantFrames `json:"participantFrames"`
	Events            []MatchTimelineEvent           `json:"events"`
}

// ParticipantFrames stores frames corresponding to each participant. The order
// is not defined (i.e. do not assume the order is ascending by participant ID).
type MatchTimelineParticipantFrames struct {
	Frames []MatchTimelineParticipantFrame `json:"frames"`
}

func (p *MatchTimelineParticipantFrames) UnmarshalJSON(b []byte) error {
	var obj map[int]MatchTimelineParticipantFrame
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}
	var vals []MatchTimelineParticipantFrame

	for _, v := range obj {
		vals = append(vals, v)
	}
	p.Frames = vals
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
	ParticipantId            int                        `json:"participantId"`
	Position                 MatchTimelinePosition      `json:"position"`
	TimeEnemySpentControlled int                        `json:"timeEnemySpentControlled"`
	TotalGold                int                        `json:"totalGold"`
	TeamScore                int                        `json:"teamScore"`
	XP                       int                        `json:"xp"`
}

type MatchTimelineChampionStats struct {
	AblityHaste          int `json:"abilityHaste"`
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
}

type MatchTimelinePosition struct {
	Y int `json:"y"`
	X int `json:"x"`
}

type MatchTimelineEventType string

const (
	ChampionKill      MatchTimelineEventType = "CHAMPION_KILL"
	WardPlaced        MatchTimelineEventType = "WARD_PLACED"
	WardKill          MatchTimelineEventType = "WARD_KILL"
	BuildingKill      MatchTimelineEventType = "BUILDING_KILL"
	EliteMonsterKill  MatchTimelineEventType = "ELITE_MONSTER_KILL"
	ItemPurchased     MatchTimelineEventType = "ITEM_PURCHASED"
	ItemSold          MatchTimelineEventType = "ITEM_SOLD"
	ItemDestroyed     MatchTimelineEventType = "ITEM_DESTROYED"
	ItemUndo          MatchTimelineEventType = "ITEM_UNDO"
	SkillLevelUp      MatchTimelineEventType = "SKILL_LEVEL_UP"
	AscendedEvent     MatchTimelineEventType = "ASCENDED_EVENT"
	CapturePoint      MatchTimelineEventType = "CAPTURE_POINT"
	PoroKingSummon    MatchTimelineEventType = "PORO_KING_SUMMON"
	ChampionTransform MatchTimelineEventType = "CHAMPION_TRANSFORM"
)

type MatchTimelineEvent struct {
	EventType               string                 `json:"eventType"`
	TowerType               string                 `json:"towerType"`
	TeamID                  int                    `json:"teamId"`
	AscendedType            string                 `json:"ascendedType"`
	KillerId                int                    `json:"killerId"`
	LevelUpType             string                 `json:"levelUpType"`
	PointCaptured           string                 `json:"pointCaptured"`
	AssistingParticipantIds []int                  `json:"assistingParticipantIds"`
	WardType                string                 `json:"wardType"`
	MonsterType             string                 `json:"monsterType"`
	Type                    MatchTimelineEventType `json:"type"`
	TransformType           string                 `json:"transformType"`
	SkillSlot               int                    `json:"skillSlot"`
	VictimId                int                    `json:"victimId"`
	Timestamp               int64                  `json:"timestamp"`
	AfterId                 int                    `json:"afterId"`
	MonsterSubType          string                 `json:"monsterSubType"`
	LaneType                string                 `json:"laneType"`
	ItemId                  int                    `json:"itemId"`
	ParticipantId           int                    `json:"participantId"`
	BuildingType            string                 `json:"buildingType"`
	CreatorId               int                    `json:"creatorId"`
	Position                MatchTimelinePosition  `json:"position"`
	BeforeId                int                    `json:"beforeId"`
}

func (c *client) GetMatchTimeline(continent continent.Continent, matchID string) (*MatchTimeline, error) {
	var res MatchTimeline
	_, err := c.dispatchAndUnmarshal(continent, "/lol/match/v5/matches", fmt.Sprintf("/%s/timeline", matchID), nil, ratelimiter.GetMatchTimeline, &res)
	return &res, err
}
