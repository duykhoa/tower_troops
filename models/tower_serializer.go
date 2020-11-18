package models

type TowerSerializer struct {
  HPLevel int             `json:"hpLevel"`
  HP int                  `json:"hp"`
  ArmorLevel int          `json:"armorLevel"`
  Armor int               `json:"armor"`
  Troops []Troop          `json:"troops"`
  Golds int               `json:"golds"`
  Missles []Missle        `json:"missles"`
}
