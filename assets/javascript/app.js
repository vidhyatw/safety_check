$("#write_reviews").on("click",(function(event){
    event.preventDefault();
    $.get("/review/create", function(data, status){
        $( "#content" ).html( data )
        initAutocomplete()
        updateRating()
        window.history.pushState({urlPath:'#/review/create'},"",'/review/create')
    });
}));

$(document).on("click", "#submit-review", (function(event){
    event.preventDefault();
    var form_data = $("#create-review-form").serializeFormJSON();
    var review = {
        rating: parseFloat(form_data.rating),
        content: form_data.content,
        timestamp: new Date().getTime(),
        visitTime: form_data.visitTime,
        place: {
            placeid: form_data.placeId,
            coordinates: [parseFloat(form_data.lng), parseFloat(form_data.lat)]
        },
        reviewer: {
            email: form_data.email
        }
    }

    $.post( "/review/create", JSON.stringify (review), function(data, status){
    } );
}));


$(document).on("rateyo.set", "#rateYo", (function (e, data) {
    $("#rating").val(data.rating);
}));

(function ($) {
    $.fn.serializeFormJSON = function () {

        var o = {};
        var a = this.serializeArray();
        $.each(a, function () {
            if (o[this.name]) {
                if (!o[this.name].push) {
                    o[this.name] = [o[this.name]];
                }
                o[this.name].push(this.value || '');
            } else {
                o[this.name] = this.value || '';
            }
        });
        return o;
    };
})(jQuery);


function updateRating() {
 
    $("#rateYo").rateYo({
      rating: 0
    });
   
  };

  function getReviews(placeData) {
      var url = "/view/" + placeData.place_id + "/" + placeData.geometry.location.lat() + "/" 
      + placeData.geometry.location.lng()
    var url = "/review/view/ChIJGQ6k2QhYqDsRgkxMNsJi8Jw/76.9940433/11.054779"
    $.get(url, function(data) {
        $("#review-section").html(data);
    })
  }

  function getSafetyScores(placeData) {
    var url = "/review/score/ChIJGQ6k2QhYqDsRgkxMNsJi8Jw/76.9940433/11.054779"
    $.get(url, function(data) {
        renderSafetyScoreGraph(data)
    })
  } 

$.subscribe('place::changed', function (e, data) {
    $("#placeId").val(data.place_id)
    $("#lat").val(data.geometry.location.lat())
    $("#lng").val(data.geometry.location.lng())
    $("#reviewLocation").html(data.formatted_address);
    if($("#current_page").val() == "home") {
        getReviews(data);
        getSafetyScores(data);
    } 
});
