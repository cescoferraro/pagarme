package recipient

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Routes is amazing
func Routes(r chi.Router) {
	r.Route("/recipient", func(j chi.Router) {
		j.Use(shared.JWTAuth.Handler)
		j.Use(middlewares.AttachRecipientsCollection)

		// TODO: somente Profile ONNi
		j.Get("/", ListRecipients)

		j.Get("/{recipientID}", Read)
		mC := middlewares.ReadRecipientPostRequestFromBody
		j.With(mC).Post("/", Create)
		mP := middlewares.ReadRecipientPatchFromBody
		j.With(mP).Patch("/{recipientID}", Patch)

		j.Route("/club", func(n chi.Router) {
			n.Get("/{clubId}", ListClubRecipients)
		})

		j.Route("/balance", func(n chi.Router) {
			n.With(middlewares.ReadFinanceQueryFromBody).Post("/", BalanceTransactions)
		})
		j.Route("/pagarme", func(n chi.Router) {
			n.Get("/{recipientID}", PagarMeRecipient)
			n.With(middlewares.ReadRecipientTweaksPatchFromBody).
				Patch("/tweaks/{recipientID}", PagarMeRecipientTweaks)
			n.With(middlewares.ReadRecipientWithDrawRequestFromBody).
				Post("/withdraw/{recipientID}", WithDraw)
			n.Get("/withdraws/{recipientID}", WithDraws)
		})

		j.Route("/antecipation", func(n chi.Router) {
			m1 := middlewares.ReadAntecipationPostRequestFromBody
			n.Delete("/{recipientID}/{bulkID}", DeleteAnteciapation)
			n.Post("/confirm/{recipientID}/{bulkID}", ConfirmAnteciapation)
			n.Post("/cancel/{recipientID}/{bulkID}", CancelAnteciapation)
			n.With(m1).Post("/", CreateAnteciapation)
			n.Get("/{recipientID}", Anteciapations)
			n.With(m1).Put("/{bulkID}", EditAnteciapation)
			n.With(middlewares.ReadFinanceQueryFromBody).Post("/limits", AnteciapationsLimit)
		})
		j.Route("/transactions", func(n chi.Router) {
			n.Use(middlewares.ReadFinanceQueryFromBody)
			n.Post("/", ListTransactions)
		})

		j.Route("/payables", func(n chi.Router) {
			n.Use(middlewares.ReadFinanceQueryFromBody)
			n.Post("/timeline", PayablesTimeline)
			n.Post("/", Payables)
		})

		j.Route("/operations", func(n chi.Router) {
			n.Use(middlewares.ReadFinanceQueryFromBody)
			n.Post("/timeline", DaysBalanceTransactions)
			n.Post("/", Operations)
			n.Post("/xlsx", Xlsx)
		})

	})
}
