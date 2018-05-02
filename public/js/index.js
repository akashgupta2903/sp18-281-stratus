function login() {
  var un = $("#logUsr").val();
  var ps = $("#logPwd").val();
  $.ajax({
    type: 'GET',
    url: '/login',
    data: {
      username: un,
      password: ps
    },
    success: function(res) {
      window.location.href = '/catalog';
    }
  });
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
  var un = $("#regUsr").val();
  var ps = $("#regPwd").val();
  $.ajax({
    type: 'GET',
    url: '/register',
    data: {
      username: un,
      password: ps
    },
    success: function(res) {
      window.location.href = '/catalog';
    }
  });
}