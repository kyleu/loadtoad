# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_amd64_v1/loadtoad"]
bundle_id = "dev.kyleu.loadtoad"

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/loadtoad_0.0.6_darwin_amd64.dmg"
  volume_name = "Load Toad"
}

zip {
  output_path = "./build/dist/loadtoad_0.0.6_darwin_amd64_notarized.zip"
}
