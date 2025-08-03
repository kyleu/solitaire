import type { MessageType } from "./messageType";

export class Message {
  id?: string;
  t: MessageType;
  data: { [key: string]: unknown };
  occurred?: Date;
}

export type Messages = Array<Message>;
