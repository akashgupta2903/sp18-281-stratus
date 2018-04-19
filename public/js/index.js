function login() {
  // TODO: ajax request to login route, redirect to catalog
}

function showForm(register) {
  if(register) {
    $("#loginForm").hide();
    $("#registerForm").show();
  } else {
    $("#loginForm").show();
    $("#registerForm").hide();
  }
}

function register() {
  // TODO: ajax request to register route, redirect to catalog
}

$(document).ready(function() {
  console.log("hello world");
});