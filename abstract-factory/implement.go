package abstract_factory

// 杭州SPA工厂
type HZFactory struct{}

// 杭州SPA精油开背
func (h *HZFactory) EssentialOilOpenBack() CommodityServiceInterface {
	return &HZSpaOpenBack{}
}

// 杭州SPA足浴按摩
func (h *HZFactory) FootBash() CommodityServiceInterface {
	return &HZSpaFootBash{}
}

// 上海SPA工厂
type SHFactory struct{}

// 上海SPA精油开背
func (s *SHFactory) EssentialOilOpenBack() CommodityServiceInterface {
	return &SHSpaOpenBack{}
}

// 上海SPA足浴按摩
func (s *SHFactory) FootBash() CommodityServiceInterface {
	return &SHSpaFootBash{}
}
