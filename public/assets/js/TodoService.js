class TodoService {

  constructor(authService) {
    this.authService = authService;
  }

  create(title, callback) {
    return this.authRequest("/todo", "POST", {title})
      .fail(TodoService.handleError)
      .done(callback);
  }

  updateCompleted(id, complete) {
    return this.authRequest(`/todo/${id}/complete`, "GET", {complete})
      .fail(TodoService.handleError);
  }

  findAll(hideComplete, callback) {
    return this.authRequest("/todo", "GET", {"hide_complete": hideComplete})
      .fail(TodoService.handleError)
      .done(callback);
  }

  authRequest(url, method = "GET", data = {}) {
    const token = this.authService.getToken();

    return $.ajax({
      url,
      method,
      data,
      headers: {"X-Auth-Token": token}
    }).fail(res => {
      if (res.status === 403) {
        this.authService.onAuthStateChanged();
      }
    });
  }

  static handleError(err) {
    console.error(err);
  }

}