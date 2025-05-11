package main

import (
	"fmt"
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.name = []byte(name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.coordinates = struct {
			x int32
			y int32
			z int32
		}{
			x: int32(x),
			y: int32(y),
			z: int32(z),
		}
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = int32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(mana) << (64 - 12))
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(health) << (64 - 22))
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(respect) << (64 - 26))
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(strength) << (64 - 30))
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(experience) << (64 - 34))
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(level) << (64 - 38))
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(1) << (64 - 39))
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(1) << (64 - 40))
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(1) << (64 - 41))
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= (int64(personType) << (64 - 43))
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

// Магическая сила (мана) [0…1000] значений
// Здоровье [0…1000] значений
// Уважение [0…10] значений
// Сила [0…10] значений
// Опыт [0…10] значений
// Уровень [0…10] значений
// Есть ли у игрока дом [true/false] значения
// Есть ли у игрока оружие [true/false] значения
// Есть ли у игрока семья [true/false] значения
// Тип игрока [строитель/кузнец/воин] значения

// 10bit, 10bit, 4bit, 4bit, 4bit, 4bit, 1bit, 1bit, 1bit, 2bit,

type GamePerson struct {
	name        []byte
	coordinates struct {
		x int32
		y int32
		z int32
	}
	gold  int32
	stats int64
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}
	for _, opt := range options {
		opt(&person)
	}
	return person
}

func (p *GamePerson) Name() string {
	return string(p.name)
}

func (p *GamePerson) X() int {
	return int(p.coordinates.x)
}

func (p *GamePerson) Y() int {
	return int(p.coordinates.y)
}

func (p *GamePerson) Z() int {
	return int(p.coordinates.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int((p.stats & (int64(2047) << (64 - 12))) >> (64 - 12))
}

func (p *GamePerson) Health() int {
	return int((p.stats & (int64(2047) << (64 - 22))) >> (64 - 22))
}

func (p *GamePerson) Respect() int {
	return int((p.stats & (int64(15) << (64 - 26))) >> (64 - 26))
}

func (p *GamePerson) Strength() int {
	return int((p.stats & (int64(15) << (64 - 30))) >> (64 - 30))
}

func (p *GamePerson) Experience() int {
	return int((p.stats & (int64(15) << (64 - 34))) >> (64 - 34))
}

func (p *GamePerson) Level() int {
	return int((p.stats & (int64(15) << (64 - 38))) >> (64 - 38))
}

func (p *GamePerson) HasHouse() bool {
	return int((p.stats&(int64(1)<<(64-39)))>>(64-39)) == 1
}

func (p *GamePerson) HasGun() bool {
	return int((p.stats&(int64(1)<<(64-40)))>>(64-40)) == 1
}

func (p *GamePerson) HasFamilty() bool {
	return int((p.stats&(int64(1)<<(64-41)))>>(64-41)) == 1
}

func (p *GamePerson) Type() int {
	fmt.Printf("%b %b", p.stats, (int64(3) << (64 - 43)))
	return int((p.stats & (int64(3) << (64 - 43))) >> (64 - 43))
}

func TestGamePerson(t *testing.T) {
	assert.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

	const x, y, z = math.MinInt32, math.MaxInt32, 0
	const name = "aaaaaaaaaaaaa_bbbbbbbbbbbbb_cccccccccccccc"
	const personType = BuilderGamePersonType
	const gold = math.MaxInt32
	const mana = 1000
	const health = 1000
	const respect = 10
	const strength = 10
	const experience = 10
	const level = 10

	options := []Option{
		WithName(name),
		WithCoordinates(x, y, z),
		WithGold(gold),
		WithMana(mana),
		WithHealth(health),
		WithRespect(respect),
		WithStrength(strength),
		WithExperience(experience),
		WithLevel(level),
		WithHouse(),
		WithFamily(),
		WithType(personType),
	}

	person := NewGamePerson(options...)
	assert.Equal(t, name, person.Name())
	assert.Equal(t, x, person.X())
	assert.Equal(t, y, person.Y())
	assert.Equal(t, z, person.Z())
	assert.Equal(t, gold, person.Gold())
	assert.Equal(t, mana, person.Mana())
	assert.Equal(t, health, person.Health())
	assert.Equal(t, respect, person.Respect())
	assert.Equal(t, strength, person.Strength())
	assert.Equal(t, experience, person.Experience())
	assert.Equal(t, level, person.Level())
	assert.True(t, person.HasHouse())
	assert.True(t, person.HasFamilty())
	assert.False(t, person.HasGun())
	assert.Equal(t, personType, person.Type())
}
