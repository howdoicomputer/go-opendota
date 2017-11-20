package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

const openDotaAPI = "https://api.opendota.com/api/"

// Client for making Open Dota API requests
type Client struct {
	sling               *sling.Sling
	DistributionService *DistributionService
	ExplorerService     *ExplorerService
	MatchService        *MatchService
	MetadataService     *MetadataService
	PlayerService       *PlayerService
	ProMatchService     *ProMatchService
	ProPlayerService    *ProPlayerService
	PublicMatchService  *PublicMatchService
	RankingService      *RankingService
	SearchService       *SearchService
	TeamService         *TeamService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(openDotaAPI)

	return &Client{
		sling:               base,
		DistributionService: newDistributionService(base.New()),
		ExplorerService:     newExplorerService(base.New()),
		MatchService:        newMatchService(base.New()),
		MetadataService:     newMetadataService(base.New()),
		PlayerService:       newPlayerService(base.New()),
		ProMatchService:     newProMatchService(base.New()),
		ProPlayerService:    newProPlayerService(base.New()),
		PublicMatchService:  newPublicMatchService(base.New()),
		RankingService:      newRankingService(base.New()),
		SearchService:       newSearchService(base.New()),
		TeamService:         newTeamService(base.New()),
	}
}
