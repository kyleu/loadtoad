# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_arm64/loadtoad"]
bundle_id = "dev.kyleu.loadtoad"

notarize {
  path = "./build/dist/loadtoad_0.0.13_darwin_arm64_desktop.dmg"
  bundle_id = "dev.kyleu.loadtoad"
}

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/loadtoad_0.0.13_darwin_arm64.dmg"
  volume_name = "Load Toad"
}

zip {
  output_path = "./build/dist/loadtoad_0.0.13_darwin_arm64_notarized.zip"
}
