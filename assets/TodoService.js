
class TodoService {

  create(title, callback) {
    return $.post("/todo", {title})
      .fail(TodoService.handleError)
      .done(callback);
  }

  updateCompleted(id, complete, callback) {
    return $.get(`/todo/${id}/complete?complete=${complete}`)
      .fail(TodoService.handleError)
      .done(callback);
  }

  findAll(hideComplete, callback) {
    return $.get(`/todo?hide_complete=${hideComplete}`)
      .fail(TodoService.handleError)
      .done(callback);
  }

  static handleError(err) {
    console.error(err);
  }

}
