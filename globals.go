package globals

import (
	"context"
	"sync"

	cam "google.golang.org/api/cloudasset/v1"
	crm "google.golang.org/api/cloudresourcemanager/v1"
	// "github.com/briandowns/spinner"
)
	
	
var (
	GetContext func() context.Context
	CrmService crm.Service
	CamService *cam.Service
	MapOfAllTFResourcesWithData *sync.Map
	ListOfAllTFResourcesWithData []*Asset
	ListOfAllTFResourcesPerPath *sync.Map
	// CmdProgress *spinner.Spinner
)


type (
	Ancestors []string
	Asset struct {
		Ancestors 		Ancestors `json:"folders"`
		AssetType 		string `json:"assetType"`
		Name 			string `json:"name"`
		DisplayName 	string `json:"displayName"`
		ParentAssetType string `json:"parentAssetType"`
		Project 		string `json:"project"`
		Organization 	string `json:"organization"`
		Parent			string `json:"parentFullResourceName"`
		Data 			map[string]any `json:"data"`
		// updateTime string `json:update_time`
		IsSupported 	bool `json:"supported"`	
	}
)


var GlobalCMDRunner func(pwd string, pipe []any, cmdp ...string) any


func init(){
	GetContext = func() context.Context { return context.Background()}
	CamService, _ = cam.NewService(GetContext()) 
	MapOfAllTFResourcesWithData = new(sync.Map)
	ListOfAllTFResourcesWithData = []*Asset{}
	// CmdProgress = spinner.New(spinner.CharSets[88], 100*time.Second)
}