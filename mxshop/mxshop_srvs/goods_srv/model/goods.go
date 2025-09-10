package model

// string默认会映射成sql中的text或者long text，处理不方便，我们喜欢varchar或者char
// 类型，字段是否能为null，尽量设置不为null，可以设置default
// 尽量使用int32，减少类型转换
type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(20);not null"`
	ParentCategoryID int32
	ParentCategory   *Category
	Level            int32 `gorm:"type:int;not null;default:1"`
	IsTab            bool  `gorm:"not null;default:false"`
}

type Brand struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);not null;default:''"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brand   Brand
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null;default:''"`
	url   string `gorm:"type:varchar(200);not null;default:''"`
	Index int32  `gorm:"type:int;not null;default:1"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandID    int32 `gorm:"type:int;not null"`
	Brand      Brand

	OnSale   bool `gorm:"type:bool;not null;default:false"` //是否上货架
	ShipFree bool `gorm:"type:bool;not null;default:false"` //是否免运费
	IsNew    bool `gorm:"type:bool;not null;default:false"`
	IsHot    bool `gorm:"type:bool;not null;default:false"`

	Name            string   `gorm:"type:varchar(50);not null"`
	GoodSn          string   `gorm:"type:varchar(50);not null"`   //商品单号
	ClickNum        int32    `gorm:"type:int;not null;default:0"` //点击量
	SoldNum         int32    `gorm:"type:int;not null;default:0"` //售出量
	FavNum          int32    `gorm:"type:int;not null;default:0"` //收藏数量
	MarketPrice     float32  `gorm:"not null"`                    //市场价
	ShopPrice       float32  `gorm:"not null"`                    //店铺价
	GoodsBrief      string   `gorm:"type:varchar(100);not null;default:''"`
	Images          GormList `gorm:"type:varchar(1000);not null"` //商品展示图
	DescImages      GormList `gorm:"type:varchar(1000);not null"` //描述信息中的商品图
	GoodsFrontImage string   `gorm:"type:varchar(200);not null"`  //封面图片
}
