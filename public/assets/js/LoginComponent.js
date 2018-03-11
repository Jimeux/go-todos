class AuthComponent {

  constructor(authService) {
    this.authService = authService;

    this.container = $(".login-component");
    this.title = $("#auth-title");
    this.primaryBtn = $("#auth-primary-btn");
    this.secondaryBtn = $("#auth-secondary-btn");
    this.username = $("#username");
    this.password = $("#password");
    this.alert = $("#auth-alert");
    this.loginMode = true;

    this.attachListeners();
  }

  toggle() {
    this.alert.hide();
    this.container.toggle();
    this.init();
  }

  init() {
    if (this.loginMode) {
      this.title.text("Login");
      this.primaryBtn.text("Login");
      this.secondaryBtn.text("Register");
    } else {
      this.title.text("Register");
      this.primaryBtn.text("Register");
      this.secondaryBtn.text("Login");
    }
    this.alert.hide();
  }

  attachListeners() {
    this.primaryBtn.on("click", (event) => {
      event.preventDefault();
      const username = this.username.val();
      const password = this.password.val();
      this.sendAuthRequest(username, password);
    });
    this.secondaryBtn.on("click", () => {
      event.preventDefault();
      this.loginMode = !this.loginMode;
      this.init();
    });
  }

  sendAuthRequest(username, password) {
    if (this.loginMode)
      this.authService.authenticate(username, password)
        .fail(() => {
          this.alert.text("Please check your username and password.").show();
        });
    else
      this.authService.register(username, password)
        .fail(() => {
          this.alert.text("Unable to register with those details.").show();
        });
  }

}
