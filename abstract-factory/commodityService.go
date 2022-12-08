package abstract_factory

import "fmt"

type HZSpaOpenBack struct{}

func (hz *HZSpaOpenBack) Intro() {
	fmt.Println("杭州SPA精油开背")
}

type HZSpaFootBash struct{}

func (hz *HZSpaFootBash) Intro() {
	fmt.Println("杭州SPA足浴按摩")
}

type SHSpaOpenBack struct{}

func (sh *SHSpaOpenBack) Intro() {
	fmt.Println("上海SPA精油开背")
}

type SHSpaFootBash struct{}

func (sh *SHSpaFootBash) Intro() {
	fmt.Println("上海SPA足浴按摩")
}
