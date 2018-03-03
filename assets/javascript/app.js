$("#write_reviews").click(function(){
    alert("hi");
    $.get("/review/new", function(data, status){
        document.write(data);
    });
});