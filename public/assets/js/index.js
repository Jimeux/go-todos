
$(() => {

  const authService = new AuthService();
  const todoService = new TodoService(authService);

  const todoComponent = new TodoListComponent(todoService);
  const loginComponent = new LoginComponent(authService);

  authService.setOnAuthStateChanged(() => {
    todoComponent.toggle();
    loginComponent.toggle();
  });

  if (authService.isAuthenticated())
    todoComponent.toggle();
  else
    loginComponent.toggle();

});
