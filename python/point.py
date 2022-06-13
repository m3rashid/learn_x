import math
from __future__ import annotations


class Point:
  def __init__(self, x, y, z) -> None:
    self.x = x
    self.y = y
    self.z = z

  def __eq__(self, other: Point) -> bool:
    return self.x == other.x and self.y == other.y and self.z == other.z

  def __add__(self, other: Point) -> Point:
    return Point(self.x + other.x, self.y + other.y, self.z + other.z)

  def __str__(self) -> str:
    return f"({self.x}, {self.y}, {self.z})"

  def getDistance(self, other: Point) -> float:
    a = math.pow((self.x - other.x))
    b = math.pow((self.y - other.y))
    c = math.pow((self.z - other.z))
    return math.sqrt(a + b + c)

  def getDistanceFromOrigin(self) -> float:
    return math.sqrt(math.pow(self.x, 2) + math.pow(self.y, 2)+ math.pow(self.z, 2))

  def getEquationOfLine(self, other: Point) -> str:
    pass

