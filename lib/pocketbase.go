package lib

import (
  "github.com/pocketbase/pocketbase"
  "github.com/pocketbase/pocketbase/plugins/migratecmd"

  _ "boats/migrations"
)

func InitPocketbase(runningInDev bool) *pocketbase.PocketBase {
  pocketBaseApp := pocketbase.NewWithConfig(pocketbase.Config{
    DefaultDev: runningInDev,
  })
  
  migratecmd.MustRegister(pocketBaseApp, pocketBaseApp.RootCmd, migratecmd.Config{
      Automigrate: runningInDev,
  })
  
  return pocketBaseApp
}
