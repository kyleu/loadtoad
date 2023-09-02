// Content managed by Project Forge, see [projectforge.md] for details.
import SwiftUI
import LoadtoadServer

@main
struct Project: App {
    init() {
        print("starting Load Toad...")
        let path = NSSearchPathForDirectoriesInDomains(.libraryDirectory, .userDomainMask, true)
        let port = LoadtoadServer.CmdLib(path[0])
        print("Load Toad started on port [\(port)]")
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
