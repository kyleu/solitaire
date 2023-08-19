# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_all/solitaire"]
bundle_id = "dev.kyleu.solitaire"

notarize {
  path = "./build/dist/solitaire_0.0.13_darwin_all_desktop.dmg"
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
  output_path = "./build/dist/solitaire_0.0.13_darwin_all.dmg"
  volume_name = "Solitaire"
}

zip {
  output_path = "./build/dist/solitaire_0.0.13_darwin_all_notarized.zip"
}
