package opensea

type Asset struct {
	AnimationOriginalURL    string             `json:"animation_original_url"`
	AnimationURL            string             `json:"animation_url"`
	AssetContract           AssetAssetContract `json:"asset_contract"`
	BackgroundColor         string             `json:"background_color"`
	Collection              AssetCollection    `json:"collection"`
	Creator                 AssetCreator       `json:"creator"`
	Decimals                string             `json:"decimals"`
	Description             string             `json:"description"`
	ExternalLink            string             `json:"external_link"`
	ID                      int                `json:"id"`
	ImageOriginalURL        string             `json:"image_original_url"`
	ImagePreviewURL         string             `json:"image_preview_url"`
	ImageThumbnailURL       string             `json:"image_thumbnail_url"`
	ImageURL                string             `json:"image_url"`
	IsPresale               bool               `json:"is_presale"`
	LastSale                string             `json:"last_sale"`
	ListingDate             string             `json:"listing_date"`
	Name                    string             `json:"name"`
	NumSales                int                `json:"num_sales"`
	Owner                   AssetOwner         `json:"owner"`
	Permalink               string             `json:"permalink"`
	SellOrders              string             `json:"sell_orders"`
	TokenID                 string             `json:"token_id"`
	TokenMetadata           string             `json:"token_metadata"`
	TopBid                  string             `json:"top_bid"`
	Traits                  []string           `json:"traits"`
	TransferFee             string             `json:"transfer_fee"`
	TransferFeePaymentToken string             `json:"transfer_fee_payment_token"`
}

type AssetAssetContract struct {
	Address                     string `json:"address"`
	AssetContractType           string `json:"asset_contract_type"`
	BuyerFeeBasisPoints         int    `json:"buyer_fee_basis_points"`
	CreatedDate                 string `json:"created_date"`
	DefaultToFiat               bool   `json:"default_to_fiat"`
	Description                 string `json:"description"`
	DevBuyerFeeBasisPoints      int    `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int    `json:"dev_seller_fee_basis_points"`
	ExternalLink                string `json:"external_link"`
	ImageURL                    string `json:"image_url"`
	Name                        string `json:"name"`
	NftVersion                  string `json:"nft_version"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int    `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int    `json:"opensea_seller_fee_basis_points"`
	OpenseaVersion              string `json:"opensea_version"`
	Owner                       int    `json:"owner"`
	PayoutAddress               string `json:"payout_address"`
	SchemaName                  string `json:"schema_name"`
	SellerFeeBasisPoints        int    `json:"seller_fee_basis_points"`
	Symbol                      string `json:"symbol"`
	TotalSupply                 string `json:"total_supply"`
}

type AssetCollection struct {
	BannerImageURL              string           `json:"banner_image_url"`
	ChatURL                     string           `json:"chat_url"`
	CreatedDate                 string           `json:"created_date"`
	DefaultToFiat               bool             `json:"default_to_fiat"`
	Description                 string           `json:"description"`
	DevBuyerFeeBasisPoints      string           `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string           `json:"dev_seller_fee_basis_points"`
	DiscordURL                  string           `json:"discord_url"`
	DisplayData                 AssetDisplayData `json:"display_data"`
	ExternalURL                 string           `json:"external_url"`
	Featured                    bool             `json:"featured"`
	FeaturedImageURL            string           `json:"featured_image_url"`
	Hidden                      bool             `json:"hidden"`
	ImageURL                    string           `json:"image_url"`
	InstagramUsername           string           `json:"instagram_username"`
	IsSubjectToWhitelist        bool             `json:"is_subject_to_whitelist"`
	LargeImageURL               string           `json:"large_image_url"`
	MediumUsername              string           `json:"medium_username"`
	Name                        string           `json:"name"`
	OnlyProxiedTransfers        bool             `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string           `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string           `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string           `json:"payout_address"`
	RequireEmail                bool             `json:"require_email"`
	SafelistRequestStatus       string           `json:"safelist_request_status"`
	ShortDescription            string           `json:"short_description"`
	Slug                        string           `json:"slug"`
	TelegramURL                 string           `json:"telegram_url"`
	TwitterUsername             string           `json:"twitter_username"`
	WikiURL                     string           `json:"wiki_url"`
}

type AssetOwner struct {
	Address       string    `json:"address"`
	Config        string    `json:"config"`
	ProfileImgURL string    `json:"profile_img_url"`
	User          AssetUser `json:"user"`
}

type AssetCreator struct {
	Address       string    `json:"address"`
	Config        string    `json:"config"`
	ProfileImgURL string    `json:"profile_img_url"`
	User          AssetUser `json:"user"`
}

type AssetUser struct {
	Username string `json:"username"`
}

type AssetDisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}
