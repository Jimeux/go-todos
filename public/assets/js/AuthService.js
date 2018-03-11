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

  register(username, password) {
    return $.post("/register", {username, password})
      .fail(TodoService.handleError)
      .done(token => {
        this.setToken(token);
        this.onAuthStateChanged();
      });
  }

  logout() {
    const token = this.getToken();
    this.setToken(null);

    return $.ajax({
      url: "/logout",
      method: "GET",
      headers: {"X-Auth-Token": token}
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
