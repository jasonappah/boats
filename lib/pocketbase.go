package lib

import (
  "github.com/pocketbase/pocketbase"
  "github.com/pocketbase/pocketbase/plugins/migratecmd"

  _ "boats/migrations"
  "os"
  "path/filepath"
)

func InitPocketbase(runningInDev bool) *pocketbase.PocketBase {
  packageName := "me.jasonaa.boats"
  
  home, err := os.UserConfigDir()
  if err != nil {
    panic(err)
  }
  
  appDir := filepath.Join(home, packageName)

  if _, err := os.Stat(appDir); os.IsNotExist(err) {
    os.Mkdir(appDir, 0755)
  }
  
  pocketBaseApp := pocketbase.NewWithConfig(pocketbase.Config{
    DefaultDev: runningInDev,
    DefaultDataDir: appDir,
  })
  
  migratecmd.MustRegister(pocketBaseApp, pocketBaseApp.RootCmd, migratecmd.Config{
      Automigrate: runningInDev,
  })
  
  return pocketBaseApp
}
