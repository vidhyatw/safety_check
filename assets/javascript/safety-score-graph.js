$(document).ready(function(){

// set the dimensions and margins of the graph
var margin = {top: 20, right: 20, bottom: 30, left: 40},
width = 560 - margin.left - margin.right,
height = 500 - margin.top - margin.bottom;

// set the ranges
var x = d3.scaleBand()
.range([0, width])
.padding(0.1);
var y = d3.scaleLinear();
//.domain([0, 5]).range([height, 0]);

// append the svg object to the body of the page
// append a 'group' element to 'svg'
// moves the 'group' element to the top left margin
var svg = d3.select("#graph").append("svg")
.attr("width", width + margin.left + margin.right)
.attr("height", height + margin.top + margin.bottom)
.append("g")
.attr("transform",
    "translate(" + margin.left + "," + margin.top + ")");
data = {
    safety_scores: [
    {time: 7, score: 1},
    {time: 9, score: 2}
]}

var safety_scores = data.safety_scores;
    var time_period = Array.apply(null, Array(24)).map(function (_, i) {return i;});

    var range = Array.apply(null, Array(10)).map(function (_, i) {return i+1;});
    //safety_scores.map(function (value, index) {return value.time}).sort();
    // Scale the range of the data in the domains
    x.domain(time_period);
    y.domain([0, d3.max(safety_scores, function(d) { return d.score; })]).range([height, 0]);
var ticks=axis.selectAll("line")
  .data(y.ticks(4)) // 20, 40, 60 and 80
  .enter().append("svg:line")
    y.nice()
    
    // append the rectangles for the bar chart
    svg.selectAll(".bar")
        .data(safety_scores)
        .enter().append("rect")
        .attr("class", "bar")
        .attr("x", function(d) { return x(d.time); })
        .attr("width", x.bandwidth())
        .attr("y", function(d) { return y(d.score); })
        .attr("height", function(d) { return height - y(d.score); });
    
    // add the x Axis
    svg.append("g")
        .attr("transform", "translate(0," + height + ")")
        .call(d3.axisBottom(x));
    
    // add the y Axis
    svg.append("g")
        .call(d3.axisLeft(y));
    



// get the data
// d3.csv("/assets/javascript/scores.csv", function(error, data) {
// if (error) throw error;

// // format the data
// data.forEach(function(d) {
//     d.sales = +d.sales;
// });

// // Scale the range of the data in the domains
// x.domain(data.map(function(d) { return d.salesperson; }));
// y.domain([0, d3.max(data, function(d) { return d.sales; })]);

// // append the rectangles for the bar chart
// svg.selectAll(".bar")
//     .data(data)
//     .enter().append("rect")
//     .attr("class", "bar")
//     .attr("x", function(d) { return x(d.salesperson); })
//     .attr("width", x.bandwidth())
//     .attr("y", function(d) { return y(d.sales); })
//     .attr("height", function(d) { return height - y(d.sales); });

// // add the x Axis
// svg.append("g")
//     .attr("transform", "translate(0," + height + ")")
//     .call(d3.axisBottom(x));

// // add the y Axis
// svg.append("g")
//     .call(d3.axisLeft(y));

// });
})

function renderGraph(data) {
    // if (error) throw error;
    
    // format the data
    // data.forEach(function(d) {
    //     d.sales = +d.sales;
    // });
    var safety_scores = data.safety_scores;
    var time_period = Array.apply(null, Array(24)).map(function (_, i) {return i+1;});
    //safety_scores.map(function (value, index) {return value.time}).sort();
    // Scale the range of the data in the domains
    x.domain(time_period);
    y.domain([1, d3.max(safety_scores, function(d) { return d.score; })]);
    
    // append the rectangles for the bar chart
    svg.selectAll(".bar")
        .data(safety_scores)
        .enter().append("rect")
        .attr("class", "bar")
        .attr("x", function(d) { return x(d.time); })
        .attr("width", x.bandwidth())
        .attr("y", function(d) { return y(d.score); })
        .attr("height", function(d) { return height - y(d.score); });
    
    // add the x Axis
    svg.append("g")
        .attr("transform", "translate(0," + height + ")")
        .call(d3.axisBottom(x));
    
    // add the y Axis
    svg.append("g")
        .call(d3.axisLeft(y));
    
    }