import * as PIXI from "pixi.js";

export async function loadAllAssets() {
  const assets = [
    { alias: "background", src: "assets/temp_background.png" },
    { alias: "wood", src: "assets/icons/wood.png" },
    { alias: "brick", src: "assets/icons/brick.png" },
    { alias: "sheep", src: "assets/icons/sheep.png" },
    { alias: "wheat", src: "assets/icons/wheat.png" },
    { alias: "ore", src: "assets/icons/rock.png" },
    { alias: "wood_tile", src: "assets/tiles/wood.png" },
    { alias: "brick_tile", src: "assets/tiles/brick.png" },
    { alias: "sheep_tile", src: "assets/tiles/sheep.png" },
    { alias: "wheat_tile", src: "assets/tiles/wheat.png" },
    { alias: "ore_tile", src: "assets/tiles/ore.png" },
    { alias: "desert_tile", src: "assets/tiles/desert.png" },
    { alias: "settlea_red", src: "assets/structures/settlea_red.svg" },
    { alias: "Bungee", type: "font", src: "assets/fonts/bungee.woff2" },
    { alias: "Rubik", type: "font", src: "assets/fonts/rubik.woff2" },
  ];

  assets.forEach(({ alias, src }) => PIXI.Assets.add({ alias, src }));

  return PIXI.Assets.load(assets.map(({ alias }) => alias));
}
