import "./game.css";
import { opt } from "../dom";

export function appendLog(msg: string) {
  const panel = opt<HTMLDivElement>("#log-panel");
  if (panel) {
    const row = document.createElement("tr");
    const numTH = document.createElement("th");
    numTH.innerText = panel.children.length.toString();
    const textTD = document.createElement("td");
    textTD.innerText = msg;
    row.append(numTH, textTD);
    panel.append(row);
    const c = document.getElementById("content");
    c.scrollTo(0, c.scrollHeight);
  } else {
    console.log("[no-log-panel]: " + msg);
  }
}
