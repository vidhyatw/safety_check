function renderSafetyScoreGraph(data){
    var safety_scores = data.safety_scores;
    var time_period = Array.apply(null, Array(24)).map(function (_, i) {return i;});

    var time_scores = Array.apply(0, Array(5));
    var backgroundColors = Array.apply(0, Array(24));
    var borderColors = Array.apply(0, Array(24));
    safety_scores.forEach(element => {
        time_scores[element.time] = element.score
        if(element.score < 2.5) {
            backgroundColors[element.time] = 'rgba(255, 99, 132, 0.2)'
            borderColors[element.time] = 'rgba(255,99,132,1)'
        } else if (element.score < 3.5) {
            backgroundColors[element.time] = 'rgba(255, 206, 86, 0.2)',
            borderColors[element.time] = 'rgba(255, 206, 86, 1)'
        }  else {
            backgroundColors[element.time] = 'rgba(75, 192, 192, 0.2)'
            borderColors[element.time] = 'rgba(75, 192, 192, 1)'
        }
    });
    var range = Array.apply(null, Array(10)).map(function (_, i) {return i+1;});
    var ctx = document.getElementById("myChart").getContext('2d');
var myChart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: time_period,
        datasets: [{
            data: time_scores,
            backgroundColor: backgroundColors,
            borderColor: borderColors,
            borderWidth: 1
        }]
    },
    options: {
        legend: {
            display: false,
        },
        scales: {
            yAxes: [{
                ticks: {
                    beginAtZero:true,
                    suggestedMax: 5
                },
                scaleLabel: {
                    display: true,
                    labelString: 'Safety Index'
                  }
            }],
            xAxes: [{
                barPercentage: 1,
                scaleLabel: {
                    display: true,
                    labelString: 'Hour of the day'
                  }
            },
        ]
        }
    }
});

}