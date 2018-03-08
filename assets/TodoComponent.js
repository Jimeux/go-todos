
class TodoListComponent {

  constructor(todoService) {
    this.todoService = todoService;
    this.addBtn = $("#add-btn");
    this.titleInput = $("#add-input");
    this.todoList = $(".list-group");
    this.hideCompleteBtn = $("#hide-complete-btn");
    this.hideComplete = localStorage.getItem("hideComplete") === "true";

    this.updateHideCompleteBtn();
    this.attachListeners();

    this.todoService.findAll(this.hideComplete, todos => {
      todos.forEach(todo => this.addTodoToDOM(todo, false, false))
    });
  }

  attachListeners() {
    this.addBtn.click(() => this.createTodo());
    this.todoList.on("click", ".form-check-input", event =>
      this.completeTodo($(event.target))
    );
    this.hideCompleteBtn.click(() => {
      this.todoService.findAll(!this.hideComplete, todos => {
        this.toggleHideComplete();
        this.todoList.empty();
        todos.forEach(todo => this.addTodoToDOM(todo, false, false));
      });
    });
  }

  toggleHideComplete() {
    this.hideComplete = !this.hideComplete;
    localStorage.setItem("hideComplete", this.hideComplete);
    this.updateHideCompleteBtn();
  }

  updateHideCompleteBtn() {
    const value = this.hideComplete === true ? "Show complete" : "Hide complete";
    this.hideCompleteBtn.html(value);
  }

  completeTodo(checkbox) {
    const checked = checkbox.prop("checked");
    const id = checkbox.data("id");
    const container = checkbox.closest("li");

    this.todoService.updateCompleted(id, checked, () => {
      container.toggleClass("list-group-item-primary", checked);
      container.find("a").toggleClass("complete", checked);
      if (this.hideComplete)
        container.slideUp();
    });
  }

  createTodo() {
    const title = this.titleInput.val();

    this.addBtn.prop("disabled", true);

    this.todoService.create(title, todo => {
      this.addTodoToDOM(todo, true, true);
      this.addBtn.prop("disabled", false);
      this.titleInput.val("");
    });
  }

  addTodoToDOM(todo, prepend, animate) {
    const todoElement = this.renderTodo(todo);
    todoElement.hide();

    prepend ?
      todoElement.prependTo(this.todoList) :
      todoElement.appendTo(this.todoList);

    animate ?
      todoElement.slideDown() :
      todoElement.show();
  }

  renderTodo(todo) {
    return $(`
       <li class="list-group-item ${todo.complete ? "list-group-item-primary" : ""}">
         <a href="#" data-id=${todo.id}" class="${todo.complete ? "complete" : ""}">
           ${todo.title}
         </a>
         <div class="form-check form-check-inline float-right">
           <input class="form-check-input"
                  type="checkbox"
                  data-id=${todo.id}
                  ${todo.complete ? "checked" : ""}>
         </div>
       </li>
    `);
  }

}
