# Netflix
Netflix api. Uses mongoDB Atlas

router.HandleFunc("/", controller.GetMyAllMovies).Methods("GET")
router.HandleFunc("/", controller.CreateMovie).Methods("POST")
router.HandleFunc("/{id}", controller.DeleteOneMovie).Methods("DELETE")
router.HandleFunc("/{id}", controller.MarkAsWatched).Methods("PUT")
