export enum Behavior {
  Stock = "stock",
  Waste = "waste",
  Pocket = "pocket",
  Foundation = "foundation",
  Tableau = "tableau",
  Cell = "cell",
  Reserve = "reserve",
  Pyramid = "pyramid"
}

export function parseBehavior(value: string): Behavior | undefined {
  if (Object.values(Behavior).includes(value as Behavior)) {
    return value as Behavior;
  }
  return undefined;
}

export function getBehavior(value: string): Behavior {
  const x = parseBehavior(value);
  if (x === undefined) {
    throw new Error(`invalid [Behavior]: ${value}`);
  }
  return x;
}
