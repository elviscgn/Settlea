import * as PIXI from "pixi.js";
import {
  Hex,
  Layout,
  Orientation,
  Point,
  hexToPixel,
  vertexToPixel,
} from "@/library/Hex";

import type { Corner, HexTile, PortData, PortRender } from "@/library/types";
import config from "@/config";

let currentContainer: PIXI.Container | null = null;

function drawSquare(
  container: PIXI.Container,
  cornerContainer: PIXI.Container,
  texture: PIXI.Texture
): void {
  const settlePopup = new PIXI.Container();
  settlePopup.interactive = true;
  settlePopup.cursor = "pointer";

  const point = cornerContainer.position;

  const coords = {
    x: point.x - 40,
    y: point.y - 120,
  };
  const square = new PIXI.Graphics();
  const structureSprite = PIXI.Sprite.from(texture);
  structureSprite.position.set(coords.x + 15, coords.y + 10);
  structureSprite.scale.set(0.25);

  square.roundRect(coords.x, coords.y, 80, 80);
  square.fill(0xf8f2dc);
  square.stroke({ width: 5, color: 0x1d5d64 });

  currentContainer = settlePopup;

  settlePopup.on("mouseover", () => {
    if (currentContainer) {
      // odd bug where settlePop container freaks out and gets sent to the origin so i went w this :c
      // structureSprite.position.set(coords.x + 18, coords.y + 20);
      // structureSprite.scale.set(0.23);
      square.clear();
      square.roundRect(coords.x + 7.5, coords.y + 5, 75, 75);
      square.fill(0xf8f2dc);
      square.stroke({ width: 5, color: 0x1d5d64 });
    }
  });

  settlePopup.on("mouseout", () => {
    if (currentContainer) {
      structureSprite.position.set(coords.x + 15, coords.y + 10);
      structureSprite.scale.set(0.25);
      square.clear();
      square.roundRect(coords.x, coords.y, 80, 80);
      square.fill(0xf8f2dc);
      square.stroke({ width: 5, color: 0x1d5d64 });
    }
  });

  settlePopup.on("click", () => {
    if (currentContainer) {
      console.log("clicked on");
      removeCurrentSquare(container);
    }
  });

  settlePopup.addChild(square);
  settlePopup.addChild(structureSprite);
  container.addChild(settlePopup);
}

function removeCurrentSquare(container: PIXI.Container): void {
  if (currentContainer) {
    container.removeChild(currentContainer);
    currentContainer.destroy();
    currentContainer = null;
  }
}

