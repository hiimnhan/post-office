package theme

import "github.com/spf13/viper"

type Theme struct {
	Border           string `mapstructure:border`
	BackgroundWindow string `mapstructure:backgroundWindow`
	Cursor           string `mapstructure:cursor`

	SideBarTitle    string `mapstructure:sideBarTitle`
	SideBarItem     string `mapstructure:sideBarItem`
	SidebarSelected string `mapstructure:sidebarSelected`
	SidebarFocus    string `mapstructure:sidebarFocus`

	MainPanelFocus string `mapstructure:mainPanelFocus`

	MethodGet     string `mapstructure:methodGet`
	MethodPost    string `mapstructure:methodPost`
	MethodPut     string `mapstructure:methodPut`
	MethodPatch   string `mapstructure:methodPatch`
	MethodDelete  string `mapstructure:methodDelete`
	MethodOptions string `mapstructure:methodOptions`
}

func LoadThemeConfig(cfgFile string) (*Theme, error) {
	config := Theme{}
	err := viper.UnmarshalKey("theme", &config)

	return &config, err
}
