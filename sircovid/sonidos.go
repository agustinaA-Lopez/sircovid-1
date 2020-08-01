package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	vol  float64
	temp = float64(vol)
	up   bool
	down bool
	fade bool
	// sonido
	audioContext *audio.Context
	deadSound    *audio.Player
	deadSound2   *audio.Player
	sonidoFondo  *audio.InfiniteLoop
	fondo        *audio.Player
	sonidoIntro  *audio.InfiniteLoop
	sIntro       *audio.Player
	sPuerta      *audio.Player
	sDinero      *audio.Player
	sNube        *audio.Player
	sFast        *audio.Player
	sBarbijo     *audio.Player
	sLevelUp     *audio.Player
)

// Inicio valores de sonido del juego
func initSonido() {

	audioContext, _ = audio.NewContext(44100)
	// sonido fondo
	s, err := os.Open(`sircovid\data\audio\SIR-COVID sin moneditas (1).wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data := make([]byte, 11491248)
	c, err := s.Read(data)
	// fmt.Println(c)

	fondoD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sonidoFondo = audio.NewInfiniteLoop(fondoD, int64(c))
	if err != nil {
		log.Fatal(err)
	}
	fondo, err = audio.NewPlayer(audioContext, sonidoFondo)
	if err != nil {
		log.Fatal(err)
	}
	// sonido intro
	s, err = os.Open(`sircovid\data\audio\introconteclas.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 8178592)
	c, err = s.Read(data)
	// fmt.Println(c)
	introD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sonidoIntro = audio.NewInfiniteLoop(introD, int64(c))
	if err != nil {
		log.Fatal(err)
	}
	sIntro, err = audio.NewPlayer(audioContext, sonidoIntro)
	if err != nil {
		log.Fatal(err)
	}

	//sonido Puerta
	s, err = os.Open(`sircovid\data\audio\puertas ingresos.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 86668)
	c, err = s.Read(data)
	// fmt.Println(c)

	puertaD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sPuerta, err = audio.NewPlayer(audioContext, puertaD)
	if err != nil {
		log.Fatal(err)
	}

	//sonido Monedas
	s, err = os.Open(`sircovid\data\audio\DINERO.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 194948)
	c, err = s.Read(data)
	// fmt.Println(c)

	dineroD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sDinero, err = audio.NewPlayer(audioContext, dineroD)
	if err != nil {
		log.Fatal(err)
	}
	// sonido Fast
	s, err = os.Open(`sircovid\data\audio\ALRIGHT! COFFE.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 151636)
	c, err = s.Read(data)
	// fmt.Println(c)

	fastD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sFast, err = audio.NewPlayer(audioContext, fastD)
	if err != nil {
		log.Fatal(err)
	}

	// sonido barbijo o alcohol

	s, err = os.Open(`sircovid\data\audio\ponerse barbijo.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 151636)
	c, err = s.Read(data)
	// fmt.Println(c)

	barbijoD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sBarbijo, err = audio.NewPlayer(audioContext, barbijoD)
	if err != nil {
		log.Fatal(err)
	}

	//sonido Pasar Nivel

	s, err = os.Open(`sircovid\data\audio\PASAR DE NIVEL.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 99346536)
	c, err = s.Read(data)

	levelD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sLevelUp, err = audio.NewPlayer(audioContext, levelD)
	if err != nil {
		log.Fatal(err)
	}

	// sonido perder vida
	s, err = os.Open(`sircovid\data\audio\tos1.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 346536)
	c, err = s.Read(data)

	tosD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	deadSound, err = audio.NewPlayer(audioContext, tosD)
	if err != nil {
		log.Fatal(err)
	}

	// sonido muerte
	s, err = os.Open(`sircovid\data\audio\sonido muerte o daño por nube.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 178704)
	_, err = s.Read(data)
	// fmt.Println(c)
	jabD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	deadSound2, err = audio.NewPlayer(audioContext, jabD)
	if err != nil {
		log.Fatal(err)
	}

}

func sonido() {

	// volumen on/off
	if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		switch {
		case vol != .01:
			vol = .01
		case vol == .01:
			vol = temp
		}
	}

	// volumen +/-
	if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd) || inpututil.IsKeyJustPressed(ebiten.Key9) {
		up = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyKPAdd) || inpututil.IsKeyJustReleased(ebiten.Key9) {
		up = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract) || inpututil.IsKeyJustPressed(ebiten.Key8) {
		down = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyKPSubtract) || inpututil.IsKeyJustReleased(ebiten.Key8) {
		down = false
	}
	switch {
	case vol < .99 && up:
		vol += .01
	case vol > .01 && down:
		vol -= .01
	}
	fmt.Println(vol)

	fondo.SetVolume(vol)
	deadSound.SetVolume(vol)
	deadSound2.SetVolume(vol)
	sBarbijo.SetVolume(vol)
	sDinero.SetVolume(vol)
	sFast.SetVolume(vol)
	sIntro.SetVolume(vol)
	sLevelUp.SetVolume(vol)
	sPuerta.SetVolume(vol)

	// sonido ModePause
	if ModePause {
		fondo.Pause()
	}

<<<<<<< HEAD
	if ModeTitle || ModeGame {
		// fadeIn()
=======
	if ModeTitle >= 0 {
		fadeIn()

>>>>>>> dd54ccb39fdb11ada4d7902048f2354ba5bf9e10
		fondo.Pause()
		sIntro.Play()
	}

}
func sonidoGame() {
	sIntro.Pause()
	sIntro.Rewind()
	deadSound2.Rewind()
	fondo.Play()
}

func sonidoGameover() {
	fondo.Pause()
	fondo.Rewind()
	deadSound.Pause()
	deadSound2.Play()
}

func sonidoVidas() {
	deadSound.Play()
	deadSound.Rewind()
}

func fadeIn() {
	if vol == 0 {
		fade = true
	} else if vol > .99 {
		fade = false
	}
	if fade {
		vol += .01
	}
}
