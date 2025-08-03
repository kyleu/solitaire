import { startGame } from "./game";

export function appInit(): void {
  // eslint-disable-line @typescript-eslint/no-empty-function
}

declare global {
  interface Window {
    // eslint-disable-line @typescript-eslint/consistent-type-definitions
    gameInit: () => void;
  }
}

export function gameInit(): void {
  startGame();
}

window.gameInit = gameInit;
