package main

type Missle struct {
  BaseRange int
  BaseDamage int
  AttackSpeed int
  Cost int
}

var WoodenArcherMissle = &Missle{5, 10, 10, 80}
var LongBowArrowMissle = &Missle{10, 25, 7, 200}
var CobraSlingShotMissle = &Missle{7, 14, 10, 500}

func (missle *Missle) Clone() *Missle {
  return &Missle{
    missle.BaseRange,
    missle.BaseDamage,
    missle.AttackSpeed,
    missle.Cost,
  }
}
