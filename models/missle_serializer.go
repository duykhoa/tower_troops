package models

type MissleSerializer struct {
  Damage int           `json:"damage"`
  DamageLevel int      `json:"damageLevel"`
  Range int            `json:"range"`
  RangeLevel int       `json:"rangeLevel"`
  AttackSpeed int      `json:"attackSpeed"`
  AttackSpeedLevel int `json:"attackSpeedLevel"`
  Name string          `json:"name"`
}