export async function generateMap(
  app: PIXI.Application,
  hexMap: HexTile[],
  hexCorners: Corner[],
  ports: Record<string, PortData>,
  layoutPointy: Orientation,
  textures: Record<string, PIXI.Texture>,
  container: PIXI.Container
): Promise<void> {
  // Fonts

  const map: Set<Hex> = new Set();
  const N = 2;

  const centerX = (config.width * window.devicePixelRatio) / 2;
  const centerY = (config.height * window.devicePixelRatio) / 2;

  const layout = new Layout(
    layoutPointy,
    new Point(92, 92),
    new Point(centerX, centerY)
  );
  const graphics = new PIXI.Graphics();

  const hexLookup = (q: number, r: number, s: number) => {
    return hexMap.find((tile) => tile.Q === q && tile.R === r && tile.S === s);
  };

  for (let q = -N; q <= N; q++) {
    for (let r = Math.max(-N, -q - N); r <= Math.min(N, -q + N); r++) {
      map.add(new Hex(q, r));
    }
  }

  for (const hex of hexMap) {
    const tileContainer = new PIXI.Container();

    const hexData = hexLookup(hex.Q, hex.R, hex.S);
    if (!hexData) {
      console.error(
        `No data found for hex at q=${hex.Q}, r=${hex.R}, s=${hex.S}`
      );
      continue;
    }

    const randTile = new PIXI.Sprite(textures[`${hexData.Type}_tile`]);

    const tileCenter: Point = hexToPixel(layout, hex);
    randTile.anchor.set(0.5);
    randTile.position.set(tileCenter.x, tileCenter.y);
    randTile.scale.set(0.47);

    tileContainer.interactive = true;

    tileContainer.addChild(randTile);

    if (hex.Type !== "desert") {
      const tokenNumber = new PIXI.Text({
        text: hex.Token.toString(),
        style: {
          fontFamily: "Bungee, sans-serif",
          fontSize: 30,
          fill: hex.Token === 6 || hex.Token === 8 ? "#FF0000" : "#183A37",
        },
      });

      const tokenPip = new PIXI.Text({
        text: "•".repeat(hex.Token === 7 ? 0 : 6 - Math.abs(7 - hex.Token)),

        style: {
          fontSize: 20,
          // red if token is 6 or 8
          fill: hex.Token === 6 || hex.Token === 8 ? "#FF0000" : "#183A37",
        },
      });

      tokenNumber.anchor.set(0.5);
      tokenPip.anchor.set(0.5);
      tokenNumber.position.set(tileCenter.x, tileCenter.y + 22);
      tokenPip.position.set(tileCenter.x, tileCenter.y + 45);

      tileContainer.addChild(tokenNumber);
      tileContainer.addChild(tokenPip);
    }

    container.addChild(tileContainer);

    tileContainer.on("click", () => {
      removeCurrentSquare(container);
    });
  }

  const portTextures: Record<string, PIXI.Sprite> = {
    wood: PIXI.Sprite.from(textures.wood),
    brick: PIXI.Sprite.from(textures.brick),
    sheep: PIXI.Sprite.from(textures.sheep),
    ore: PIXI.Sprite.from(textures.ore),
    wheat: PIXI.Sprite.from(textures.wheat),
  };

  const portMappings: Record<string, PortRender> = {};

  for (const [key, value] of Object.entries(ports)) {
    portMappings[key] = {
      exchangeRate: {
        text: value.exchangeRate.text.toString(),
        coord: value.exchangeRate.coord,
      },
      portType: {
        text:
          value.portType.text === "?" ? "?" : portTextures[value.portType.text],
        coord: value.portType.coord,
        size: value.portType.size, // Include size if it exists
      },
    };
  }

  for (const portName in portMappings) {
    addPorts(portMappings[portName], container);
  }

  const cornerContainers: PIXI.Container[] = [];
  for (const corner of hexCorners) {
    const point = vertexToPixel(layout, corner);
    const circle = new PIXI.Graphics();
    const cornerContainer = new PIXI.Container();
    const drawCirc = (alpha: number) => {
      circle.clear();
      circle.circle(0, 0, 21);
      circle.fill({ color: 0xffffff, alpha: alpha });
      circle.stroke({ color: 0x000 });
    };

    cornerContainer.interactive = true;
    cornerContainer.cursor = "pointer";

    drawCirc(0.2);

    cornerContainer.position.set(point.x, point.y);
    cornerContainer.addChild(circle);

    cornerContainers.push(cornerContainer);
    container.addChild(cornerContainer);

    cornerContainer.on("mouseover", () => {
      drawCirc(0.8);
    });

    cornerContainer.on("mouseout", () => {
      drawCirc(0.2);
    });

    cornerContainer.on("click", () => {
      removeCurrentSquare(container);
      drawSquare(container, cornerContainer, textures["settlea_red"]);
    });

    container.on("click", (e) => {
      //   removeCurrentSquare(container);
      console.log("click in", container.uid, cornerContainer.uid);
      e.stopPropagation();
    });
  }

  const minScale = 0.8;
  const maxScale = 1.2;
  const scaleSpeed = 0.008;

  app.ticker.add(() => {
    cornerContainers.forEach((container) => {
      if (!("scaleDirection" in container)) {
        (container as any).scaleDirection = 1;
      }

      const scaleDirection = (container as any).scaleDirection;

      container.scale.x += scaleDirection * scaleSpeed;
      container.scale.y += scaleDirection * scaleSpeed;

      if (container.scale.x >= maxScale || container.scale.x <= minScale) {
        (container as any).scaleDirection *= -1;
      }
    });
  });

  container.addChild(graphics);
}

export function addPorts(
  portData: PortRender,
  container: PIXI.Container
): void {
  const ratesText = new PIXI.Text({
    text: portData.exchangeRate.text,
    style: { fontFamily: "Rubik, Arial, sans-serif" },
  });

  ratesText.x = portData.exchangeRate.coord[0];
  ratesText.y = portData.exchangeRate.coord[1];
  ratesText.style.fontSize = 20;
  container.addChild(ratesText);

  // if its a sprite
  if (portData.portType.text instanceof PIXI.Sprite) {
    const typeSprite = portData.portType.text as PIXI.Sprite;
    typeSprite.anchor.set(0.5);
    typeSprite.position.set(
      portData.portType.coord[0],
      portData.portType.coord[1]
    );

    typeSprite.scale.set(portData.portType.size);
    container.addChild(typeSprite);
  } else {
    const typeText = new PIXI.Text({
      text: portData.portType.text,
      style: { fontFamily: "Rubik" },
    });

    typeText.x = portData.portType.coord[0];
    typeText.y = portData.portType.coord[1];
    container.addChild(typeText);
  }
}

export function setupBackground(
  texture: PIXI.Texture,
  container: PIXI.Container
): void {
  texture.source.scaleMode = "linear";
  const background = new PIXI.Sprite(texture);
  background.anchor.set(0.5);
  background.position.set(
    (config.width * window.devicePixelRatio) / 2,
    (config.height * window.devicePixelRatio) / 2
  );
  background.scale.set(0.65);
  container.addChild(background);
}

export function centerCanvas(app: PIXI.Application): void {
  app.canvas.style.width = `${window.innerWidth}px`;
  app.canvas.style.height = `${window.innerHeight}px`;
  app.canvas.style.position = "absolute";
  // app.canvas.style.top = "50%";
  // app.canvas.style.left = "50%";
  // app.canvas.style.transform = "translate(-50%, -50%)";
}
