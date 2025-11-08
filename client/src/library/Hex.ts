// adapted from https://www.Redblobgames.com/grids/hexagons/

import { HexTile } from "./types";

const SQRT_3_2 = 0.8660254037844386;
// hexagons are the bestagons 😻
export class Hex {
  private static readonly DIRECTIONS: [number, number, number][] = [
    [1, 0, -1],
    [1, -1, 0],
    [0, -1, 1],
    [-1, 0, 1],
    [-1, 1, 0],
    [0, 1, -1],
  ];

  public Q: number;
  public R: number;
  public S: number;

  constructor(q: number, r: number) {
    this.Q = q;
    this.R = r;
    this.S = -q - r;
  }

  equals(other: Hex): boolean {
    return this.Q === other.Q && this.R === other.R && this.S === other.S;
  }

  notEquals(other: Hex): boolean {
    return !this.equals(other);
  }

  hashCode(): number {
    return (this.Q * 31 + this.R) * 31 + this.S;
  }

  toString(): string {
    return `Hex(q=${this.Q}, r=${this.R}, s=${this.S})`;
  }

  static getDirection(direction: number): Hex {
    if (direction < 0 || direction >= 6) {
      throw new Error("Direction must be between 0 and 5");
    }
    const dir = Hex.DIRECTIONS[direction];
    return new Hex(dir[0], dir[1]);
  }

  add(other: Hex): Hex {
    return new Hex(this.Q + other.Q, this.R + other.R);
  }

  subtract(other: Hex): Hex {
    return new Hex(this.Q - other.Q, this.R - other.R);
  }

  multiply(scalar: number): Hex {
    return new Hex(this.Q * scalar, this.R * scalar);
  }

  length(): number {
    return (Math.abs(this.Q) + Math.abs(this.R) + Math.abs(this.S)) / 2;
  }

  distanceFrom(other: Hex): number {
    return this.subtract(other).length();
  }

  getNeighbour(direction: number): Hex {
    return this.add(Hex.getDirection(direction));
  }
}

export type Vertex = {
  Q: number;
  R: number;
  Direction: string;
};

export class Point {
  constructor(
    public x: number,
    public y: number
  ) {}
}

export class Orientation {
  public forwardMatrix: number[][];
  public invMatrix: number[][];
  public startAngle: number;

  constructor(
    forwardMatrix: number[][],
    invMatrix: number[][],
    startAngle: number
  ) {
    if (
      !Array.isArray(forwardMatrix) ||
      !forwardMatrix.every(
        (row) =>
          Array.isArray(row) && row.every((val) => typeof val === "number")
      )
    ) {
      throw new Error("forwardMatrix must be a 2D array of numbers");
    }

    if (
      !Array.isArray(invMatrix) ||
      !invMatrix.every(
        (row) =>
          Array.isArray(row) && row.every((val) => typeof val === "number")
      )
    ) {
      throw new Error("invMatrix must be a 2D array of numbers");
    }

    this.forwardMatrix = forwardMatrix;
    this.invMatrix = invMatrix;
    this.startAngle = startAngle;
  }
}

export class Layout {
  public orientation: Orientation;
  public size: Point;
  public origin: Point;

  constructor(orientation: Orientation, size: Point, origin: Point) {
    this.orientation = orientation;
    this.size = size;
    this.origin = origin;
  }
}

export function hexToPixel(layout: Layout, h: Hex | HexTile): Point {
  const M = layout.orientation.forwardMatrix;
  const x = (M[0][0] * h.Q + M[0][1] * h.R) * layout.size.x;
  const y = (M[1][0] * h.Q + M[1][1] * h.R) * layout.size.y;

  return new Point(x + layout.origin.x, y + layout.origin.y);
}

export function hexCornerOffset(layout: Layout, corner: number): Point {
  const size = layout.size;
  const angle = (2.0 * Math.PI * (layout.orientation.startAngle + corner)) / 6;

  return new Point(size.x * Math.cos(angle), size.y * Math.sin(angle));
}

export function vertexToPixel(layout: Layout, v: Vertex): Point {
  const M = layout.orientation.forwardMatrix;

  const x = (M[0][0] * v.Q + M[0][1] * v.R) * layout.size.x;
  let y = (M[1][0] * v.Q + M[1][1] * v.R) * layout.size.y;

  if (v.Direction === "S") {
    y += layout.size.y;
  } else {
    y -= layout.size.y;
  }

  return new Point(x + layout.origin.x, y + layout.origin.y);
}
