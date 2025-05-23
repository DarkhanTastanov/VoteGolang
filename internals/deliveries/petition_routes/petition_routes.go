package petition_routes

import (
	"VoteGolang/internals/utils"
	"VoteGolang/pkg/domain"
	"log"
	"net/http"
)

func RegisterPetitionRoutes(mux *http.ServeMux, handler *PetitionHandler, tokenManager domain.TokenManager) {
	logRequest := func(route string, handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token, err := utils.ExtractTokenFromRequest(r)
			if err != nil {
				log.Printf("Failed to extract token: %v", err)
			} else {
				log.Printf("Accessing %s route | Method: %s | URL: %s | Token: %s", route, r.Method, r.URL.Path, token)
			}

			handlerFunc(w, r)
		}
	}

	mux.Handle("/petition/create", utils.JWTMiddleware(tokenManager)(
		logRequest("/petition_data/petition_repository/create", handler.CreatePetition),
	))

	mux.Handle("/petition/all", utils.JWTMiddleware(tokenManager)(
		logRequest("/petition_data/petition_repository/all", handler.GetAllPetitions),
	))
	mux.Handle("/petition/all/", utils.JWTMiddleware(tokenManager)(
		logRequest("/petition_data/petition_repository/all_by_page", handler.GetPetitionsByPage),
	))

	mux.Handle("/petition/get", utils.JWTMiddleware(tokenManager)(
		logRequest("/petition_data/petition_repository/get", handler.GetPetitionByID),
	))

	mux.Handle("/petition/vote", utils.JWTMiddleware(tokenManager)(
		logRequest("/petition_data/petition_repository/petition_data", handler.Vote),
	))

	mux.Handle("/petition/delete", utils.JWTMiddleware(tokenManager)(
		logRequest("/petition_data/petition_repository/delete", handler.DeletePetition),
	))
}
