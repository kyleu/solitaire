# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_arm64/solitaire"]
bundle_id = "dev.kyleu.solitaire"

notarize {
  path = "./build/dist/solitaire_0.1.0_darwin_arm64_desktop.dmg"
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
  output_path = "./build/dist/solitaire_0.1.0_darwin_arm64.dmg"
  volume_name = "Solitaire"
}

zip {
  output_path = "./build/dist/solitaire_0.1.0_darwin_arm64_notarized.zip"
}
