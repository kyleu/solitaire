import "./game.css";
import { Socket } from "../socket";
import type { SocketMessage } from "../socket";
import { client, GameClient, setClient } from "./client";
import type { Message } from "./message/message";

declare global {
  var log: (msg: unknown) => void;
  var sendGameMessage: (t: string, data: Record<string, unknown>) => void;
  var wasmInit: (ms: number, ruleset: string) => void;
  var websocketInit: (ms: number, url: string, rs: string) => void;
  var onClientMessage: (msg: Message) => void;
}

function wasmInit(ms: number, ruleset: string) {
  setClient(new GameClient());
  client().init(ms, ruleset);
}

function websocketInit(ms: number, url: string, rs: string) {
  setClient(new GameClient());

  function open() {
    client().init(ms, rs);
  }

  function recv(m: SocketMessage) {
    handleMessage(m.param as unknown as Message);
  }

  function err(svc: string, err: string) {
    console.log(`[${svc} error]: ${err}`);
  }

  const sock = new Socket(false, open, recv, err, url);

  globalThis.log = (m: unknown) => console.log(m);
  globalThis.sendGameMessage = (t: string, data: Record<string, unknown>) => {
    sock.send({ channel: "game", cmd: "message", param: { t, data } });
  };
}

function handleMessage(msg: Message) {
  console.log("received message: " + JSON.stringify(msg));
}

globalThis.wasmInit = wasmInit;
globalThis.websocketInit = websocketInit;

globalThis.onClientMessage = (msg: Message) => handleMessage(msg);
