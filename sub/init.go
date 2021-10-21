package sub

import (
	"OneAuth/cfg"
	"OneAuth/libs/auth"
	"OneAuth/models"
	"github.com/urfave/cli/v2"
	"github.com/veypi/utils/cmd"
	"github.com/veypi/utils/log"
	"strconv"
)

var Init = &cli.Command{
	Name:   "init",
	Action: runInit,
}

func runInit(c *cli.Context) error {
	return InitSystem()
}

// 初始化项目

func InitSystem() error {
	db()
	self, err := selfApp()
	if err != nil {
		return err
	}
	cfg.CFG.APPID = self.ID
	cfg.CFG.APPKey = self.Key
	err = cmd.DumpCfg(cfg.Path, cfg.CFG)
	// TODO
	//if err != nil {
	//	return err
	//}
	err = role(self.InitRoleID == 0)
	return nil
}

func db() {
	db := cfg.DB()
	log.HandlerErrs(
		db.SetupJoinTable(&models.User{}, "Roles", &models.UserRole{}),
		db.SetupJoinTable(&models.Role{}, "Users", &models.UserRole{}),
		db.SetupJoinTable(&models.User{}, "Apps", &models.AppUser{}),
		db.SetupJoinTable(&models.App{}, "Users", &models.AppUser{}),
		db.AutoMigrate(&models.User{}, &models.Role{}, &models.Auth{}, &models.App{}),
	)
	log.HandlerErrs(
		db.AutoMigrate(&models.Wechat{}, &models.Resource{}),
	)
}

func selfApp() (*models.App, error) {
	self := &models.App{
		Name:           "OA",
		Icon:           "",
		UUID:           "jU5Jo5hM",
		Des:            "",
		Creator:        0,
		UserCount:      0,
		Hide:           false,
		Host:           "",
		UserRefreshUrl: "/",
		Key:            "cB43wF94MLTksyBK",
		EnableRegister: true,
		EnableUserKey:  true,
		EnableUser:     true,
		EnableWx:       false,
		EnablePhone:    false,
		EnableEmail:    false,
		Wx:             nil,
	}
	return self, cfg.DB().Where("uuid = ?", self.UUID).FirstOrCreate(self).Error
}

func role(reset_init_role bool) error {
	authMap := make(map[string]*models.Resource)
	n := []string{
		auth.APP,
		auth.User,
		auth.Res,
		auth.Auth,
		auth.Role,
	}
	var err error
	adminRole := &models.Role{
		AppID:    cfg.CFG.APPID,
		Name:     "admin",
		IsUnique: false,
	}
	err = cfg.DB().Where(adminRole).FirstOrCreate(adminRole).Error
	if err != nil {
		return err
	}
	for _, na := range n {
		a := &models.Resource{
			AppID: cfg.CFG.APPID,
			Name:  na,
			Tag:   "",
			Des:   "",
		}
		err = cfg.DB().Where(a).FirstOrCreate(a).Error
		if err != nil {
			return err
		}
		authMap[na] = a
		err = auth.BindRoleAuth(cfg.DB(), adminRole.ID, a.ID, models.AuthAll, "")
		if err != nil {
			return err
		}
	}
	userRole := &models.Role{
		AppID:    cfg.CFG.APPID,
		Name:     "user",
		IsUnique: false,
	}
	err = cfg.DB().Where(userRole).FirstOrCreate(userRole).Error
	if err != nil {
		return err
	}
	err = auth.BindRoleAuth(cfg.DB(), userRole.ID, authMap[auth.APP].ID, models.AuthRead, strconv.Itoa(int(cfg.CFG.APPID)))
	if err != nil {
		return err
	}
	if reset_init_role {
		return cfg.DB().Model(&models.App{}).Where("id = ?", cfg.CFG.APPID).Update("init_role_id", adminRole.ID).Error
	}
	return nil
}
