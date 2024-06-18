import SwiftUI
import SolitaireServer

@main
struct Project: App {
    init() {
        print("starting Solitaire...")
        let path = NSSearchPathForDirectoriesInDomains(.libraryDirectory, .userDomainMask, true)
        let port = SolitaireServer.CmdLib(path[0])
        print("Solitaire started on port [\(port)]")
        let url = URL.init(string: "http://localhost:\(port)/")!
        self.cv = ContentView(url: URLRequest(url: url))
    }

    var cv: ContentView

    var body: some Scene {
        WindowGroup {
            cv
        }
    }
}
