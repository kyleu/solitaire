import "./game.css";
import {Socket} from "../socket";
import {client, GameClient, setClient} from "./client";
import type {Message} from "./message/message";

function wasmInit(ms: number, ruleset: string) {
  setClient(new GameClient());
  client().init(ms, ruleset);
}

function websocketInit(ms: number, url: string, rs: string) {
  setClient(new GameClient());

  function open() {
    client().init(ms, rs);
  }

  function recv(m) {
    handleMessage(client(), m.param as Message);
  }

  function err(e) {
    console.log("[socket error]: " + e);
  }

  const sock = new Socket(false, open, recv, err, url);

  globalThis.log = (m: unknown) => console.log(m);
  globalThis.sendGameMessage = (t: string, data: { [key: string]: unknown }) => {
    sock.send({channel: "game", cmd: "message", param: {t, data}});
  };
}

function handleMessage(gameClient: GameClient, msg: Message) {
  console.log("received message: " + JSON.stringify(msg));
}

globalThis.wasmInit = wasmInit;
globalThis.websocketInit = websocketInit;

globalThis.onClientMessage = (msg) => handleMessage(client(), msg);
