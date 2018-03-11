class AuthService {

  constructor() {
    this.onAuthStateChanged = () => {}
  }

  setOnAuthStateChanged(fn) {
    this.onAuthStateChanged = fn;
  }

  authenticate(username, password) {
    return $.post("/login", {username, password})
      .fail(TodoService.handleError)
      .done(token => {
        this.setToken(token);
        this.onAuthStateChanged();
      });
  }

  setToken(token) {
    localStorage.setItem("auth-token", token);
  }

  getToken() {
    return localStorage.getItem("auth-token");
  }

  isAuthenticated() {
    return this.getToken() !== null;
  }

}
