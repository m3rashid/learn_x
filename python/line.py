import math
from __future__ import annotations


class Line:
  def __init__(self, x0: float, y0: float, z0: float, constant: float) -> None:
    self.x0 = x0
    self.y0 = y0
    self.z0 = z0
    self.constant = constant
    self.equation = f"({x0})x + ({y0})y + ({z0})z + ({constant}) = 0" 

  def __eq__(self, other: Line) -> bool:
    pass 

  def getDistanceFromOrigin(self) -> float:
    pass

  def getDistance(self, other:Line) -> float:
    pass

  def getSlope(self) -> float:
    pass

  