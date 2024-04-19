package api

// type ApiServer struct {
// 	listenAddr string
// 	store      *gorm.DB
// }

// func NewApiServer(store *gorm.DB, listenAddr string) *ApiServer {
// 	return &ApiServer{
// 		listenAddr: listenAddr,
// 		store:      store,
// 	}
// }

// func WriteJsonResponse(w http.ResponseWriter, code int, data interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	json.NewEncoder(w).Encode(data)
// }

// func (a *ApiServer) Start() error {
// 	// Health check handler
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Health!\n"))
// 	})

// 	// Account handlers
// 	http.HandleFunc("/account", a.HandleAccount)

// 	// Login handler
// 	http.HandleFunc("/login", a.handleLogin)

// 	// Roulette handlers
// 	http.Handle("/roulette", authMiddleware(http.HandlerFunc(a.handleRoulette)))

// 	// Dogecoin handler
// 	http.HandleFunc("/balance", a.handleDogeCoin)
// 	http.HandleFunc("/wallet", a.handleCreateWallet) //

// 	// Start the server
// 	log.Println("Starting server on", a.listenAddr)
// 	return http.ListenAndServe(a.listenAddr, nil)
// }

// func (a *ApiServer) HandleAccount(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == "GET" {
// 		a.handleGetAccount(w, r)
// 	} else if r.Method == "POST" {
// 		a.handleCreateAccount(w, r)
// 	} else {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		w.Write([]byte("Method not allowed\n"))
// 	}
// }

// func (a *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) {
// 	WriteJsonResponse(w, http.StatusOK, "Account get handler\n")
// }

// func (a *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) {

// 	type Input struct {
// 		Name     string `json:"name"`
// 		Password string `json:"password"`
// 		Email    string `json:"email"`
// 	}

// 	// Decode JSON input
// 	var input Input
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if input.Name == "" {
// 		http.Error(w, "Name is required", http.StatusBadRequest)
// 		return
// 	}
// 	if input.Password == "" {
// 		http.Error(w, "Password is required", http.StatusBadRequest)
// 		return
// 	}
// 	if input.Email == "" {
// 		http.Error(w, "Email is required", http.StatusBadRequest)
// 		return
// 	}

// 	// check if email exists in the database
// 	var account Account
// 	AuthService := NewAuthService()

// 	a.store.Where("email = ?", input.Email).First(&account)
// 	if account.Id != 0 {
// 		http.Error(w, "Email already exists", http.StatusBadRequest)
// 		return
// 	} else {
// 		// Create the account
// 		account = AuthService.CreateAccount(input.Name, input.Email, input.Password)
// 		// Store the account
// 		a.store.Create(&account)
// 	}

// 	// Generate JWT
// 	jwt, err := AuthService.GenerateJWT(account)
// 	if err != nil {
// 		log.Println("Failed to generate JWT")
// 		return
// 	}

// 	WriteJsonResponse(w, http.StatusOK,
// 		map[string]string{
// 			"message": "Account created",
// 			"token":   jwt,
// 		})
// }

// func (a *ApiServer) handleLogin(w http.ResponseWriter, r *http.Request) {

// 	type Input struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	// Decode JSON input
// 	var input Input
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if input.Password == "" {
// 		http.Error(w, "Password is required", http.StatusBadRequest)
// 		return
// 	}
// 	if input.Email == "" {
// 		http.Error(w, "Email is required", http.StatusBadRequest)
// 		return
// 	}

// 	var account Account
// 	a.store.Where("email = ?", input.Email).First(&account)
// 	if account.Id == 0 {
// 		http.Error(w, "Account not found", http.StatusBadRequest)
// 		return
// 	}

// 	// Generate JWT
// 	AuthService := NewAuthService()
// 	jwt, err := AuthService.GenerateJWT(account)
// 	if err != nil {
// 		log.Println("Failed to generate JWT")
// 		return
// 	}

// 	WriteJsonResponse(w, http.StatusOK,
// 		map[string]string{
// 			"message": "Login successful",
// 			"token":   jwt,
// 		})
// }

// // roulette play handler
// func (a *ApiServer) handleRoulette(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != "POST" {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		w.Write([]byte("Method not allowed\n"))
// 		return
// 	}

// 	userId := r.Header.Get("User")

// 	type Input struct {
// 		Bets []Bet `json:"bets"`
// 	}

// 	// Decode JSON input
// 	var input Input
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if len(input.Bets) == 0 {
// 		http.Error(w, "Bets are required", http.StatusBadRequest)
// 		return
// 	}

// 	var account Account
// 	a.store.Where("id = ?", userId).First(&account)
// 	if account.Id == 0 {
// 		http.Error(w, "Account not found", http.StatusBadRequest)
// 		return
// 	}

// 	game, err := PlayRoulette(account, input.Bets)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// update the account balance
// 	a.store.Model(&account).Update("balance", account.Balance+int64(game.Winnings)-int64(game.Amount))

// 	// Return the result to the user
// 	w.Write([]byte(fmt.Sprintf("Result: %d\n", game.Result)))
// 	w.Write([]byte(fmt.Sprintf("Winnings: %d\n", game.Winnings)))
// 	w.Write([]byte(fmt.Sprintf("Balance: %d\n", game.Balance)))
// }

// func (a *ApiServer) handleDogeCoin(w http.ResponseWriter, r *http.Request) {
// 	GetBallance()
// 	w.Write([]byte("Dogecoin balance\n"))
// }

// func (a *ApiServer) handleCreateWallet(w http.ResponseWriter, r *http.Request) {
// 	CreateWallet()
// 	w.Write([]byte("Wallet created\n"))
// }

// func authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Check if the request is authenticated
// 		AuthService := NewAuthService()

// 		tokenString := extractToken(r)

// 		if tokenString == "" {
// 			// Respond with unauthorized status code
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Unauthorized\n"))
// 			log.Print("Unauthorized request with no token")
// 			return
// 		}

// 		jwt, err := AuthService.VerifyJWT(tokenString)
// 		if err != nil {
// 			// Respond with unauthorized status code
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Unauthorized\n"))
// 			log.Print("Unauthorized request with invalid token", err)
// 			return
// 		}

// 		r.Header.Set("User", strconv.Itoa(jwt.UserID))

// 		// If authenticated, call the next handler
// 		next.ServeHTTP(w, r)
// 	})
// }

// func extractToken(r *http.Request) string {
// 	authHeader := r.Header.Get("Authorization")
// 	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 		return ""
// 	}
// 	return authHeader[len("Bearer "):]
// }
