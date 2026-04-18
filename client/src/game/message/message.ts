import type { MessageType } from "./messageType";

export class Message {
  id?: string;
  t!: MessageType;
  data!: Record<string, unknown>;
  occurred?: Date;
}

export type Messages = Message[];
