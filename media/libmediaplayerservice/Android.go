package libRKmediaplayerservice

import (
    "android/soong/android"
    "android/soong/cc"
    "fmt"
    "strings"
)

func init() {
    fmt.Println("libRKmediaplayerservice want to conditional Compile")
    android.RegisterModuleType("cc_libRKmediaplayerservice", DefaultsFactory)
}

func DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, AddMediaplayerserviceShardLibs)
    return module
}

func AddMediaplayerserviceShardLibs(ctx android.LoadHookContext) {
    type props struct {
        Shared_libs []string
    }
    var src_fix string = getPlatformLibname(ctx)
    p := &props{}
    p.Shared_libs = append(p.Shared_libs, src_fix)
    ctx.AppendProperties(p)
}

func getPlatformLibname(ctx android.BaseContext) (string) {
    var interface_version string = "librockit_client@1.0"
    if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM"),"rk3326") || strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM"),"rk3126c") || strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM"),"rk3399") || strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM"),"rk3399pro")) {
        interface_version = "librockit_interface"
    }
    fmt.Println("PlatformLibname: " + interface_version)
    return interface_version
}
