package main

/* 策略模式
 */

// WeaponStrategy 武器策略(抽象的策略)
type WeaponStrategy interface {
	UseWeapon()
}

// Ak47 具体的策略
type Ak47 struct{}

func (a *Ak47) UseWeapon() {
	println("use ak47")
}

type Knife struct{}

func (k *Knife) UseWeapon() {
	println("use Knife")
}

// Hero 环境类
type Hero struct {
	strategy WeaponStrategy // 拥有一个具体的抽象策略
}

func (h *Hero) SetStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon() // 调用策略
}

func main() {
	hero := Hero{}
	hero.SetStrategy(&Ak47{})
	hero.Fight()
	hero.SetStrategy(&Knife{})
	hero.Fight()
}
