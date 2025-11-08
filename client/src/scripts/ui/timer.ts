import * as PIXI from "pixi.js";

function formatTime(time: number): string {
  const minutes = Math.floor(time / 60);
  const seconds = time % 60;
  return `${minutes}:${seconds < 10 ? "0" : ""}${seconds}`;
}

export function drawTimer(app: PIXI.Application) {
  const countdownTime = 5 * 60;
  let timeRemaining = countdownTime;

  const timerContainer = new PIXI.Container();
  const timerText = new PIXI.Text({
    text: "05:00",
    style: {
      fontFamily: "Bungee, sans-serif",
      fontSize: 35,
      fill: 0xffffff,
    },
  });
  const timerBorder = new PIXI.Graphics();

  const startX = app.renderer.screen.width * (1 - 0.31);
  const startY = app.renderer.screen.height - 101;
  const borderWidth = app.renderer.screen.width * 0.1;
  const borderHeight = 100;

  timerBorder.roundRect(startX, startY, borderWidth, borderHeight, 10);
  timerBorder.fill(0x177e89);
  timerBorder.stroke({ width: 7, color: 0xffffff });

  timerText.anchor.set(0.5);
  timerText.x = startX + borderWidth / 2;
  timerText.y = startY + borderHeight / 2;

  timerContainer.addChild(timerBorder);
  timerContainer.addChild(timerText);

  let elapsedTime = 0;

  function updateCountdown(ticker: PIXI.Ticker) {
    elapsedTime += ticker.elapsedMS / 1000;

    if (elapsedTime >= 1) {
      elapsedTime = 0;

      if (timeRemaining > 0) {
        timeRemaining--;
        timerText.text = formatTime(timeRemaining);
      } else {
        timerText.text = "00:00";
        app.ticker.remove(updateCountdown);
      }
    }
  }

  app.ticker.add(updateCountdown);
  app.stage.addChild(timerContainer);

  console.log(timerContainer.x);
}
