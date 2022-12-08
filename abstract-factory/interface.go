package abstract_factory

// 抽象工厂
// 以SPA按摩为例
// 杭州和上海准备开两家SPA会所，杭州店和上海店属于两个产品族
// 精油开背和足浴属于同一个服务等级
// 所以在抽象工厂模式中，我们要添加两个工厂（每个工厂对应一家SPA会所），每个工厂实现两个商品服务的创建方法

// FactoryInterface
// 工厂接口
type FactoryInterface interface {
	EssentialOilOpenBack() CommodityServiceInterface // 精油开背服务
	FootBash() CommodityServiceInterface             // 足浴服务
}

// CommodityService
// 商品服务接口
type CommodityServiceInterface interface {
	Intro()
}
