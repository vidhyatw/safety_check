// This is called with the results from from FB.getLoginStatus().
function statusChangeCallback(response) {
    console.log('statusChangeCallback');
    console.log(response);
    // The response object is returned with a status field that lets the
    // app know the current login status of the person.
    // Full docs on the response object can be found in the documentation
    // for FB.getLoginStatus().
    if (response.status === 'connected') {
      // Logged into your app and Facebook.
      testAPI();
    } else if (response.status === 'unknown'){
      sessionStorage.setItem("login", "false")
      // The person is not logged into your app or we are unable to tell.
      document.getElementById('login_status').innerHTML = 'Please login to write reviews';
    }
  }


  // This function is called when someone finishes with the Login
  // Button.  See the onlogin handler attached to it in the sample
  // code below.
  function checkLoginState() {
    FB.getLoginStatus(function(response) {
      statusChangeCallback(response);
    });
  }

  $(document).on("click", "#logout", function logout() {
    FB.logout(function(response) {
      sessionStorage.setItem("login", "false")
      document.getElementById('login_status').innerHTML = 'Please login to write reviews';
      $("#logout").hide()
      $("#login").show()
      if($("#current_page").val()!="home"){
        window.location.href="/"
      }
    })
});

  window.fbAsyncInit = function() {
    var appId = $("#fb-app-id").val();
    FB.init({
      appId      : appId,
      cookie     : true,  // enable cookies to allow the server to access 
                          // the session
      xfbml      : true,  // parse social plugins on this page
      version    : 'v2.12' // use graph api version 2.12
    });

    // Now that we've initialized the JavaScript SDK, we call 
    // FB.getLoginStatus().  This function gets the state of the
    // person visiting this page and can return one of three states to
    // the callback you provide.  They can be:
    //
    // 1. Logged into your app ('connected')
    // 2. Logged into Facebook, but not your app ('not_authorized')
    // 3. Not logged into Facebook and can't tell if they are logged into
    //    your app or not.
    //
    // These three cases are handled in the callback function.

    FB.getLoginStatus(function(response) {
      statusChangeCallback(response);
    });

  };

  // Load the SDK asynchronously
  (function(d, s, id) {
    var js, fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return;
    js = d.createElement(s); js.id = id;
    js.src = "https://connect.facebook.net/en_US/sdk.js";
    fjs.parentNode.insertBefore(js, fjs);
  }(document, 'script', 'facebook-jssdk'));

  // Here we run a very simple test of the Graph API after login is
  // successful.  See statusChangeCallback() for when this call is made.
  function testAPI() {
    console.log('Welcome!  Fetching your information.... ');
    FB.api('/me', { locale: 'en_US', fields: 'name,email,gender' }, function(response) {
      debugger;
      sessionStorage.setItem("login",true)
      sessionStorage.setItem("reviewer_name",response.name)
      sessionStorage.setItem("reviewer_gender",response.gender)
      sessionStorage.setItem("reviewer_email",response.email)

      $("#logout").show()
      $("#login").hide()
      console.log('Successful login for: ' + response.name);
      
      document.getElementById('login_status').innerHTML =
        'Thanks for logging in, ' + response.name + '!';
    });
  }

  

  
