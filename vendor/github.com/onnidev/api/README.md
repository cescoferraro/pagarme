# github.com/onnidev/api

A Golang API to support ONNi

## Routes

<details>
<summary>`/antitheft`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/antitheft**
	- _GET_
		- [AttachBanCollection](/middlewares/ban.go#L16)
		- [AttachAntiTheftCollection](/middlewares/anti.go#L44)
		- [List](/antitheft/list.go#L16)

</details>
<details>
<summary>`/antitheft/{mode}/{antitheftId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/antitheft/{mode}/{antitheftId}**
	- _POST_
		- [AttachBanCollection](/middlewares/ban.go#L16)
		- [AttachAntiTheftCollection](/middlewares/anti.go#L44)
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [GetUserClubFromToken](/middlewares/userClub.go#L198)
		- [Activate](/antitheft/activate.go#L20)

</details>
<details>
<summary>`/appclub/v1/userclub/login`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/appclub/v1/userclub/login**
	- _POST_
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [AttachTokenCollection](/middlewares/tokens.go#L21)
		- [ReadSoftLoginRequestFromBody](/middlewares/userClub.go#L76)
		- [Login](/appclub/login.go#L20)

</details>
<details>
<summary>`/appclub/v1/voucher/read`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/appclub/v1/voucher/read**
	- _POST_
		- [ReadVoucherSoftReadReqFromBody](/middlewares/vouchers.go#L122)
		- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
		- [AttachPartiesCollection](/middlewares/party.go#L21)
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [AttachTokenCollection](/middlewares/tokens.go#L21)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [HeaderScan](/middlewares/tokens.go#L49)
		- [CheckToken](/middlewares/tokens.go#L81)
		- [Read](/appclub/read.go#L15)

</details>
<details>
<summary>`/appclub/v1/voucher/validate`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/appclub/v1/voucher/validate**
	- _POST_
		- [ReadVoucherSoftValidateReqFromBody](/middlewares/vouchers.go#L147)
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [AttachTokenCollection](/middlewares/tokens.go#L21)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [HeaderScan](/middlewares/tokens.go#L49)
		- [CheckToken](/middlewares/tokens.go#L81)
		- [Validate](/appclub/validate.go#L17)

</details>
<details>
<summary>`/appclub/v1/vouchers/read/club/{clubId}/userClub/{userClubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/appclub/v1/vouchers/read/club/{clubId}/userClub/{userClubId}**
	- _GET_
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [AttachTokenCollection](/middlewares/tokens.go#L21)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [CheckToken](/middlewares/tokens.go#L81)
		- [HeaderScan](/middlewares/tokens.go#L49)
		- [History](/appclub/history.go#L16)

</details>
<details>
<summary>`/auth`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/auth**
	- _POST_
		- [ReadCreateJWTRefreshRequestFromBody](/middlewares/auth.go#L41)
		- [AttachCustomerCollection](/middlewares/customer.go#L21)
		- [Auth](/auth/customer_reauth.go#L16)

</details>
<details>
<summary>`/auth/customer/login`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/auth/customer/login**
	- _POST_
		- [ReadCreateJWTRefreshRequestFromBody](/middlewares/auth.go#L41)
		- [AttachCustomerCollection](/middlewares/customer.go#L21)
		- [CustomerLogin](/auth/customer_login.go#L17)

</details>
<details>
<summary>`/auth/leitor`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/auth/leitor**
	- _POST_
		- [ReadCreateJWTRefreshRequestFromBody](/middlewares/auth.go#L41)
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [AttachClubsCollection](/middlewares/clubs.go#L23)
		- [LeitorReAuth](/auth/leitor_reauth.go#L17)

</details>
<details>
<summary>`/auth/sms/again/{phone}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/auth/sms/again/{phone}**
	- _GET_
		- [Sms](/auth/sms.go#L16)

</details>
<details>
<summary>`/auth/sms/totalvoice/{phone}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/auth/sms/totalvoice/{phone}**
	- _GET_
		- [TotalVoice](/auth/sms.go#L44)

</details>
<details>
<summary>`/auth/sms/twilio/{phone}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/auth/sms/twilio/{phone}**
	- _GET_
		- [Twilio](/auth/sms.go#L30)

</details>
<details>
<summary>`/auth/sms/{phone}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/auth/sms/{phone}**
	- _GET_
		- [Sms](/auth/sms.go#L16)

</details>
<details>
<summary>`/auth/userClub`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/auth/userClub**
	- _POST_
		- [ReadCreatePagarMeReAuthRequestFromBody](/middlewares/auth.go#L17)
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [AttachClubsCollection](/middlewares/clubs.go#L23)
		- [UserClubReAuth](/auth/userClub_reauth.go#L16)

</details>
<details>
<summary>`/ban/{customerId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/ban/{customerId}**
	- _POST_
		- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
		- [AttachCustomerCollection](/middlewares/customer.go#L21)
		- [AttachPushRegistryCollection](/middlewares/pushRegistry.go#L20)
		- [AttachBanCollection](/middlewares/ban.go#L16)
		- [AttachInvoicesCollection](/middlewares/invoice.go#L16)
		- [AttachCardsCollection](/middlewares/cards.go#L92)
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [GetUserClubFromToken](/middlewares/userClub.go#L198)
		- [BanEndpoint](/bans/ban.go#L16)

</details>
<details>
<summary>`/banner`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/banner**
	- **/**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachBannerCollection](/middlewares/banner.go#L20)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ReadCreateBannerRequestFromBody](/middlewares/banner.go#L75)
			- [PersistToDB](/banner/create.go#L19)
		- _GET_
			- [AttachBannerCollection](/middlewares/banner.go#L20)
			- [PublishedBanners](/banner/list.go#L11)

</details>
<details>
<summary>`/banner/all`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/banner**
	- **/all**
		- _GET_
			- [AttachBannerCollection](/middlewares/banner.go#L20)
			- [AllBanners](/banner/list.go#L23)

</details>
<details>
<summary>`/banner/{bannerID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/banner**
	- **/{bannerID}**
		- **/**
			- _PATCH_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachBannerCollection](/middlewares/banner.go#L20)
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
				- [ReadPatchBannerRequestFromBody](/middlewares/banner.go#L51)
				- [GetUserClubFromToken](/middlewares/userClub.go#L198)
				- [Patch](/banner/patch.go#L19)
			- _GET_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachBannerCollection](/middlewares/banner.go#L20)
				- [Read](/banner/read.go#L16)
			- _DELETE_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachBannerCollection](/middlewares/banner.go#L20)
				- [Delete](/banner/delete.go#L15)

</details>
<details>
<summary>`/buy`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/buy**
	- _POST_
		- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
		- [AttachCustomerCollection](/middlewares/customer.go#L21)
		- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
		- [AttachCardsCollection](/middlewares/cards.go#L92)
		- [AttachBanCollection](/middlewares/ban.go#L16)
		- [AttachPushRegistryCollection](/middlewares/pushRegistry.go#L20)
		- [AttachPartiesCollection](/middlewares/party.go#L21)
		- [AttachClubsCollection](/middlewares/clubs.go#L23)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [AttachInvoicesCollection](/middlewares/invoice.go#L16)
		- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
		- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
		- [GetCustomerFromToken](/middlewares/customer.go#L109)
		- [ReadCreateBuyPostFromBody](/middlewares/buy.go#L18)
		- [Business](/buy/buy.go#L17)

</details>
<details>
<summary>`/card`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/card**
	- **/**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachTokenCollection](/middlewares/tokens.go#L21)
			- [AttachCardsCollection](/middlewares/cards.go#L92)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [ListCards](/card/list.go#L38)
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachCardsCollection](/middlewares/cards.go#L92)
			- [AttachTokenCollection](/middlewares/tokens.go#L21)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [ReadCreateCardRequestFromBody](/middlewares/cards.go#L23)
			- [PagarMeCreateCardHash](/middlewares/cards.go#L123)
			- [PersistToDB](/card/persist.go#L44)

</details>
<details>
<summary>`/card/{id}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/card**
	- **/{id}**
		- **/**
			- _GET_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachCardsCollection](/middlewares/cards.go#L92)
				- [Read](/card/read.go#L40)
			- _DELETE_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachCardsCollection](/middlewares/cards.go#L92)
				- [AttachCustomerCollection](/middlewares/customer.go#L21)
				- [GetCustomerFromToken](/middlewares/customer.go#L109)
				- [Delete](/card/delete.go#L39)
			- _PATCH_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachCardsCollection](/middlewares/cards.go#L92)
				- [ReadUpdateCardRequestFromBody](/middlewares/cards.go#L68)
				- [AttachCustomerCollection](/middlewares/customer.go#L21)
				- [GetCustomerFromToken](/middlewares/customer.go#L109)
				- [Patch](/card/patch.go#L41)

</details>
<details>
<summary>`/cart/{partyID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/cart**
	- **/{partyID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [Cart](/cart/cart.go#L15)

</details>
<details>
<summary>`/club`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/club**
	- **/**
		- _POST_
			- [ReadClubPostRequestFromBody](/middlewares/clubs.go#L117)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
			- [PersistToDB](/club/persist.go#L22)
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ListClubs](/club/list.go#L13)

</details>
<details>
<summary>`/club/dashboard/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/club**
	- **/dashboard/{clubId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [Dashboard](/club/dashboard.go#L36)

</details>
<details>
<summary>`/club/image/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/club**
	- **/image/{clubId}**
		- _GET_
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [Image](/club/image.go#L14)

</details>
<details>
<summary>`/club/next/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/club**
	- **/next/{clubId}**
		- _GET_
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
			- [NextRead](/club/read.go#L13)

</details>
<details>
<summary>`/club/soft`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/club**
	- **/soft**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ComboSoft](/club/combo_soft.go#L15)

</details>
<details>
<summary>`/club/soft/image/background/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/club**
	- **/soft/image/background/{clubId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [BCKUpdateImage](/club/update_bck_image.go#L21)

</details>
<details>
<summary>`/club/soft/image/main/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/club**
	- **/soft/image/main/{clubId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [UpdateImage](/club/update_image.go#L21)

</details>
<details>
<summary>`/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/club**
	- **/{clubId}**
		- _PATCH_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [ReadClubPatchFromBody](/middlewares/clubs.go#L69)
			- [Patch](/club/patch.go#L19)
		- _GET_
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AppRead](/club/app_read.go#L13)

</details>
<details>
<summary>`/club/{mode}/{clubID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/club**
	- **/{mode}/{clubID}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Activate](/club/activate.go#L21)

</details>
<details>
<summary>`/clubLead`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/clubLead**
	- [AttachClubLeadCollection](/middlewares/clubLead.go#L20)
	- **/**
		- _POST_
			- [ReadCreateClubLeadRequestFromBody](/middlewares/clubLead.go#L48)
			- [CreateEndpoint](/clubLead/create.go#L21)
		- _GET_
			- [ListEndpoint](/clubLead/list.go#L14)

</details>
<details>
<summary>`/clubLead/done/{clubLeadID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/clubLead**
	- [AttachClubLeadCollection](/middlewares/clubLead.go#L20)
	- **/done/{clubLeadID}**
		- _POST_
			- [ReadPatchClubLeadRequestFromBody](/middlewares/clubLead.go#L72)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
			- [Done](/clubLead/done.go#L16)

</details>
<details>
<summary>`/clubLead/{clubLeadID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/clubLead**
	- [AttachClubLeadCollection](/middlewares/clubLead.go#L20)
	- **/{clubLeadID}**
		- _GET_
			- [ReadEndpoint](/clubLead/read.go#L14)
		- _PATCH_
			- [ReadPatchClubLeadRequestFromBody](/middlewares/clubLead.go#L72)
			- [Patch](/clubLead/patch.go#L18)

</details>
<details>
<summary>`/clublead`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/clublead**
	- [AttachClubLeadCollection](/middlewares/clubLead.go#L20)
	- **/**
		- _POST_
			- [ReadCreateClubLeadRequestFromBody](/middlewares/clubLead.go#L48)
			- [CreateEndpoint](/clubLead/create.go#L21)
		- _GET_
			- [ListEndpoint](/clubLead/list.go#L14)

</details>
<details>
<summary>`/clublead/done/{clubLeadID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/clublead**
	- [AttachClubLeadCollection](/middlewares/clubLead.go#L20)
	- **/done/{clubLeadID}**
		- _POST_
			- [ReadPatchClubLeadRequestFromBody](/middlewares/clubLead.go#L72)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
			- [Done](/clubLead/done.go#L16)

</details>
<details>
<summary>`/clublead/{clubLeadID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/clublead**
	- [AttachClubLeadCollection](/middlewares/clubLead.go#L20)
	- **/{clubLeadID}**
		- _GET_
			- [ReadEndpoint](/clubLead/read.go#L14)
		- _PATCH_
			- [ReadPatchClubLeadRequestFromBody](/middlewares/clubLead.go#L72)
			- [Patch](/clubLead/patch.go#L18)

</details>
<details>
<summary>`/customer`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/**
		- _PATCH_
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [ReadCustomerPatchFromBody](/middlewares/customer.go#L137)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [Patch](/customer/patch.go#L22)
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [ListCustomers](/customer/list.go#L12)
		- _POST_
			- [ReadCustomerPostFromBody](/middlewares/customer.go#L195)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [CreateEndpoint](/customer/create.go#L21)

</details>
<details>
<summary>`/customer/check`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/check**
		- _POST_
			- [ReadCustomerCheckFromBody](/middlewares/customer.go#L171)
			- [Check](/customer/check.go#L15)

</details>
<details>
<summary>`/customer/favorite`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/favorite**
		- _PATCH_
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [ReadCustomerPatchFromBody](/middlewares/customer.go#L137)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [PatchFavorite](/customer/patch_favorite.go#L18)

</details>
<details>
<summary>`/customer/fixup`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/fixup**
		- _PATCH_
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [ReadCustomerPatchFromBody](/middlewares/customer.go#L137)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [FixUpPatch](/customer/patch.go#L32)

</details>
<details>
<summary>`/customer/image/{customerId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/image/{customerId}**
		- _GET_
			- [Photo](/customer/image.go#L16)

</details>
<details>
<summary>`/customer/intel/{customerId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/intel/{customerId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachBanCollection](/middlewares/ban.go#L16)
			- [AttachInvoicesCollection](/middlewares/invoice.go#L16)
			- [AttachCardsCollection](/middlewares/cards.go#L92)
			- [ReadFullCustomers](/customer/read_full_customer.go#L16)

</details>
<details>
<summary>`/customer/login`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/login**
		- _POST_
			- [ReadLoginRequestFromBody](/middlewares/userClub.go#L101)
			- [Login](/customer/login.go#L14)

</details>
<details>
<summary>`/customer/login/facebook`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/login/facebook**
		- _POST_
			- [ReadFacebookLoginRequest](/middlewares/facebook.go#L17)
			- [FacebookLogin](/customer/facebook_login.go#L14)

</details>
<details>
<summary>`/customer/newcomer`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/newcomer**
		- _POST_
			- [ReadCustomerPostFromBody](/middlewares/customer.go#L195)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [NewComerEndpoint](/customer/newcomer.go#L21)

</details>
<details>
<summary>`/customer/password/reset`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/password**
		- **/reset**
			- _POST_
				- [ReadResetFromBody](/middlewares/customer.go#L217)
				- [Reset](/customer/password_reset.go#L23)

</details>
<details>
<summary>`/customer/password/update`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/password**
		- **/update**
			- _POST_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [ReadChangePasswordRequestFromBody](/middlewares/userClub.go#L52)
				- [GetCustomerFromToken](/middlewares/customer.go#L109)
				- [UpdatePassword](/customer/password_update.go#L18)

</details>
<details>
<summary>`/customer/profile/facebook`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/profile/facebook**
		- _POST_
			- [ReadFacebookSignUpRequest](/middlewares/facebook.go#L41)
			- [FacebookProfile](/customer/facebook_profile.go#L13)

</details>
<details>
<summary>`/customer/query`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/query**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [ReadCustomerQueryFromBody](/middlewares/customer.go#L85)
			- [Query](/customer/query.go#L15)

</details>
<details>
<summary>`/customer/redirect`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/redirect**
		- _GET_
			- [Redirect](/customer/redirect.go#L16)

</details>
<details>
<summary>`/customer/signup/facebook`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/signup/facebook**
		- _POST_
			- [ReadFacebookSignUpRequest](/middlewares/facebook.go#L41)
			- [FacebookSignUp](/customer/facebook_signup.go#L19)

</details>
<details>
<summary>`/customer/trust/{mode}/{customerId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/trust/{mode}/{customerId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Activate](/customer/activate.go#L21)

</details>
<details>
<summary>`/customer/weblogin/{location}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/weblogin/{location}**
		- _GET_
			- [FB](/customer/web_login.go#L18)

</details>
<details>
<summary>`/customer/{customerId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/customer**
	- [AttachCustomerCollection](/middlewares/customer.go#L21)
	- **/{customerId}**
		- _GET_
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [Read](/customer/read.go#L15)

</details>
<details>
<summary>`/deeplink`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/deeplink**
	- **/**
		- _GET_
			- [MainDeeplink](/deeplink/main.go#L12)

</details>
<details>
<summary>`/deeplink/club/{clubID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/deeplink**
	- **/club/{clubID}**
		- _GET_
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [CLubDeepLInk](/deeplink/club.go#L17)

</details>
<details>
<summary>`/deeplink/party/{partyID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/deeplink**
	- **/party/{partyID}**
		- _GET_
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [PartyDeeplink](/deeplink/party.go#L17)

</details>
<details>
<summary>`/file`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/file**
	- _PUT_
		- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
		- [ReadFileFromBody](/middlewares/gridfs.go#L49)
		- [Add](/file/add.go#L23)

</details>
<details>
<summary>`/file/s3/{id}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/file/s3/{id}**
	- _GET_
		- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
		- [Aws](/file/aws.go#L22)

</details>
<details>
<summary>`/file/{id}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/file/{id}**
	- _GET_
		- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
		- [Get](/file/get.go#L16)

</details>
<details>
<summary>`/file/{min}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/file/{min}**
	- _PUT_
		- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
		- [ReadFileFromBody](/middlewares/gridfs.go#L49)
		- [AddMin](/file/add_min.go#L26)

</details>
<details>
<summary>`/invitedCustomer`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/invitedCustomer**
	- **/**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachTokenCollection](/middlewares/tokens.go#L21)
			- [ListCards](/invitedCustomer/list.go#L38)

</details>
<details>
<summary>`/invitedCustomer/fb/{customerId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/invitedCustomer**
	- **/fb/{customerId}**
		- _GET_
			- [FB](/invitedCustomer/fb.go#L18)

</details>
<details>
<summary>`/invitedCustomer/link/{inviteId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/invitedCustomer**
	- **/link/{inviteId}**
		- _POST_
			- [ReadInvitedLinkCustomerPostRequestFromBody](/middlewares/invitedCustomer.go#L51)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [Link](/invitedCustomer/patch.go#L21)

</details>
<details>
<summary>`/invitedCustomer/return`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/invitedCustomer**
	- **/return**
		- _GET_
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [Return](/invitedCustomer/return.go#L18)

</details>
<details>
<summary>`/invitedcustomer`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/invitedcustomer**
	- **/**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachTokenCollection](/middlewares/tokens.go#L21)
			- [ListCards](/invitedCustomer/list.go#L38)

</details>
<details>
<summary>`/invitedcustomer/fb/{customerId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/invitedcustomer**
	- **/fb/{customerId}**
		- _GET_
			- [FB](/invitedCustomer/fb.go#L18)

</details>
<details>
<summary>`/invitedcustomer/link/{inviteId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/invitedcustomer**
	- **/link/{inviteId}**
		- _POST_
			- [ReadInvitedLinkCustomerPostRequestFromBody](/middlewares/invitedCustomer.go#L51)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [Link](/invitedCustomer/patch.go#L21)

</details>
<details>
<summary>`/invitedcustomer/return`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/invitedcustomer**
	- **/return**
		- _GET_
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [Return](/invitedCustomer/return.go#L18)

</details>
<details>
<summary>`/locate`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/locate**
	- **/**
		- _POST_
			- [ReadGeoLocalizationPostRequestFromBody](/middlewares/location.go#L17)
			- [Locate](/location/locate.go#L18)

</details>
<details>
<summary>`/menuProduct/clone/{menuId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuProduct**
	- **/clone/{menuId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [Clone](/menuProduct/clone.go#L17)

</details>
<details>
<summary>`/menuProduct/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuProduct**
	- **/club/{clubId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [ByClub](/menuProduct/list.go#L16)

</details>
<details>
<summary>`/menuProduct/{menuId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuProduct**
	- **/{menuId}**
		- _DELETE_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [Delete](/menuProduct/delete.go#L20)
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [Read](/menuProduct/read.go#L15)

</details>
<details>
<summary>`/menuTicket/clone/{menuId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuTicket**
	- **/clone/{menuId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [Clone](/menuTicket/clone.go#L17)

</details>
<details>
<summary>`/menuTicket/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuTicket**
	- **/club/{clubId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [ByClub](/menuTicket/list.go#L16)

</details>
<details>
<summary>`/menuTicket/{menuId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuTicket**
	- **/{menuId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [Read](/menuTicket/read.go#L16)
		- _DELETE_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [Delete](/menuTicket/delete.go#L20)

</details>
<details>
<summary>`/menuproduct/clone/{menuId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuproduct**
	- **/clone/{menuId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [Clone](/menuProduct/clone.go#L17)

</details>
<details>
<summary>`/menuproduct/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuproduct**
	- **/club/{clubId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [ByClub](/menuProduct/list.go#L16)

</details>
<details>
<summary>`/menuproduct/{menuId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuproduct**
	- **/{menuId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [Read](/menuProduct/read.go#L15)
		- _DELETE_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [Delete](/menuProduct/delete.go#L20)

</details>
<details>
<summary>`/menuticket/clone/{menuId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuticket**
	- **/clone/{menuId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [Clone](/menuTicket/clone.go#L17)

</details>
<details>
<summary>`/menuticket/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuticket**
	- **/club/{clubId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [ByClub](/menuTicket/list.go#L16)

</details>
<details>
<summary>`/menuticket/{menuId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/menuticket**
	- **/{menuId}**
		- _DELETE_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [Delete](/menuTicket/delete.go#L20)
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [Read](/menuTicket/read.go#L16)

</details>
<details>
<summary>`/musicStyles`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/musicStyles**
	- **/**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachMusicStylesCollection](/middlewares/musicStyles.go#L16)
			- [List](/musicStyles/list.go#L14)

</details>
<details>
<summary>`/musicstyles`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/musicstyles**
	- **/**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachMusicStylesCollection](/middlewares/musicStyles.go#L16)
			- [List](/musicStyles/list.go#L14)

</details>
<details>
<summary>`/notification`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/notification**
	- **/**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachNotificationCollection](/middlewares/notification.go#L20)
			- [ListNotifications](/notification/list.go#L20)
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachBannerCollection](/middlewares/banner.go#L20)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachNotificationCollection](/middlewares/notification.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ReadCreateNotificationRequestFromBody](/middlewares/notification.go#L75)
			- [PersistToDB](/notification/persist.go#L19)

</details>
<details>
<summary>`/notification/customer/{customerID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/notification**
	- **/customer/{customerID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachNotificationCollection](/middlewares/notification.go#L20)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [ByCustomer](/notification/byCustomer.go#L19)

</details>
<details>
<summary>`/notification/publish/{notificationID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/notification**
	- **/publish/{notificationID}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachNotificationCollection](/middlewares/notification.go#L20)
			- [AttachPushRegistryCollection](/middlewares/pushRegistry.go#L20)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Publish](/notification/publish.go#L13)

</details>
<details>
<summary>`/notification/publish/{notificationID}/{partyProductID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/notification**
	- **/publish/{notificationID}/{partyProductID}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachNotificationCollection](/middlewares/notification.go#L20)
			- [AttachPushRegistryCollection](/middlewares/pushRegistry.go#L20)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Publish](/notification/publish.go#L13)

</details>
<details>
<summary>`/notification/{notificationID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/notification**
	- **/{notificationID}**
		- **/**
			- _DELETE_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachNotificationCollection](/middlewares/notification.go#L20)
				- [Delete](/notification/delete.go#L15)
			- _PATCH_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachClubsCollection](/middlewares/clubs.go#L23)
				- [AttachNotificationCollection](/middlewares/notification.go#L20)
				- [AttachPartiesCollection](/middlewares/party.go#L21)
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [GetUserClubFromToken](/middlewares/userClub.go#L198)
				- [ReadCreateNotificationPatchRequestFromBody](/middlewares/notification.go#L51)
				- [Patch](/notification/patch.go#L19)
			- _GET_
				- [AttachNotificationCollection](/middlewares/notification.go#L20)
				- [Read](/notification/read.go#L17)

</details>
<details>
<summary>`/party`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [ListParties](/party/list.go#L12)
		- _POST_
			- [ReadPartyListFilterRequestFromBody](/middlewares/party.go#L70)
			- [AppListPartiesFiltered](/party/app_filter-list.go#L12)
		- _PUT_
			- [ReadSoftPartyPostRequestRequestFromBody](/middlewares/party.go#L94)
			- [AttachMusicStylesCollection](/middlewares/musicStyles.go#L16)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [Create](/party/create.go#L20)

</details>
<details>
<summary>`/party/antitheft/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/antitheft/{partyId}**
		- _POST_
			- [AttachInvoicesCollection](/middlewares/invoice.go#L16)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachCardsCollection](/middlewares/cards.go#L92)
			- [ReadAntiTheftModelRequestFromBody](/middlewares/anti.go#L20)
			- [AntiTheft](/party/antitheft.go#L23)

</details>
<details>
<summary>`/party/club/site/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/club/site/{clubId}**
		- _GET_
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadClubPartiesSite](/party/read_club_parties_site_now.go#L16)

</details>
<details>
<summary>`/party/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/club/{clubId}**
		- _GET_
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [IsRequestFromApp](/interfaces/jwt.go#L16)
			- [ReadClubParties](/party/read_club_parties.go#L38)

</details>
<details>
<summary>`/party/image/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/image/{partyId}**
		- _GET_
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [Image](/party/image.go#L15)

</details>
<details>
<summary>`/party/next`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/next**
		- _POST_
			- [ReadPartyListFilterRequestFromBody](/middlewares/party.go#L70)
			- [ListPartiesFiltered](/party/filtered_list.go#L12)

</details>
<details>
<summary>`/party/soft/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/soft/club/{clubId}**
		- _GET_
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [SoftClubInfo](/party/soft_club_info.go#L16)

</details>
<details>
<summary>`/party/soft/image/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/soft/image/{partyId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [UpdateImage](/party/update_image.go#L21)

</details>
<details>
<summary>`/party/soft/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/soft/{partyId}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [ReadPartySoft](/party/read_soft.go#L15)

</details>
<details>
<summary>`/party/userClub`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/userClub**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [UserClubListParties](/party/userClub_list.go#L16)

</details>
<details>
<summary>`/party/{mode}/{partyID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/{mode}/{partyID}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Activate](/party/activate.go#L22)

</details>
<details>
<summary>`/party/{partyID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/{partyID}**
		- _PATCH_
			- [ReadSoftPartyPostRequestRequestFromBody](/middlewares/party.go#L94)
			- [AttachMusicStylesCollection](/middlewares/musicStyles.go#L16)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachClubMenuTicketCollection](/middlewares/menuTicket.go#L16)
			- [AttachClubMenuProductCollection](/middlewares/menuProduct.go#L16)
			- [Patch](/party/patch.go#L17)

</details>
<details>
<summary>`/party/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/party**
	- [AttachPartiesCollection](/middlewares/party.go#L21)
	- **/{partyId}**
		- _GET_
			- [ReadParty](/party/read.go#L13)

</details>
<details>
<summary>`/partyProduct/party/soft/combo/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProduct**
	- **/party/soft/combo/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [SoftCombo](/partyProduct/soft_combo.go#L14)

</details>
<details>
<summary>`/partyProduct/promotion/customer/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProduct**
	- **/promotion/customer/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [CustomersSummary](/promotion/customer.go#L17)

</details>
<details>
<summary>`/partyProduct/promotion/party/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProduct**
	- **/promotion/party/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ListPartyPromotions](/promotion/list_party_promotion.go#L17)

</details>
<details>
<summary>`/partyProduct/promotion/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProduct**
	- **/promotion/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadPromotion](/promotion/read_promotion.go#L19)

</details>
<details>
<summary>`/partyProduct/tickets/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProduct**
	- **/tickets/{partyId}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [TicketProducts](/partyProduct/tickets.go#L16)

</details>
<details>
<summary>`/partyProduct/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProduct**
	- **/{partyId}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [Products](/partyProduct/list.go#L16)

</details>
<details>
<summary>`/partyProducts/party/soft/combo/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProducts**
	- **/party/soft/combo/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [SoftCombo](/partyProduct/soft_combo.go#L14)

</details>
<details>
<summary>`/partyProducts/promotion/customer/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProducts**
	- **/promotion/customer/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [CustomersSummary](/promotion/customer.go#L17)

</details>
<details>
<summary>`/partyProducts/promotion/party/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProducts**
	- **/promotion/party/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ListPartyPromotions](/promotion/list_party_promotion.go#L17)

</details>
<details>
<summary>`/partyProducts/promotion/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProducts**
	- **/promotion/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadPromotion](/promotion/read_promotion.go#L19)

</details>
<details>
<summary>`/partyProducts/tickets/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProducts**
	- **/tickets/{partyId}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [TicketProducts](/partyProduct/tickets.go#L16)

</details>
<details>
<summary>`/partyProducts/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyProducts**
	- **/{partyId}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [Products](/partyProduct/list.go#L16)

</details>
<details>
<summary>`/partyproduct/party/soft/combo/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproduct**
	- **/party/soft/combo/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [SoftCombo](/partyProduct/soft_combo.go#L14)

</details>
<details>
<summary>`/partyproduct/promotion/customer/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproduct**
	- **/promotion/customer/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [CustomersSummary](/promotion/customer.go#L17)

</details>
<details>
<summary>`/partyproduct/promotion/party/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproduct**
	- **/promotion/party/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ListPartyPromotions](/promotion/list_party_promotion.go#L17)

</details>
<details>
<summary>`/partyproduct/promotion/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproduct**
	- **/promotion/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadPromotion](/promotion/read_promotion.go#L19)

</details>
<details>
<summary>`/partyproduct/tickets/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproduct**
	- **/tickets/{partyId}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [TicketProducts](/partyProduct/tickets.go#L16)

</details>
<details>
<summary>`/partyproduct/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproduct**
	- **/{partyId}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [Products](/partyProduct/list.go#L16)

</details>
<details>
<summary>`/partyproducts/party/soft/combo/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproducts**
	- **/party/soft/combo/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [SoftCombo](/partyProduct/soft_combo.go#L14)

</details>
<details>
<summary>`/partyproducts/promotion/customer/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproducts**
	- **/promotion/customer/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [CustomersSummary](/promotion/customer.go#L17)

</details>
<details>
<summary>`/partyproducts/promotion/party/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproducts**
	- **/promotion/party/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ListPartyPromotions](/promotion/list_party_promotion.go#L17)

</details>
<details>
<summary>`/partyproducts/promotion/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproducts**
	- **/promotion/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadPromotion](/promotion/read_promotion.go#L19)

</details>
<details>
<summary>`/partyproducts/tickets/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproducts**
	- **/tickets/{partyId}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [TicketProducts](/partyProduct/tickets.go#L16)

</details>
<details>
<summary>`/partyproducts/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/partyproducts**
	- **/{partyId}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [Products](/partyProduct/list.go#L16)

</details>
<details>
<summary>`/product`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/product**
	- **/**
		- _PUT_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [Add](/product/add.go#L16)
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [Products](/product/list.go#L15)
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [ReadProductPostRequestFromBody](/middlewares/product.go#L51)
			- [Post](/product/post.go#L16)

</details>
<details>
<summary>`/product/soft/image/{productId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/product**
	- **/soft/image/{productId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [UpdateImage](/product/update_image.go#L21)

</details>
<details>
<summary>`/product/soft/{productId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/product**
	- **/soft/{productId}**
		- _PATCH_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [ReadProductSoftPatchRequestFromBody](/middlewares/product.go#L99)
			- [SoftPatch](/product/soft_patch.go#L19)

</details>
<details>
<summary>`/product/{id}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/product**
	- **/{id}**
		- _DELETE_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [Delete](/product/delete.go#L14)

</details>
<details>
<summary>`/product/{productId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/product**
	- **/{productId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [Read](/product/read.go#L14)
		- _PATCH_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachProductsCollection](/middlewares/product.go#L20)
			- [AttachGridFSCollection](/middlewares/gridfs.go#L20)
			- [ReadProductPatchRequestFromBody](/middlewares/product.go#L75)
			- [Patch](/product/patch.go#L16)

</details>
<details>
<summary>`/promotion`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotion**
	- **/**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [ReadPromotionPostRequestRequestFromBody](/middlewares/ppromotion.go#L41)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [CreatePromotion](/promotion/create.go#L19)

</details>
<details>
<summary>`/promotion/customer/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotion**
	- **/customer/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [CustomersSummary](/promotion/customer.go#L17)

</details>
<details>
<summary>`/promotion/party/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotion**
	- **/party/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ListPartyPromotions](/promotion/list_party_promotion.go#L17)

</details>
<details>
<summary>`/promotion/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotion**
	- **/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadPromotion](/promotion/read_promotion.go#L19)
		- _PATCH_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [ReadPromotionPatchRequestRequestFromBody](/middlewares/ppromotion.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [PatchPromotion](/promotion/patch_promotion.go#L20)

</details>
<details>
<summary>`/promotionalCustomer`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotionalCustomer**
	- **/**
		- _GET_
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [List](/promotionalCustomer/list.go#L12)
		- _POST_
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ReadPromotionalCustomerQueryFromBody](/middlewares/promotionalCustomer.go#L50)
			- [Create](/promotionalCustomer/create.go#L25)

</details>
<details>
<summary>`/promotionalCustomer/invite`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotionalCustomer**
	- **/invite**
		- _POST_
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ReadPromotionalCustomerQueryFromBody](/middlewares/promotionalCustomer.go#L50)
			- [Invite](/promotionalCustomer/invite.go#L20)

</details>
<details>
<summary>`/promotionalcustomer`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotionalcustomer**
	- **/**
		- _GET_
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [List](/promotionalCustomer/list.go#L12)
		- _POST_
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ReadPromotionalCustomerQueryFromBody](/middlewares/promotionalCustomer.go#L50)
			- [Create](/promotionalCustomer/create.go#L25)

</details>
<details>
<summary>`/promotionalcustomer/invite`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotionalcustomer**
	- **/invite**
		- _POST_
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ReadPromotionalCustomerQueryFromBody](/middlewares/promotionalCustomer.go#L50)
			- [Invite](/promotionalCustomer/invite.go#L20)

</details>
<details>
<summary>`/promotions`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotions**
	- **/**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [ReadPromotionPostRequestRequestFromBody](/middlewares/ppromotion.go#L41)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [CreatePromotion](/promotion/create.go#L19)

</details>
<details>
<summary>`/promotions/customer/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotions**
	- **/customer/{promotionID}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [CustomersSummary](/promotion/customer.go#L17)

</details>
<details>
<summary>`/promotions/party/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotions**
	- **/party/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ListPartyPromotions](/promotion/list_party_promotion.go#L17)

</details>
<details>
<summary>`/promotions/{promotionID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/promotions**
	- **/{promotionID}**
		- _PATCH_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [ReadPromotionPatchRequestRequestFromBody](/middlewares/ppromotion.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [PatchPromotion](/promotion/patch_promotion.go#L20)
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadPromotion](/promotion/read_promotion.go#L19)

</details>
<details>
<summary>`/proxy`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/proxy**
	- _*_
		- [Proxy](/proxy/proxy.go#L16)

</details>
<details>
<summary>`/proxy/app/v4/party/products/buy`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/proxy/app/v4/party/products/buy**
	- _POST_
		- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
		- [AttachCustomerCollection](/middlewares/customer.go#L21)
		- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
		- [AttachCardsCollection](/middlewares/cards.go#L92)
		- [AttachBanCollection](/middlewares/ban.go#L16)
		- [AttachPushRegistryCollection](/middlewares/pushRegistry.go#L20)
		- [AttachPartiesCollection](/middlewares/party.go#L21)
		- [AttachClubsCollection](/middlewares/clubs.go#L23)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [AttachInvoicesCollection](/middlewares/invoice.go#L16)
		- [AttachCardsCollection](/middlewares/cards.go#L92)
		- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
		- [ReadCreateBuyPostFromBody](/middlewares/buy.go#L18)
		- [GetCustomerFromToken](/middlewares/customer.go#L109)
		- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
		- [AttachBanCollection](/middlewares/ban.go#L16)
		- [Business](/buy/buy.go#L17)

</details>
<details>
<summary>`/proxy/app/v4/shopping/cart/customer/{customerId}/party/{partyID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/proxy/app/v4/shopping/cart/customer/{customerId}/party/{partyID}**
	- _GET_
		- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
		- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
		- [AttachCustomerCollection](/middlewares/customer.go#L21)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [AttachClubsCollection](/middlewares/clubs.go#L23)
		- [AttachPartiesCollection](/middlewares/party.go#L21)
		- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
		- [GetCustomerFromToken](/middlewares/customer.go#L109)
		- [Cart](/cart/cart.go#L15)

</details>
<details>
<summary>`/pushRegistry`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/pushRegistry**
	- **/**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPushRegistryCollection](/middlewares/pushRegistry.go#L20)
			- [ListPushRegistry](/pushRegistry/list.go#L16)
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPushRegistryCollection](/middlewares/pushRegistry.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [ReadCreatePushRegistryRequestFromBody](/middlewares/pushRegistry.go#L51)
			- [Create](/pushRegistry/create.go#L19)

</details>
<details>
<summary>`/recipient`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/**
		- _GET_
			- [ListRecipients](/recipient/recipient_list.go#L12)
		- _POST_
			- [ReadRecipientPostRequestFromBody](/middlewares/recipient.go#L125)
			- [Create](/recipient/recipient_create.go#L17)

</details>
<details>
<summary>`/recipient/antecipation`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/antecipation**
		- **/**
			- _POST_
				- [ReadAntecipationPostRequestFromBody](/middlewares/recipient.go#L51)
				- [CreateAnteciapation](/recipient/antecipation_create.go#L15)

</details>
<details>
<summary>`/recipient/antecipation/cancel/{recipientID}/{bulkID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/antecipation**
		- **/cancel/{recipientID}/{bulkID}**
			- _POST_
				- [CancelAnteciapation](/recipient/antecipation_cancel.go#L13)

</details>
<details>
<summary>`/recipient/antecipation/confirm/{recipientID}/{bulkID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/antecipation**
		- **/confirm/{recipientID}/{bulkID}**
			- _POST_
				- [ConfirmAnteciapation](/recipient/antecipation_confirm.go#L13)

</details>
<details>
<summary>`/recipient/antecipation/limits`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/antecipation**
		- **/limits**
			- _POST_
				- [ReadFinanceQueryFromBody](/middlewares/transaction.go#L17)
				- [AnteciapationsLimit](/recipient/antecipation_limit.go#L17)

</details>
<details>
<summary>`/recipient/antecipation/{bulkID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/antecipation**
		- **/{bulkID}**
			- _PUT_
				- [ReadAntecipationPostRequestFromBody](/middlewares/recipient.go#L51)
				- [EditAnteciapation](/recipient/antecipation_edit.go#L17)

</details>
<details>
<summary>`/recipient/antecipation/{recipientID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/antecipation**
		- **/{recipientID}**
			- _GET_
				- [Anteciapations](/recipient/antecipations_list.go#L15)

</details>
<details>
<summary>`/recipient/antecipation/{recipientID}/{bulkID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/antecipation**
		- **/{recipientID}/{bulkID}**
			- _DELETE_
				- [DeleteAnteciapation](/recipient/antecipation_delete.go#L14)

</details>
<details>
<summary>`/recipient/balance`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/balance**
		- **/**
			- _POST_
				- [ReadFinanceQueryFromBody](/middlewares/transaction.go#L17)
				- [BalanceTransactions](/recipient/recipient_pagarme_balance.go#L12)

</details>
<details>
<summary>`/recipient/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/club**
		- **/{clubId}**
			- _GET_
				- [ListClubRecipients](/recipient/recipient_byclub.go#L13)

</details>
<details>
<summary>`/recipient/operations`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/operations**
		- [ReadFinanceQueryFromBody](/middlewares/transaction.go#L17)
		- **/**
			- _POST_
				- [Operations](/recipient/operations_list.go#L21)

</details>
<details>
<summary>`/recipient/operations/timeline`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/operations**
		- [ReadFinanceQueryFromBody](/middlewares/transaction.go#L17)
		- **/timeline**
			- _POST_
				- [DaysBalanceTransactions](/recipient/operations_timeline.go#L21)

</details>
<details>
<summary>`/recipient/operations/xlsx`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/operations**
		- [ReadFinanceQueryFromBody](/middlewares/transaction.go#L17)
		- **/xlsx**
			- _POST_
				- [Xlsx](/recipient/operations_xlsx.go#L17)

</details>
<details>
<summary>`/recipient/pagarme/tweaks/{recipientID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/pagarme**
		- **/tweaks/{recipientID}**
			- _PATCH_
				- [ReadRecipientTweaksPatchFromBody](/middlewares/recipient.go#L77)
				- [PagarMeRecipientTweaks](/recipient/recipient_pagarme_tweaks.go#L20)

</details>
<details>
<summary>`/recipient/pagarme/withdraw/{recipientID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/pagarme**
		- **/withdraw/{recipientID}**
			- _POST_
				- [ReadRecipientWithDrawRequestFromBody](/middlewares/recipient.go#L151)
				- [WithDraw](/recipient/recipient_pagarme_withdraw.go#L17)

</details>
<details>
<summary>`/recipient/pagarme/withdraws/{recipientID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/pagarme**
		- **/withdraws/{recipientID}**
			- _GET_
				- [WithDraws](/recipient/recipient_pagarme_withdraws.go#L16)

</details>
<details>
<summary>`/recipient/pagarme/{recipientID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/pagarme**
		- **/{recipientID}**
			- _GET_
				- [PagarMeRecipient](/recipient/recipient_pagarme.go#L17)

</details>
<details>
<summary>`/recipient/payables`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/payables**
		- [ReadFinanceQueryFromBody](/middlewares/transaction.go#L17)
		- **/**
			- _POST_
				- [Payables](/recipient/payables_list.go#L18)

</details>
<details>
<summary>`/recipient/payables/timeline`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/payables**
		- [ReadFinanceQueryFromBody](/middlewares/transaction.go#L17)
		- **/timeline**
			- _POST_
				- [PayablesTimeline](/recipient/payables_timeline.go#L21)

</details>
<details>
<summary>`/recipient/transactions`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/transactions**
		- [ReadFinanceQueryFromBody](/middlewares/transaction.go#L17)
		- **/**
			- _POST_
				- [ListTransactions](/recipient/transactions_list.go#L20)

</details>
<details>
<summary>`/recipient/{recipientID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/recipient**
	- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
	- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
	- **/{recipientID}**
		- _PATCH_
			- [ReadRecipientPatchFromBody](/middlewares/recipient.go#L101)
			- [Patch](/recipient/recipient_patch.go#L24)
		- _GET_
			- [Read](/recipient/recipient_read.go#L40)

</details>
<details>
<summary>`/refund/{voucherId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/refund/{voucherId}**
	- _POST_
		- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
		- [AttachCustomerCollection](/middlewares/customer.go#L21)
		- [AttachRecipientsCollection](/middlewares/recipient.go#L20)
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [AttachCardsCollection](/middlewares/cards.go#L92)
		- [AttachBanCollection](/middlewares/ban.go#L16)
		- [AttachAntiTheftCollection](/middlewares/anti.go#L44)
		- [AttachPushRegistryCollection](/middlewares/pushRegistry.go#L20)
		- [AttachPartiesCollection](/middlewares/party.go#L21)
		- [AttachClubsCollection](/middlewares/clubs.go#L23)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [AttachInvoicesCollection](/middlewares/invoice.go#L16)
		- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
		- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
		- [GetUserClubFromToken](/middlewares/userClub.go#L198)
		- [Refund](/buy/refund.go#L18)

</details>
<details>
<summary>`/report/download/{id}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/report/download/{id}**
	- _GET_
		- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [GetVouchers](/middlewares/vouchers.go#L54)
		- [AttachPartiesCollection](/middlewares/party.go#L21)
		- [GetPartyByID](/middlewares/party.go#L51)
		- [AttachClubsCollection](/middlewares/clubs.go#L23)
		- [GetClubFromPartyMiddleware](/middlewares/clubs.go#L50)
		- [GenerateExcel](/report/generate.go#L34)

</details>
<details>
<summary>`/report/mail/{id}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/report/mail/{id}**
	- _POST_
		- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [GetVouchers](/middlewares/vouchers.go#L54)
		- [AttachPartiesCollection](/middlewares/party.go#L21)
		- [GetPartyByID](/middlewares/party.go#L51)
		- [AttachClubsCollection](/middlewares/clubs.go#L23)
		- [GetClubFromPartyMiddleware](/middlewares/clubs.go#L50)
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [GetUserClubFromToken](/middlewares/userClub.go#L198)
		- [SendMail](/report/send_mail.go#L33)

</details>
<details>
<summary>`/report/party/{id}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/report/party/{id}**
	- _POST_
		- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
		- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
		- [GetVouchers](/middlewares/vouchers.go#L54)
		- [AttachPartiesCollection](/middlewares/party.go#L21)
		- [GetPartyByID](/middlewares/party.go#L51)
		- [AttachClubsCollection](/middlewares/clubs.go#L23)
		- [GetClubFromPartyMiddleware](/middlewares/clubs.go#L50)
		- [AttachUserClubCollection](/middlewares/userClub.go#L21)
		- [CloseParty](/report/closeParty_mail.go#L34)

</details>
<details>
<summary>`/site/cart/promo/{partyID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/site**
	- **/cart/promo/{partyID}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [PromoCart](/site/cart.go#L38)

</details>
<details>
<summary>`/site/cart/ticket/{partyID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/site**
	- **/cart/ticket/{partyID}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [TicketCart](/site/cart.go#L16)

</details>
<details>
<summary>`/site/cart/{partyID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/site**
	- **/cart/{partyID}**
		- _GET_
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachPromotionalCustomerCollection](/middlewares/promotionalCustomer.go#L20)
			- [TicketCart](/site/cart.go#L16)

</details>
<details>
<summary>`/site/customer/next`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/site**
	- **/customer/next**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [ByCustomer](/site/byCustomer.go#L14)

</details>
<details>
<summary>`/tickets`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/tickets**
	- **/**
		- _GET_
			- [BudDeeplink](/deeplink/bud.go#L12)

</details>
<details>
<summary>`/token`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/token**
	- _GET_
		- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
		- [AttachTokenCollection](/middlewares/tokens.go#L21)
		- [ListTokens](/token/list.go#L12)

</details>
<details>
<summary>`/userClub`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [ReadUserClubPostRequestFromBody](/middlewares/userClub.go#L174)
			- [Create](/userClub/create.go#L17)
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ListUsersClub](/userClub/list.go#L31)

</details>
<details>
<summary>`/userClub/android/vouchers`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/android/vouchers**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [SimpleVoucherReadByUserClub](/userClub/vouchers_read.go#L69)

</details>
<details>
<summary>`/userClub/check`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/check**
		- _POST_
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadCustomerCheckFromBody](/middlewares/customer.go#L171)
			- [Check](/userClub/check.go#L16)

</details>
<details>
<summary>`/userClub/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/club**
		- **/{clubId}**
			- _GET_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ByClub](/userClub/byParty.go#L32)

</details>
<details>
<summary>`/userClub/image/{token}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/image/{token}**
		- _GET_
			- [FBImage](/userClub/fbImage.go#L13)

</details>
<details>
<summary>`/userClub/login`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/login**
		- **/**
			- _POST_
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ReadLoginRequestFromBody](/middlewares/userClub.go#L101)
				- [AttachClubsCollection](/middlewares/clubs.go#L23)
				- [LoginUser](/userClub/login.go#L12)

</details>
<details>
<summary>`/userClub/login/leitor`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/login**
		- **/leitor**
			- **/**
				- _POST_
					- [AttachUserClubCollection](/middlewares/userClub.go#L21)
					- [ReadLoginRequestFromBody](/middlewares/userClub.go#L101)
					- [AttachClubsCollection](/middlewares/clubs.go#L23)
					- [AttachPartiesCollection](/middlewares/party.go#L21)
					- [LoginLeitor](/userClub/login_leitor.go#L15)

</details>
<details>
<summary>`/userClub/login/leitor/oauth2`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/login**
		- **/leitor**
			- **/oauth2**
				- _POST_
					- [AttachUserClubCollection](/middlewares/userClub.go#L21)
					- [AttachClubsCollection](/middlewares/clubs.go#L23)
					- [ReadOauth2LoginEmail](/middlewares/userClub.go#L126)
					- [Oauth2](/userClub/oauth2.go#L16)

</details>
<details>
<summary>`/userClub/login/soft`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/login**
		- **/soft**
			- _POST_
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ReadSoftLoginRequestFromBody](/middlewares/userClub.go#L76)
				- [AttachClubsCollection](/middlewares/clubs.go#L23)
				- [SoftLoginUser](/userClub/soft_login.go#L12)

</details>
<details>
<summary>`/userClub/onni`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/onni**
		- **/**
			- _GET_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ONNi](/userClub/onni.go#L31)

</details>
<details>
<summary>`/userClub/parties`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/parties**
		- **/**
			- _GET_
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
				- [AttachClubsCollection](/middlewares/clubs.go#L23)
				- [AttachPartiesCollection](/middlewares/party.go#L21)
				- [GetUserClubFromToken](/middlewares/userClub.go#L198)
				- [PartiesUserClub](/userClub/parties.go#L15)

</details>
<details>
<summary>`/userClub/password/reset`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/password**
		- **/reset**
			- _POST_
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ReadResetFromBody](/middlewares/customer.go#L217)
				- [Reset](/userClub/password_reset.go#L23)

</details>
<details>
<summary>`/userClub/promoter/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/promoter/{clubId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [SoftCombo](/userClub/soft_combo.go#L16)

</details>
<details>
<summary>`/userClub/vouchers`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/vouchers**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [VoucherReadByUserClub](/userClub/vouchers_read.go#L18)

</details>
<details>
<summary>`/userClub/{mode}/{userClubID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/{mode}/{userClubID}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Activate](/userClub/activate.go#L21)

</details>
<details>
<summary>`/userClub/{userClubID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userClub**
	- **/{userClubID}**
		- **/**
			- _PATCH_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [ReadUserClubPatchRequestFromBody](/middlewares/userClub.go#L150)
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [GetUserClubFromToken](/middlewares/userClub.go#L198)
				- [Patch](/userClub/patch.go#L20)

</details>
<details>
<summary>`/userclub`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [ReadUserClubPostRequestFromBody](/middlewares/userClub.go#L174)
			- [Create](/userClub/create.go#L17)
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ListUsersClub](/userClub/list.go#L31)

</details>
<details>
<summary>`/userclub/android/vouchers`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/android/vouchers**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [SimpleVoucherReadByUserClub](/userClub/vouchers_read.go#L69)

</details>
<details>
<summary>`/userclub/check`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/check**
		- _POST_
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadCustomerCheckFromBody](/middlewares/customer.go#L171)
			- [Check](/userClub/check.go#L16)

</details>
<details>
<summary>`/userclub/club/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/club**
		- **/{clubId}**
			- _GET_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ByClub](/userClub/byParty.go#L32)

</details>
<details>
<summary>`/userclub/image/{token}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/image/{token}**
		- _GET_
			- [FBImage](/userClub/fbImage.go#L13)

</details>
<details>
<summary>`/userclub/login`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/login**
		- **/**
			- _POST_
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ReadLoginRequestFromBody](/middlewares/userClub.go#L101)
				- [AttachClubsCollection](/middlewares/clubs.go#L23)
				- [LoginUser](/userClub/login.go#L12)

</details>
<details>
<summary>`/userclub/login/leitor`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/login**
		- **/leitor**
			- **/**
				- _POST_
					- [AttachUserClubCollection](/middlewares/userClub.go#L21)
					- [ReadLoginRequestFromBody](/middlewares/userClub.go#L101)
					- [AttachClubsCollection](/middlewares/clubs.go#L23)
					- [AttachPartiesCollection](/middlewares/party.go#L21)
					- [LoginLeitor](/userClub/login_leitor.go#L15)

</details>
<details>
<summary>`/userclub/login/leitor/oauth2`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/login**
		- **/leitor**
			- **/oauth2**
				- _POST_
					- [AttachUserClubCollection](/middlewares/userClub.go#L21)
					- [AttachClubsCollection](/middlewares/clubs.go#L23)
					- [ReadOauth2LoginEmail](/middlewares/userClub.go#L126)
					- [Oauth2](/userClub/oauth2.go#L16)

</details>
<details>
<summary>`/userclub/login/soft`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/login**
		- **/soft**
			- _POST_
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ReadSoftLoginRequestFromBody](/middlewares/userClub.go#L76)
				- [AttachClubsCollection](/middlewares/clubs.go#L23)
				- [SoftLoginUser](/userClub/soft_login.go#L12)

</details>
<details>
<summary>`/userclub/onni`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/onni**
		- **/**
			- _GET_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ONNi](/userClub/onni.go#L31)

</details>
<details>
<summary>`/userclub/parties`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/parties**
		- **/**
			- _GET_
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
				- [AttachClubsCollection](/middlewares/clubs.go#L23)
				- [AttachPartiesCollection](/middlewares/party.go#L21)
				- [GetUserClubFromToken](/middlewares/userClub.go#L198)
				- [PartiesUserClub](/userClub/parties.go#L15)

</details>
<details>
<summary>`/userclub/password/reset`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/password**
		- **/reset**
			- _POST_
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [ReadResetFromBody](/middlewares/customer.go#L217)
				- [Reset](/userClub/password_reset.go#L23)

</details>
<details>
<summary>`/userclub/promoter/{clubId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/promoter/{clubId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [SoftCombo](/userClub/soft_combo.go#L16)

</details>
<details>
<summary>`/userclub/vouchers`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/vouchers**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [VoucherReadByUserClub](/userClub/vouchers_read.go#L18)

</details>
<details>
<summary>`/userclub/{mode}/{userClubID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/{mode}/{userClubID}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Activate](/userClub/activate.go#L21)

</details>
<details>
<summary>`/userclub/{userClubID}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/userclub**
	- **/{userClubID}**
		- **/**
			- _PATCH_
				- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
				- [ReadUserClubPatchRequestFromBody](/middlewares/userClub.go#L150)
				- [AttachUserClubCollection](/middlewares/userClub.go#L21)
				- [GetUserClubFromToken](/middlewares/userClub.go#L198)
				- [Patch](/userClub/patch.go#L20)

</details>
<details>
<summary>`/version`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/version**
	- _GET_
		- [version.func1](/cmd/version.go#L24)

</details>
<details>
<summary>`/voucher`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [ReadVoucherPostRequestFromBody](/middlewares/vouchers.go#L97)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Create](/voucher/create.go#L16)

</details>
<details>
<summary>`/voucher/customer`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/customer**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [AppByCustomer](/voucher/appByCustomer.go#L14)

</details>
<details>
<summary>`/voucher/customer/next`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/customer/next**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [GetCustomerFromToken](/middlewares/customer.go#L109)
			- [ByCustomer](/voucher/byCustomer.go#L14)

</details>
<details>
<summary>`/voucher/invite`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/invite**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [AttachPartyProductsCollection](/middlewares/partyProduct.go#L16)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachInvitedCustomerCollection](/middlewares/invitedCustomer.go#L20)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [ReadVoucherPostRequestFromBody](/middlewares/vouchers.go#L97)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Invite](/voucher/invite.go#L20)

</details>
<details>
<summary>`/voucher/party/graph/soft/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/party/graph/soft/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachClubsCollection](/middlewares/clubs.go#L23)
			- [DashSoft](/voucher/dash_soft.go#L16)

</details>
<details>
<summary>`/voucher/party/soft/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/party/soft/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [ByPartySoft](/voucher/byPartySoft.go#L17)

</details>
<details>
<summary>`/voucher/party/{partyId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/party/{partyId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [ByParty](/voucher/byParty.go#L31)

</details>
<details>
<summary>`/voucher/refund/{voucherId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/refund/{voucherId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachInvoicesCollection](/middlewares/invoice.go#L16)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Refund](/voucher/refund.go#L16)

</details>
<details>
<summary>`/voucher/soft/{voucherId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/soft/{voucherId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [CheckToken](/middlewares/tokens.go#L81)
			- [HeaderScan](/middlewares/tokens.go#L49)
			- [Read](/voucher/read.go#L14)

</details>
<details>
<summary>`/voucher/transfer/{id}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/transfer/{id}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachCustomerCollection](/middlewares/customer.go#L21)
			- [ReadCustomerTranferableEmailFromBody](/middlewares/customer.go#L52)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Transfer](/voucher/transfer.go#L49)

</details>
<details>
<summary>`/voucher/undo/{voucherId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/undo/{voucherId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Undo](/voucher/undo.go#L22)

</details>
<details>
<summary>`/voucher/use/android/{voucherId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/use/android/{voucherId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [UseAndroid](/voucher/use_dash.go#L52)

</details>
<details>
<summary>`/voucher/use/soft/{voucherId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/use/soft/{voucherId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [ReadVoucherUsingConstrains](/middlewares/vouchers.go#L73)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [UseDash](/voucher/use_dash.go#L19)

</details>
<details>
<summary>`/voucher/use/{voucherId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/use/{voucherId}**
		- _POST_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [ReadVoucherUsingConstrains](/middlewares/vouchers.go#L73)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [AttachPartiesCollection](/middlewares/party.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Use](/voucher/use.go#L18)

</details>
<details>
<summary>`/voucher/{voucherId}`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/voucher**
	- **/{voucherId}**
		- _GET_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Read](/voucher/read.go#L14)
		- _DELETE_
			- [github.com/onnidev/api/vendor/github.com/cescoferraro/go-jwt-middleware.(*JWTMiddleware).Handler-fm](/report/routes.go#L17)
			- [AttachVoucherCollection](/middlewares/vouchers.go#L25)
			- [AttachUserClubCollection](/middlewares/userClub.go#L21)
			- [GetUserClubFromToken](/middlewares/userClub.go#L198)
			- [Delete](/voucher/delete.go#L14)

</details>
<details>
<summary>`/ws/app`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/ws/app**
	- **/**
		- _*_
			- [Handler.func1](/ws/handler.go#L10)

</details>
<details>
<summary>`/ws/dashboard`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/ws/dashboard**
	- **/**
		- _*_
			- [Handler.func1](/ws/handler.go#L10)

</details>
<details>
<summary>`/ws/staff`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/ws/staff**
	- **/**
		- _*_
			- [Handler.func1](/ws/handler.go#L10)

</details>
<details>
<summary>`/ws/test`</summary>

- [Cors](/middlewares/cors.go#L10)
- [Logger](/vendor/github.com/go-chi/chi/middleware/logger.go#L30)
- **/ws/test**
	- **/**
		- _GET_
			- [Routes.func4.1](/ws/routes.go#L30)

</details>

Total # of routes: 231
