export enum Color {
  Red = "red",
  Black = "black",
  Green = "green",
  Blue = "blue",
  Colorless = "colorless",
  Unknown = "unknown"
}

export function parseColor(value: string): Color | undefined {
  if (Object.values(Color).includes(value as Color)) {
    return value as Color;
  }
  return undefined;
}

export function getColor(value: string): Color {
  const x = parseColor(value);
  if (x === undefined) {
    throw new Error(`invalid [Color]: ${value}`);
  }
  return x;
}
