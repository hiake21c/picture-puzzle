package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
	"log"
	"picture-puzzle/font"
	"picture-puzzle/global"
	"picture-puzzle/scenemanager"
)

type StartScene struct {
	startImg *ebiten.Image
}

func (s *StartScene) Startup() {
	var err error
	s.startImg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

func (s *StartScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(s.startImg, nil)

	titleWidth := font.TextWidth(global.Title, global.TitleFontScale)
	font.DrawTextWithShadow(screen, global.Title, global.ScreenWidth/2-titleWidth/2, global.ScreenHeight/2, global.TitleFontScale, color.Black)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}
