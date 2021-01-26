package ciossdk

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/optim-corp/cios-golang-sdk/model"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

var (
	createClient = func(basePath string) *cios.APIClient {
		c := cios.NewAPIClient(cios.NewConfiguration())
		sp := strings.Split(basePath, "://")
		if len(sp) >= 2 {
			c.GetConfig().Scheme = sp[0]
			c.GetConfig().Host = sp[1]
		}
		return c
	}
)

type (
	CiosClient struct {
		PubSub                *PubSub
		Account               *Account
		DeviceAssetManagement *DeviceAssetManagement
		DeviceManagement      *DeviceManagement
		FileStorage           *FileStorage
		Geography             *Geography
		Auth                  *Auth
		License               *License
		Contract              *Contract
		authType              model.AuthType
	}
	CiosClientConfig struct {
		AutoRefresh  bool
		Debug        bool
		Urls         model.CIOSUrl
		RefreshToken string
		ClientID     string
		ClientSecret string
		RequestScope string
	}
)

func NewCiosClient(config CiosClientConfig) *CiosClient {
	instance := new(CiosClient)
	instance.Contract = &Contract{ApiClient: createClient(config.Urls.ContractUrl), Url: config.Urls.ContractUrl}
	instance.License = &License{ApiClient: createClient(config.Urls.LicenseUrl), Url: config.Urls.LicenseUrl}
	instance.Auth = &Auth{ApiClient: createClient(config.Urls.AuthUrl), Url: config.Urls.AuthUrl}
	instance.PubSub = &PubSub{ApiClient: createClient(config.Urls.MessagingUrl), Url: config.Urls.MessagingUrl}
	instance.Account = &Account{ApiClient: createClient(config.Urls.AccountsUrl), Url: config.Urls.AccountsUrl}
	instance.DeviceManagement = &DeviceManagement{ApiClient: createClient(config.Urls.DeviceManagementUrl), Url: config.Urls.DeviceManagementUrl}
	instance.DeviceAssetManagement = &DeviceAssetManagement{ApiClient: createClient(config.Urls.DeviceAssetManagementUrl), Url: config.Urls.DeviceAssetManagementUrl}
	instance.FileStorage = &FileStorage{ApiClient: createClient(config.Urls.StorageUrl), Url: config.Urls.StorageUrl}
	instance.Geography = &Geography{ApiClient: createClient(config.Urls.LocationUrl), Url: config.Urls.LocationUrl}

	if config.ClientID != "" && config.ClientSecret != "" {
		instance.authType = model.CLIENT_TYPE
		instance.Auth.clientId = config.ClientID
		instance.Auth.clientSecret = config.ClientSecret
		instance.Auth.scope = config.RequestScope
		if config.RefreshToken != "" {
			instance.Auth.ref = config.RefreshToken
			instance.authType = model.REFRESH_TYPE
		}

	} else if false {
		// TODO: Add Device Auth
		instance.authType = model.DEVICE_TYPE
	} else {
		instance.authType = model.NONE_TYPE
	}
	if config.AutoRefresh {
		refFunc := func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error) {
			switch instance.authType {
			case model.CLIENT_TYPE:
				token, scope, tokenType, exp, err := instance.Auth.GetAccessTokenOnClient()
				instance._accessToken(token)
				return token, scope, tokenType, exp, err
			case model.REFRESH_TYPE:
				token, scope, tokenType, exp, err := instance.Auth.GetAccessTokenByRefreshToken()
				instance._accessToken(token)
				return token, scope, tokenType, exp, err
			case model.DEVICE_TYPE:
				return instance.Auth.GetAccessTokenOnDevice()
			default:
				return "", "", "", 0, fmt.Errorf("%s", "No AccessToken")
			}
		}
		instance.PubSub.refresh = &refFunc
		instance.Account.refresh = &refFunc
		instance.License.refresh = &refFunc
		instance.Contract.refresh = &refFunc
		instance.Geography.refresh = &refFunc
		instance.FileStorage.refresh = &refFunc
		instance.DeviceManagement.refresh = &refFunc
		instance.DeviceAssetManagement.refresh = &refFunc
	}
	instance.Debug(config.Debug)

	return instance
}

func (self *CiosClient) Debug(debug bool) *CiosClient {
	self.Auth.ApiClient.GetConfig().Debug = false
	self.PubSub.ApiClient.GetConfig().Debug = debug
	self.DeviceManagement.ApiClient.GetConfig().Debug = debug
	self.Account.ApiClient.GetConfig().Debug = debug
	self.Geography.ApiClient.GetConfig().Debug = debug
	self.DeviceAssetManagement.ApiClient.GetConfig().Debug = debug
	self.FileStorage.ApiClient.GetConfig().Debug = debug
	self.PubSub.debug = debug
	return self
}
func (self *CiosClient) _accessToken(accessToken string) *CiosClient {
	accessToken = "Bearer " + regexp.MustCompile(`^bearer|Bearer| `).ReplaceAllString(accessToken, "")
	self.Auth.ApiClient.GetConfig().AddDefaultHeader("Authorization", accessToken)
	self.PubSub.ApiClient.GetConfig().AddDefaultHeader("Authorization", accessToken)
	self.DeviceManagement.ApiClient.GetConfig().AddDefaultHeader("Authorization", accessToken)
	self.Account.ApiClient.GetConfig().AddDefaultHeader("Authorization", accessToken)
	self.Geography.ApiClient.GetConfig().AddDefaultHeader("Authorization", accessToken)
	self.DeviceAssetManagement.ApiClient.GetConfig().AddDefaultHeader("Authorization", accessToken)
	self.FileStorage.ApiClient.GetConfig().AddDefaultHeader("Authorization", accessToken)
	return self
}
func (self *CiosClient) RequestScope(scope string) *CiosClient {
	self.Auth.scope = scope
	return self
}

func MakeRequestCtx(token string) model.RequestCtx {
	if token == "" {
		return context.Background()
	}
	return context.WithValue(context.Background(), cios.ContextAccessToken, regexp.MustCompile(`^bearer|Bearer| `).ReplaceAllString(token, ""))
}
func ParseAccessToken(accessToken string) string {
	if !strings.Contains(accessToken, "Bearer ") {
		return "Bearer " + accessToken
	} else {
		return accessToken
	}
}

func (self CiosClient) String() string {
	// TODO: ToString
	return ""
}
