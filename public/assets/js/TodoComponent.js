class TodoListComponent {

  constructor(todoService) {
    this.todoService = todoService;

    this.container = $(".todo-list-component");
    this.addBtn = $("#add-btn");
    this.titleInput = $("#add-input");
    this.todoList = $(".list-group");
    this.hideCompleteBtn = $("#hide-complete-btn");
    this.hideComplete = localStorage.getItem("hideComplete") === "true";

    this.attachListeners();
  }

  toggle() {
    if (this.container.is(":visible")) {
      this.container.hide();
    } else {
      this.updateHideCompleteBtn();
      this.todoService.findAll(this.hideComplete, todos => {
        this.todoList.empty();
        if (todos !== null)
          todos.forEach(todo => this.addTodoToDOM(todo, false, false))
      });
      this.container.show();
    }
  }

  attachListeners() {
    this.addBtn.on("click", () => this.createTodo());
    this.todoList.on("click", ".form-check-input", event =>
      this.completeTodo($(event.target))
    );
    this.hideCompleteBtn.on("click", () => {
      this.todoService.findAll(!this.hideComplete, todos => {
        this.toggleHideComplete();
        this.todoList.empty();
        if (todos !== null)
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
      this.titleInput.val("");
    }).always(() => this.addBtn.prop("disabled", false));
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
