# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_amd64_v1/solitaire"]
bundle_id = "dev.kyleu.solitaire"

notarize {
  path = "./build/dist/solitaire_0.0.1_darwin_amd64_desktop.dmg"
  bundle_id = "dev.kyleu.solitaire"
}

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/solitaire_0.0.1_darwin_amd64.dmg"
  volume_name = "Solitaire"
}

zip {
  output_path = "./build/dist/solitaire_0.0.1_darwin_amd64_notarized.zip"
}
