package routes

import (
	"github.com/JorgeLeonardoLF/Authentication/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

/*
The routes Setup() function should:
- Take a parent router of type *chi.Mux
- It will mount on to the parent router a new route to make versioning easier vXrouter of type *chi.Mux route
- The vXrouter will have the http request type, path, and controller on to it
*/
func Setup(parentRoute *chi.Mux) {

	//Setup Parent router configurations
	parentRoute.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http//*"},                    // Allow sending through these
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow these methods
		AllowedHeaders:   []string{"*"},                                       // Allow any headers
		ExposedHeaders:   []string{"Link"},                                    //Hover over for more details and google search for better explanations {Note Stamp 2 End}
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//Initialize vXRouter
	v1Router := chi.NewRouter()

	/*Setup routes:
	- request types: {"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	- "/paths": "/examplePath"
	- request handler methods: controllers._____
	*/
	v1Router.Get("/healthz", controllers.CheckServerHealth)
	v1Router.Post("/register", controllers.RegisterNewAccount)

	//finally mount the vXRouter to the parent router
	parentRoute.Mount("/v1", v1Router)
}
