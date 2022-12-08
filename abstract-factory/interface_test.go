package abstract_factory

import "testing"

// 杭州SPA精油开背
func TestHZFactory_EssentialOilOpenBack(t *testing.T) {
	var f FactoryInterface
	f = new(HZFactory)
	f.EssentialOilOpenBack().Intro()
}

// 杭州SPA足浴
func TestHZFactory_FootBash(t *testing.T) {
	var f FactoryInterface
	f = new(HZFactory)
	f.FootBash().Intro()
}

// 上海SPA精油开背
func TestSHFactory_EssentialOilOpenBack(t *testing.T) {
	var f FactoryInterface
	f = new(SHFactory)
	f.EssentialOilOpenBack().Intro()
}

// 上海SPA足浴
func TestSHFactory_FootBash(t *testing.T) {
	var f FactoryInterface
	f = new(SHFactory)
	f.FootBash().Intro()
}
