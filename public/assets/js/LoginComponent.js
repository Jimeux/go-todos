class LoginComponent {

  constructor(authService) {
    this.authService = authService;

    this.container = $(".login-component");
    this.loginBtn = $("#login-btn");
    this.username = $("#username");
    this.password = $("#password");

    this.attachListeners();
  }

  toggle() {
    this.container.toggle();
  }

  attachListeners() {
    this.loginBtn.on("click", (event) => {
      event.preventDefault();
      const username = this.username.val();
      const password = this.password.val();
      this.authService.authenticate(username, password);
    });
  }

}
