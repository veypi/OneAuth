package sub

import (
	"github.com/urfave/cli/v2"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/auth"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneAuth/oalib"
	"github.com/veypi/utils"
)

var Init = &cli.Command{
	Name:   "init",
	Action: runInit,
}

func runInit(c *cli.Context) error {
	return InitSystem()
}

func InitSystem() error {
	err := db()
	if err != nil {
		return err
	}
	self, err := selfApp()
	if err != nil {
		return err
	}
	err = role(self.InitRoleID == 0)
	return err
}

func db() error {
	db := cfg.DB()
	err := utils.MultiErr(
		db.SetupJoinTable(&models.User{}, "Roles", &models.UserRole{}),
		db.SetupJoinTable(&models.Role{}, "Users", &models.UserRole{}),
		db.SetupJoinTable(&models.User{}, "Apps", &models.AppUser{}),
		db.SetupJoinTable(&models.App{}, "Users", &models.AppUser{}),
		db.AutoMigrate(&models.User{}, &models.App{}, &models.Auth{}, &models.Role{}),
		db.AutoMigrate(&models.Wechat{}, &models.Resource{}),
	)
	return err
}

func selfApp() (*models.App, error) {
	self := &models.App{
		Name:           "OA",
		Icon:           "",
		UUID:           cfg.CFG.APPUUID,
		Des:            "",
		Creator:        0,
		UserCount:      0,
		Hide:           false,
		Host:           "",
		UserRefreshUrl: "/",
		Key:            string(cfg.CFG.APPKey),
		EnableRegister: true,
		EnableUserKey:  true,
		EnableUser:     true,
		EnableWx:       false,
		EnablePhone:    false,
		EnableEmail:    false,
	}
	return self, cfg.DB().Where("UUID = ?", self.UUID).FirstOrCreate(self).Error
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
		AppUUID: cfg.CFG.APPUUID,
		Name:    "admin",
	}
	err = cfg.DB().Where(adminRole).FirstOrCreate(adminRole).Error
	if err != nil {
		return err
	}
	for _, na := range n {
		a := &models.Resource{
			AppUUID: cfg.CFG.APPUUID,
			Name:    na,
			Des:     "",
		}
		err = cfg.DB().Where(a).FirstOrCreate(a).Error
		if err != nil {
			return err
		}
		authMap[na] = a
		err = auth.BindRoleAuth(cfg.DB(), adminRole.ID, a.ID, oalib.AuthAll, "")
		if err != nil {
			return err
		}
	}
	userRole := &models.Role{
		AppUUID: cfg.CFG.APPUUID,
		Name:    "user",
	}
	err = cfg.DB().Where(userRole).FirstOrCreate(userRole).Error
	if err != nil {
		return err
	}
	e1 := auth.BindRoleAuth(cfg.DB(), userRole.ID, authMap[auth.APP].ID, oalib.AuthRead, "")
	if err := utils.MultiErr(e1); err != nil {
		return err
	}
	if reset_init_role {
		return cfg.DB().Model(&models.App{}).Where("UUID = ?", cfg.CFG.APPUUID).Update("InitRoleID", adminRole.ID).Error
	}
	return nil
}
