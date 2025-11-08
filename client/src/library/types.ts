import * as PIXI from "pixi.js";

export interface PortData {
  exchangeRate: {
    text: string;
    coord: [number, number];
  };
  portType: {
    text: string;
    coord: [number, number];
    size?: number;
  };
}

export interface PortRender {
  exchangeRate: {
    text: string;
    coord: [number, number];
  };
  portType: {
    text: PIXI.Sprite | string;
    coord: [number, number];
    size?: number;
  };
}

type ResourceType = "sheep" | "wood" | "brick" | "ore" | "wheat" | "desert";

export interface HexTile {
  Q: number;
  R: number;
  S: number;
  Type: ResourceType;
  Token: number;
  Blocked: boolean;
  Coords: ScreenCoord;
}

export interface Structure {
  Type: string;
  OwnerId: number;
}

export interface Corner {
  Q: number;
  R: number;
  Direction: string;
  Coords: ScreenCoord;
  Structure: Structure | null;
  IsPort: boolean;
  PortType: string;
}

export interface ScreenCoord {
  X: number;
  Y: number;
}

export interface HexMap {
  hex_map: HexTile[];
}

export interface TileAPIResponse {
  tiles: HexTile[];
  corners: Corner[];
  // edges: Edge[];
  ports: Record<string, PortData>;

  iterations: number;
  duration: string;
}

export interface Room {
  id: string;
  name: string;
}

export interface Client {
  id: string;
  name: string;
}

export interface MessageResponse {
  action: string;
  content: string;
  roomID: string;
  username: string;
}

/*
type Message struct {
	Action string  `json:"action"`
	Data   string  `json:"data"`
	Target *Room   `json:"target"`
	Sender *Client `json:"sender"`
}

*/
