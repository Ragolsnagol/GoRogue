package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/Ragolsnagol/GoRogue/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var hudImg *ebiten.Image = nil
var hudErr error = nil
var hudFont text.Face = nil

func ProcessHUD(g *Game, screen *ebiten.Image) {
	if hudImg == nil {
		hudImg, _, hudErr = ebitenutil.NewImageFromFile("assets/UIPanel.png")
		if hudErr != nil {
			log.Fatal(hudErr)
		}
	}
	if hudFont == nil {
		tt, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
		if err != nil {
			log.Fatal(err)
		}
		hudFont = &text.GoTextFace{
			Source: tt,
			Size:   16,
		}
	}
	gd := NewGameData()

	uiY := (gd.ScreenHeight - gd.UIHeight) * gd.TileHeight
	uiX := (gd.ScreenWidth * gd.TileWidth) / 2
	var fontX = uiX + 16
	var fontY = uiY + 24
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(uiX), float64(uiY))
	screen.DrawImage(userLogImg, op)

	for _, p := range g.World.Query(g.WorldTags["players"]) {
		h := p.Components[health].(*Health)
		healthText := fmt.Sprintf("Health: %d / %d", h.CurrentHealth, h.MaxHealth)
		drawHUDText(screen, hudFont, healthText, fontX, fontY)
		fontY += 16
		ac := p.Components[armor].(*Armor)
		acText := fmt.Sprintf("Armor Class: %d", ac.ArmorClass)
		drawHUDText(screen, hudFont, acText, fontX, fontY)
		fontY += 16
		defText := fmt.Sprintf("Defense: %d", ac.Defense)
		drawHUDText(screen, hudFont, defText, fontX, fontY)
		fontY += 16
		wpn := p.Components[meleeWeapon].(*MeleeWeapon)
		dmg := fmt.Sprintf("Damage: %d - %d", wpn.MinimumDamage, wpn.MaximumDamage)
		drawHUDText(screen, hudFont, dmg, fontX, fontY)
		fontY += 16
		bonus := fmt.Sprintf("To Hit Bonus: %d", wpn.ToHitBonus)
		drawHUDText(screen, hudFont, bonus, fontX, fontY)
	}
}

func drawHUDText(screen *ebiten.Image, font text.Face, msg string, x int, y int) {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, msg, font, options)
}
