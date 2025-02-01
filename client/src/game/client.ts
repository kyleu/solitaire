import "./game.css";
import {opt, req} from "../dom";
import {appendLog} from "./util";

declare const sendGameMessage: (t: string, data: { [key: string]: unknown }) => void;
declare const log: (msg: string) => void;

export class GameClient {
  init(ms: number, st: string) {
    opt("#loading-panel")?.remove();
    req("#load-status").innerText = "Loaded in [" + ms + "ms]";
    log(`game [${st}] initialized in [${ms}ms]`);
    sendGameMessage("welcome", {ruleset: st});
  }

  onInput(key: string, msg: unknown) {
    switch (key) {
      case "save":
        sendGameMessage("save", {});
        break;
      case "testbed":
        appendLog(JSON.stringify(msg, null, 2));
        break;
      default:
        console.log("unhandled input [" + key + "].");
        break;
    }
  }
}

let globalClient: GameClient;

export function setClient(cl: GameClient) {
  globalClient = cl;
}

export function client(): GameClient {
  return globalClient;
}
