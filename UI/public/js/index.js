function login() {
  $("#logUsr").css("border-color", "darkgrey"); 
  $("#logPwd").css("border-color", "darkgrey"); 
  var un = $("#logUsr").val();
  var ps = $("#logPwd").val();
  if(un === "") {
    $("#logUsr").css("border-color", "red"); 
  } else if(ps === "") {
    $("#logPwd").css("border-color", "red");
  } else {
    $.ajax({
      type: 'GET',
      url: '/login',
      data: {
        username: un,
        password: ps
      },
      success: function(res) {
        window.location.href = '/catalog';
      },
      statusCode: {
        500: function() {
          $("#logUsr").val("");
          $("#logPwd").val("");
        }
      }
    });
  }
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
  $("#regUsr").css("border-color", "darkgrey"); 
  $("#regPwd").css("border-color", "darkgrey");
  $("#regEmail").css("border-color", "darkgrey"); 
  var un = $("#regUsr").val();
  var ps = $("#regPwd").val();
  var em = $("#regEmail").val();
  if(un === "") {
    $("#regUsr").css("border-color", "red"); 
  } else if(ps === "") {
    $("#regPwd").css("border-color", "red");
  } else if(em === "") {
    $("#regEmail").css("border-color", "red");
  } else {
    $.ajax({
      type: 'GET',
      url: '/register',
      data: {
        username: un,
        password: ps,
        email: em
      },
      success: function(res) {
        window.location.href = '/catalog';
      },
      statusCode: {
        500: function() {
          $("#regUsr").val("");
          $("#regPwd").val("");
        }
      }
    });
  }
}