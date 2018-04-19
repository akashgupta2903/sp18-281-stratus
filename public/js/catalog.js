$(document).ready(function() {
  var username = getParameterByName('userid');
  $("#usernameInput").val(username);

  // get current user information
  $.ajax({
    type: 'GET',
    url: '/users/getUserInfo',
    data: {
      username: username
    },
    success: function(res) {
      var json = JSON.parse(res);
      var lgi = json.loginInfo;

      if(lgi.loggedIn) {
        $("#liLogin").hide();
        $("#liLogout").show();
      } else {
        $("#liLogout").hide();
        $("#liLogin").show();
      }

      if(lgi.loggedIn && lgi.username == username) {
        $("#spnName").html(lgi.firstName + " " + lgi.lastName);
        $("#spnUsername").html(lgi.username);
        $("#spnEmail").html(lgi.email);
      } else {
        $("button.loggedInOnly").hide();
        $("#vrtab h2").html(username + "'s Virtual Realities");
      }
    }
  });

  // setup upload
  var uploadForm = $("#fileForm");
  uploadForm.submit(function(e) {
    var formData = new FormData(this);
    e.preventDefault();
    $.ajax({
      type: 'POST',
      url: '/files/upload',
      processData: false,
      contentType: false,
      data: formData,
      error: function(err) {
        console.log("file upload error...");
      },
      success: function(res) {
        console.log("file upload success!!!");
        // clear the input fields
        document.getElementById("fileForm").reset();

        $("#uploadSuccessSpan").show(function() {
          $(this).fadeOut(5000);
        });
      }
    });
  });

  // display the user's VRs, if any
  getVRs();
});

// Switch to a tab on the user profile
function openTab(evt, tabName) {
  var i, tabcontent, tablinks;
  tabcontent = document.getElementsByClassName("tabcontent");
  for (i = 0; i < tabcontent.length; i++) {
    tabcontent[i].style.display = "none";
  }
  tablinks = document.getElementsByClassName("tablinks");
  for (i = 0; i < tablinks.length; i++) {
    tablinks[i].className = tablinks[i].className.replace(" active", "");
  }
  if(tabName === "vrtab") {
    getVRs();
  }
  document.getElementById(tabName).style.display = "block";
  evt.currentTarget.className += " active";
}

// Get & show VRs for this user
function getVRs() {
  var username = getParameterByName('userid');
  $.ajax({
    type: 'GET',
    url: '/files/uservrs',
    data: {
      username: username
    },
    success: function(res) {
      var json = JSON.parse(res);
      //console.log(json);
      var rows = json.result;

      var tbl = $("#vrtable").empty();
      var row,tr,td,img,h4,p;
      for(var i=0; i<rows.length; i++) {
        row = rows[i];
        if(i % 4 === 0) {
          if(i !== 0) {
            tr.appendTo(tbl);
          }
          tr = $("<tr/>");
        }
        td = $("<td/>");
        img = $("<img/>");
        img.attr("src", row.filepath)
           .attr("alt", row.filename)
           .attr("height", "200")
           .attr("width", "300")
           .appendTo(td);
        h4 = $("<h4/>").html(row.filename).appendTo(td);
        p = $("<p/>").html(row.description).appendTo(td);
        td.click(vrClickHandler(row.id));
        td.appendTo(tr);
      }
      tbl.append(tr);
    }
  });
}

// Returns a click handler for a VR
function vrClickHandler(id) {
  return function() {
    window.location.href = "/vr?vrid=" + id;
  }
}

// Enable submit file
function enableSubmit() {
  //$("#fileButton").prop("disabled", false);
  $("#fileSubmit").prop("disabled", false);
}

// Get a URL query parameter by name
function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}

// Show the log out dialog
function showLogoutDialog() {
  $("#confirmLogout").dialog({
    resizable: false,
    height: "auto",
    width: 400,
    modal: true,
    buttons: {
      "Log Out": function() {
        $.ajax({
          type: 'GET',
          url: '/users/logout',
          success: function(res) {
            $("#confirmLogout").dialog("close");
            window.location.href = '/';
          }
        });
      },
      Cancel: function() {
        $(this).dialog("close");
      }
    }
  });
}